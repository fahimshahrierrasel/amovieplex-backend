package errors

type MoviePlexError int

const (
	ERRReqBody MoviePlexError = iota
	ERRAuthRegHash
	ERRAuthRegInvRole
	ERRAuthLogin
	ERRAuthLoginTokGen
	ERRUnauthorized
	ERRNotFound
)

func (mpe MoviePlexError) ErrorCode() string {
	return [...]string{
		"ERRReqBody",
		"ERRAuthRegHash",
		"ERRAuthRegInvRole",
		"ERRAuthLogin",
		"ERRAuthLoginTokGen",
		"ERRUnauthorized",
		"ERRNotFound",
	}[mpe]
}

func (mpe MoviePlexError) ErrorMessage() string {
	return [...]string{
		"Given request body is not correct",
		"Error Generating Password Hash",
		"Given role is not valid",
		"User/Password not matched",
		"Token generation failed",
		"Unauthorized user or token is not valid",
		"Requested object not found",
	}[mpe]
}

func ErrorCodeMessage(error MoviePlexError) string {
	return error.ErrorCode() + ": " + error.ErrorMessage()
}
