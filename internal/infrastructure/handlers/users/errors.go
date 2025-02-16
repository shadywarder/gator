package users

type usersError struct{ msg string }

func (e usersError) Error() string { return e.msg }

var ErrUserExistance = usersError{msg: "error: user does not exist or cannot be found"}
var ErrInvalidUserName = usersError{msg: "error: name is invalid or does not meet the requirements"}
var ErrUserCreation = usersError{msg: "error: failed to create user"}
var ErrTableReset = usersError{msg: "error: failed to reset table"}
var ErrRetrieveUser = usersError{msg: "error: failed to retrieve user"}
