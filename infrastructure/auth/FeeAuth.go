package auth

/*
create by: Hoangnd
create at: 2023-01-01
des: Xử lý dữ liệu từ token của user
*/

import (
	"fmt"
	"io"
	"os"

	sAuth "aBet/adapters/auth"

	"github.com/valyala/fastjson"
)

type authObject struct {
	id            string
	userName      string
	displayName   string
	email         string
	resetPswdInfo string
	accType       string
	userAgent     string
	ip            string
	role          int
	exp           int64
	token         string
}

func NewAuthObject(data []byte, jwt string) (sAuth.AuthObject, error) {
	var p fastjson.Parser
	v, err := p.Parse(string(data))
	if err != nil {
		return nil, err
	}
	userId := string(v.GetStringBytes("id"))
	userDisplayName := string(v.GetStringBytes("userName"))
	userType := v.GetInt("userType")
	return &authObject{
			id:       userId,
			userName: userDisplayName,
			role:     userType,
			token:    jwt,
		},
		err
}

func (au *authObject) GetUserId() string {
	fmt.Println(au, "auuuuuuuuuuuuuuuuuuu")
	return au.id
}
func (au *authObject) GetUserDisplayName() string {
	return au.displayName
}
func (au *authObject) GetUserUserName() string {
	return au.userName
}
func (au *authObject) GetUserEmail() string {
	return au.email
}
func (au *authObject) GetUserResetPswdInfo() string {
	return au.resetPswdInfo
}
func (au *authObject) GetUserAccType() string {
	return au.accType
}
func (au *authObject) GetUserIp() string {
	return au.ip
}
func (au *authObject) GetUserRole() int {
	return au.role
}
func (au *authObject) GetUserExp() int64 {
	return au.exp
}
func (au *authObject) GetToken() string {
	return au.token
}

func (au *authObject) GetAll() interface{} {
	return map[string]interface{}{
		"id":            au.id,
		"userName":      au.userName,
		"displayName":   au.displayName,
		"email":         au.email,
		"resetPswdInfo": au.resetPswdInfo,
		"accType":       au.accType,
		"userAgent":     au.userAgent,
		"ip":            au.ip,
		"role":          au.role,
		"exp":           au.exp,
		"token":         au.token,
	}
}

func GetPublicKey() []byte {
	file, err := os.Open("crypt/public3.pem")
	if err != nil {
		return nil
	}
	defer file.Close()
	fileByte, _ := io.ReadAll(file)
	return fileByte
}
