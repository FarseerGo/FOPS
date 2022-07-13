package adminService

import (
	"errors"
	"fops/domain/security/admin"
)

// Add 添加管理员
func Add(repository admin.Repository, admin admin.DomainObject) (int, error) {
	err := admin.EncryptPwd(admin.UserPwd)
	if err != nil {
		return 0, err
	}

	var isExists = repository.IsExists(admin.UserName)
	if isExists {
		return 0, errors.New("管理员名称已存在")
	}
	return repository.Add(admin), nil
}
