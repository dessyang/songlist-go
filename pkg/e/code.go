package e

const (
	Success          = 0
	Fail             = 1
	NotLogin         = 2
	NotFound         = 404
	InvalidOperation = 500

	MissParam           = 40000
	UsernameFormatError = 40001
	PasswordFormatError = 40002
	EmailFormatError    = 40003

	RepeatUser  = 50000
	RepeatEmail = 50001
	AuthFail    = 50002

	PageOutBound = iota + 20000
	SongNotFound

	ParamNotNul
	RegistrationNotAllowed
	MethodError
	PageNotNum
)
