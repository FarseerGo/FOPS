package adminService

import (
	"errors"
	"fops/domain/security/admin"
	"fs/utils/encrypt"
)

// ChangePwd 修改密码
func ChangePwd(repository admin.Repository, userName string, pwd string, newPwd string) error {
	if newPwd == "" {
		return errors.New("请输入新密码")
	}

	admin := repository.ToInfoByUsername(userName, encrypt.Md5(pwd))
	if admin.Id == 0 {
		return errors.New("原密码错误，请重新输入")
	}
	err := admin.EncryptPwd(newPwd)
	if err != nil {
		return err
	}

	repository.Update(admin.Id, admin)
	return nil
}
