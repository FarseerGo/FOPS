package adminService

import (
	"errors"
	"fops/domain/security/admin"
)

// Update 修改管理员
func Update(repository admin.Repository, admin admin.DomainObject) error {
	err := admin.EncryptPwd(admin.UserPwd)
	if err != nil {
		return err
	}

	var isExists = repository.IsExistsWithoutSelf(admin.UserName, admin.Id)
	if isExists {
		return errors.New("管理员名称已存在")
	}

	repository.Update(admin.Id, admin)
	return nil
}
