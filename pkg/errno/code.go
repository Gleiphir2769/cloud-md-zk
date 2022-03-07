package errno

// Error code description:
// 				1						00						02
// 		service level error		service module code		specific error code
//
// service level error: 1:system error, 2:common error
var (
	// Common errors
	OK                   = &Errno{Code: 200, Message: "success"}
	InternalServerError  = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind              = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrValidation        = &Errno{Code: 10003, Message: "Validation failed."}
	ErrInvalidParameters = &Errno{Code: 10004, Message: "Invalid parameters"}

	// zk-operator operations errors
	ErrCreate  = &Errno{Code: 20001, Message: "ZK operator create failed"}
	ErrDelete  = &Errno{Code: 20002, Message: "ZK operator delete failed"}
	ErrUpdate  = &Errno{Code: 20003, Message: "ZK operator update failed"}
	ErrGetInfo = &Errno{Code: 20004, Message: "ZK operator get cluster info failed"}
)
