package errors

type MoviePlexError int

const (
	None MoviePlexError = iota
	ERRReqBody
	ERRAuthRegHash
	ERRAuthRegInvRole
	ERRAuthLogin
	ERRAuthLoginTokGen
	ERRUnauthorized
	ERRNotFound
	ERRInvObjID
)

func (mpe MoviePlexError) ErrorCode() string {
	return [...]string{
		"None",
		"ERRReqBody",
		"ERRAuthRegHash",
		"ERRAuthRegInvRole",
		"ERRAuthLogin",
		"ERRAuthLoginTokGen",
		"ERRUnauthorized",
		"ERRNotFound",
		"ERRInvObjID",
	}[mpe]
}

func (mpe MoviePlexError) ErrorMessage() string {
	return [...]string{
		"None",
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
