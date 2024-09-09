package e

type ErrorCode uint

const (
	UserNotFound  = 1001001
	UserIDInvalid = iota
	UserPasswordInvalid

	Success     ErrorCode = 200
	SystemError           = 500
)
