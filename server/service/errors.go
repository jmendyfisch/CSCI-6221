// All the service end errors to be propagated.

package service

import "errors"

var (
	ErrInvalidLawyerID = errors.New("invalid lawyer id")
	ErrInvalidCaseID   = errors.New("invalid case id")
	ErrQueryFailure    = errors.New("execute query failed")
)
