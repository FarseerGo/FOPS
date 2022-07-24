package admin

import (
	"fops/infrastructure/repository/context"
)

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList Admin列表
func (agent) ToList() []PO {
	return context.NewContext().Admin.Order("Id desc").ToList()
}

// ToInfo Admin信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().Admin.Where("Id = ?", id).ToEntity()
}

// ToInfoByUserName Admin信息
func (agent) ToInfoByUserName(userName string, pwd string) PO {
	return context.NewContext().Admin.Where("UserName = ? && UserPwd = ?", userName, pwd).ToEntity()
}

// Count Admin数量
func (agent) Count() int64 {
	return context.NewContext().Admin.Count()
}

// IsExists 管理员是否存在
func (agent) IsExists(adminName string) bool {
	return context.NewContext().Admin.Where("UserName = ?", adminName).IsExists()
}

// IsExistsByAdminId 管理员是否存在
func (agent) IsExistsByAdminId(adminName string, adminId int) bool {
	return context.NewContext().Admin.Where("UserName = ? and Id <> ?", adminName, adminId).IsExists()
}

// Add 添加管理员
func (agent) Add(po PO) int {
	context.NewContext().Admin.Insert(&po)
	return po.Id
}

// Update 修改管理员
func (agent) Update(id int, po PO) {
	context.NewContext().Admin.Where("Id = ?", id).Update(po)
}

// Delete 删除管理员
func (agent) Delete(id int) {
	context.NewContext().Admin.Where("Id = ?", id).Delete()
}
