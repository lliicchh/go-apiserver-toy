package errno

var (
	// common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error",}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred when binding the request to the struct"}
	ErrValidation       = &Errno{Code: 20001, Message: "Validation failed",}
	ErrDatabase         = &Errno{Code: 20002, Message: "DataBase error",}
	ErrToken            = &Errno{Code: 20003, Message: "Error occured when sign the JSON web token",}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found",}
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occured when encrypt the user password",}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The Token was invalid",}
	ErrPassWordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect",}
)
