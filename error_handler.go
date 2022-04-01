package jwtmiddleware

import (
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ErrorHandler is a handler which is called when an error occurs in the
// JWTMiddleware. Among some general errors, this handler also determines the
// response of the JWTMiddleware when a token is not found or is invalid. The
// err can be checked to be ErrJWTMissing or ErrJWTInvalid for specific cases.
// The default handler will return a status code of 400 for ErrJWTMissing,
// 401 for ErrJWTInvalid, and 500 for all other errors. If you implement your
// own ErrorHandler you MUST take into consideration the error types as not
// properly responding to them or having a poorly implemented handler could
// result in the JWTMiddleware not functioning as intended.
type ErrorHandler func(c *gin.Context, err error)

// DefaultErrorHandler is the default error handler implementation for the
// JWTMiddleware. If an error handler is not provided via the WithErrorHandler
// option this will be used.
func DefaultErrorHandler(c *gin.Context, err error) {
	switch {
	case errors.Is(err, jwtmiddleware.ErrJWTMissing):
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "JWT is missing.",
		})
	case errors.Is(err, jwtmiddleware.ErrJWTInvalid):
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "JWT is invalid.",
		})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while checking the JWT.",
		})
	}
}

// invalidError handles wrapping a JWT validation error with
// the concrete error ErrJWTInvalid. We do not expose this
// publicly because the interface methods of Is and Unwrap
// should give the user all they need.
type invalidError struct {
	details error
}

// Is allows the error to support equality to ErrJWTInvalid.
func (e invalidError) Is(target error) bool {
	return target == jwtmiddleware.ErrJWTInvalid
}

// Error returns a string representation of the error.
func (e invalidError) Error() string {
	return fmt.Sprintf("%s: %s", jwtmiddleware.ErrJWTInvalid, e.details)
}

// Unwrap allows the error to support equality to the
// underlying error and not just ErrJWTInvalid.
func (e invalidError) Unwrap() error {
	return e.details
}
