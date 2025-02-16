package util

type utilError struct{ msg string }

func (e utilError) Error() string { return e.msg }

var ErrTableReset = utilError{msg: "error: failed to reset table"}
