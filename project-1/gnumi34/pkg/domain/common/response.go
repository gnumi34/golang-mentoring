package common

const (
	DataSuccess     = "success processing data"
	DataFailed      = "failed processing data"
	UnknownError    = "unknown error detected. please try again later"
	RecordNotFound  = "your requested entity is not found"
	ValidationError = "please check the validation error"
	URINotFound     = "your requested URI is not found"

	UserAlreadyCreated = "the user is already created"
	PasswordNotMatch   = "invalid username/password. please check again"
	PasswordNotSame    = "password 1 and password 2 must be same"
	PasswordNotFilled  = "please fill the password"
	InvalidPassword    = "password must have at least 8 characters with minimum of 1 uppercase, 1 number, and 1 special character"
	InvalidUserID      = "invalid user id"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
	Code    int         `json:"code"`
}
