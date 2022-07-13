package adminService

import (
	"errors"
	"fops/domain/security/admin"
	"fs/utils/encrypt"
)

// Login 登陆
func Login(repository admin.Repository, userName string, pwd string, ip string) (admin admin.DomainObject, err error) {
	pwd = encrypt.Md5(pwd)
	admin = repository.ToInfoByUsername(userName, pwd)

	if admin.Id == 0 {
		return admin, errors.New("用户不存在，或者密码错误")
	}

	if !admin.IsEnable {
		return admin, errors.New("用户不存在，或者密码错误")
	}

	admin.SetLoginIp(ip)
	repository.Update(admin.Id, admin)
	return admin, nil
}
