package repository

import (
	"fops/domain/security/admin"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() admin.Repository {
		return &adminRepository{
			data.Init[context.MysqlContext]("fops").Admin,
		}
	})
}

type adminRepository struct {
	data.TableSet[model.AdminPO]
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
	po := mapper.Single[model.AdminPO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改管理员
func (repository adminRepository) Update(id int, do admin.DomainObject) {
	po := mapper.Single[model.AdminPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// ToList Admin列表
func (repository adminRepository) ToList() []admin.DomainObject {
	lst := repository.Order("Id desc").ToList()
	return mapper.Array[admin.DomainObject](lst)
}

// ToInfo Admin信息
func (repository adminRepository) ToInfo(id int) admin.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[admin.DomainObject](po)
}

// ToInfoByUsername Admin信息
func (repository adminRepository) ToInfoByUsername(userName string, pwd string) admin.DomainObject {
	po := repository.Where("UserName = ? && UserPwd = ?", userName, pwd).ToEntity()
	return mapper.Single[admin.DomainObject](po)
}

//// Count Admin数量
//func (repository adminRepository) Count() int64 {
//	return repository.Count()
//}

// Delete 删除管理员
func (repository adminRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
