package admin

import (
	"fops/infrastructure/repository/dbContext"
)

type agent struct {
}

func NewAgent() agent {
	return agent{}
}

// ToList Admin列表
func (agent) ToList() []PO {
	return dbContext.NewContext().Admin.Order("Id desc").ToList()
}

// ToInfo Admin信息
func ToInfo(id int) PO {
	return dbContext.NewContext().Admin.Where("Id = ?", id).ToEntity()
}

// ToInfoByUserName Admin信息
func ToInfoByUserName(userName string, pwd string) PO {
	return dbContext.NewContext().Admin.Where("UserName = ? && UserPwd = ?", userName, pwd).ToEntity()
}

// Count Admin数量
func Count() int64 {
	return dbContext.NewContext().Admin.Count()
}

// IsExists 管理员是否存在
func IsExists(adminName string) bool {
	return dbContext.NewContext().Admin.Where("UserName = ?", adminName).IsExists()
}

// IsExistsByAdminId 管理员是否存在
func IsExistsByAdminId(adminName string, adminId int) bool {
	return dbContext.NewContext().Admin.Where("UserName = ? and Id <> ?", adminName, adminId).IsExists()
}

// Add 添加管理员
func Add(po PO) int {
	dbContext.NewContext().Admin.Insert(po)
	return po.Id
}

// Update 修改管理员
func Update(id int, po PO) {
	dbContext.NewContext().Admin.Where("Id = ?", id).Update(po)
}

// Delete 删除管理员
func Delete(id int) {
	dbContext.NewContext().Admin.Where("Id = ?", id).Delete()
}
