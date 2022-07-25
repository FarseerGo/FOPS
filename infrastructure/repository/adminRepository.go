package repository

import (
	"fops/domain/security/admin"
	"fops/infrastructure/repository/agent/adminAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() admin.Repository {
		return &adminRepository{}
	})
}

type adminRepository struct {
}

// IsExists 管理员是否存在
func (repository adminRepository) IsExists(adminName string) bool {
	return adminAgent.IsExists(adminName)
}

// IsExistsWithoutSelf 管理员是否存在
func (repository adminRepository) IsExistsWithoutSelf(adminName string, adminId int) bool {
	return adminAgent.IsExistsByAdminId(adminName, adminId)
}

// Add 添加管理员
func (repository adminRepository) Add(do admin.DomainObject) int {
	po := mapper.Single[adminAgent.PO](do)
	return adminAgent.Add(po)
}

// Update 修改管理员
func (repository adminRepository) Update(id int, do admin.DomainObject) {
	po := mapper.Single[adminAgent.PO](do)
	adminAgent.Update(id, po)
}

// ToList Admin列表
func (repository adminRepository) ToList() []admin.DomainObject {
	lst := adminAgent.ToList()
	return mapper.Array[admin.DomainObject](lst)
}

// ToInfo Admin信息
func (repository adminRepository) ToInfo(id int) admin.DomainObject {
	po := adminAgent.ToInfo(id)
	return mapper.Single[admin.DomainObject](po)
}

// ToInfoByUsername Admin信息
func (repository adminRepository) ToInfoByUsername(username string, pwd string) admin.DomainObject {
	po := adminAgent.ToInfoByUserName(username, pwd)
	return mapper.Single[admin.DomainObject](po)
}

// Count Admin数量
func (repository adminRepository) Count() int64 {
	return adminAgent.Count()
}

// Delete 删除管理员
func (repository adminRepository) Delete(id int) {
	adminAgent.Delete(id)
}
