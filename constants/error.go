package constants

// general
const (
	InternalServerErrCode = 1001 + iota
	OrmHookValidationErrCode
	NotAuthorizedErrCode
	BadRequestErrCode
)

// repositories
const (
	QueryInternalServerErrCode = 2001 + iota
	QueryNotFoundErrCode
)

// user
const (
	RegisterEmailNotAvailableErrCode = 3001 + iota
	RegisterUsernameNotAvailableErrCode
	HashPasswordInternalErrCode
	LoginUsernameNotFoundErrCode
	LoginEmailNotFoundErrCode
	LoginInvalidPasswordErrCode
)
