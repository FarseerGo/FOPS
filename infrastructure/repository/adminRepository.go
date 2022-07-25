package repository

import (
	"fops/domain/security/admin"
	"fops/infrastructure/repository/agent/adminAgent"
	"fops/infrastructure/repository/context"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() admin.Repository {
		return &adminRepository{
			data.Init[context.MysqlContext]().Admin,
		}
	})
}

type adminRepository struct {
	data.TableSet[adminAgent.PO]
}

// IsExists 管理员是否存在
func (repository adminRepository) IsExists(adminName string) bool {
	return repository.Where("UserName = ?", adminName).IsExists()
}

// IsExistsWithoutSelf 管理员是否存在
func (repository adminRepository) IsExistsWithoutSelf(adminName string, adminId int) bool {
	return repository.Where("UserName = ? and Id <> ?", adminName, adminId).IsExists()
}

// Add 添加管理员
func (repository adminRepository) Add(do admin.DomainObject) int {
	po := mapper.Single[adminAgent.PO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改管理员
func (repository adminRepository) Update(id int, do admin.DomainObject) {
	po := mapper.Single[adminAgent.PO](do)
	repository.DbContext.Admin.Update(id, po)
}

// ToList Admin列表
func (repository adminRepository) ToList() []admin.DomainObject {
	lst := repository.DbContext.Admin.ToList()
	return mapper.Array[admin.DomainObject](lst)
}

// ToInfo Admin信息
func (repository adminRepository) ToInfo(id int) admin.DomainObject {
	po := repository.DbContext.Admin.ToInfo(id)
	return mapper.Single[admin.DomainObject](po)
}

// ToInfoByUsername Admin信息
func (repository adminRepository) ToInfoByUsername(username string, pwd string) admin.DomainObject {
	po := repository.DbContext.Admin.ToInfoByUserName(username, pwd)
	return mapper.Single[admin.DomainObject](po)
}

// Count Admin数量
func (repository adminRepository) Count() int64 {
	return repository.DbContext.Admin.Count()
}

// Delete 删除管理员
func (repository adminRepository) Delete(id int) {
	repository.DbContext.Admin.Delete(id)
}
