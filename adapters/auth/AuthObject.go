package auth

type AuthObject interface {
	GetUserId() string
	GetUserDisplayName() string
	GetUserUserName() string
	GetUserEmail() string
	GetUserResetPswdInfo() string
	GetUserAccType() string
	GetUserIp() string
	GetUserRole() int
	GetUserExp() int64
	GetToken() string
}
