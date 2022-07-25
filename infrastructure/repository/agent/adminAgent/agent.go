package adminAgent

import (
	"fs/data"
)

type TableSet struct {
	data.TableSet[PO]
}

// ToList Admin列表
func (set TableSet) ToList() []PO {
	return set.Order("Id desc").ToList()
}

// ToInfo Admin信息
func (set TableSet) ToInfo(id int) PO {
	return set.Where("Id = ?", id).ToEntity()
}

// ToInfoByUserName Admin信息
func (set TableSet) ToInfoByUserName(userName string, pwd string) PO {
	return set.Where("UserName = ? && UserPwd = ?", userName, pwd).ToEntity()
}

// Count Admin数量
func (set TableSet) Count() int64 {
	return set.Count()
}

// IsExists 管理员是否存在
func (set TableSet) IsExists(adminName string) bool {
	return set.Where("UserName = ?", adminName).IsExists()
}

// IsExistsByAdminId 管理员是否存在
func (set TableSet) IsExistsByAdminId(adminName string, adminId int) bool {
	return set.Where("UserName = ? and Id <> ?", adminName, adminId).IsExists()
}

// Add 添加管理员
func (set TableSet) Add(po PO) int {
	set.Insert(&po)
	return po.Id
}

// Update 修改管理员
func (set TableSet) Update(id int, po PO) {
	set.Where("Id = ?", id).Update(po)
}

// Delete 删除管理员
func (set TableSet) Delete(id int) {
	set.Where("Id = ?", id).Delete()
}
