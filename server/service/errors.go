package service

import "errors"

var (
	ErrInvalidLawyerID = errors.New("invalid lawyer id")
	ErrQueryFailure    = errors.New("execute query failed")
)
