package adminService

import (
	"errors"
	"fops/domain/security/admin"
	"github.com/farseernet/farseer.go/utils/encrypt"
)

// ChangePwd 修改密码
func ChangePwd(repository admin.Repository, userName string, pwd string, newPwd string) error {
	if newPwd == "" {
		return errors.New("请输入新密码")
	}

	do := repository.ToInfoByUsername(userName, encrypt.Md5(pwd))
	if do.Id == 0 {
		return errors.New("原密码错误，请重新输入")
	}
	do.EncryptPwd(newPwd)

	repository.Update(do.Id, do)
	return nil
}
