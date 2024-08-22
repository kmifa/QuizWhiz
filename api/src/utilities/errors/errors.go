package errors

import (
	"net/http"

	"github.com/morikuni/failure"
)

var (
	// Unauthenticated means the request does not have valid
	Unauthenticated failure.StringCode = "Unauthenticated"

	// Unauthorized means the request does not have valid
	Unauthorized failure.StringCode = "Unauthorized"

	// Forbidden means the request does not have valid
	Forbidden failure.StringCode = "Forbidden"

	// InvalidArgument indicates client specified an invalid argument.
	// Note that this differs from FailedPrecondition. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	InvalidArgument failure.StringCode = "InvalidArgument"

	// NotFound means some requested entity (e.g., file or directory) was
	// not found.
	NotFound failure.StringCode = "NotFound"

	// FailedPrecondition indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	FailedPrecondition failure.StringCode = "FailedPrecondition"

	// FailedDependency indicates operation failed because of a failed
	// dependency.
	Aborted failure.StringCode = "Aborted"

	// Internal errors. Means some invariants expected by underlying
	// system has been broken. If you see one of these errors,
	// something is very broken.
	Internal failure.StringCode = "Internal"

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	AlreadyExists failure.StringCode = "AlreadyExists"

	// NotImplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	UnImplemented failure.StringCode = "UnImplemented"
)

func GetHTTPStatus(err error) int {
	c, ok := failure.CodeOf(err)
	if !ok {
		return http.StatusInternalServerError
	}
	switch c {
	case Unauthenticated, Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case InvalidArgument:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case FailedPrecondition:
		return http.StatusPreconditionFailed
	case Aborted:
		return http.StatusConflict
	case AlreadyExists:
		return http.StatusConflict
	case UnImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}
}
