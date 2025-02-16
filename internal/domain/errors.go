package domain

type domainError struct{ msg string }

func (e domainError) Error() string { return e.msg }

var ErrInvalidCommand = domainError{msg: "error: provided command does not exist"}
