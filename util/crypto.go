package util

import "golang.org/x/crypto/bcrypt"

//EncodePwd 密码加密
func EncodePwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) //加密处理
	return string(hash), err
}

//CheckPwd 密码校验
func CheckPwd(pwd string, pwd1 string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(pwd1)) //验证（对比）
	if err != nil {
		return false, err
	}
	return true, err
}
