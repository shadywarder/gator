package middleware

type middlewareError struct{ msg string }

func (e middlewareError) Error() string { return e.msg }

var ErrReceiveUser = middlewareError{msg: "error: failed to retrieve user"}
