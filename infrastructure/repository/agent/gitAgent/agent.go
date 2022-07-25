package gitAgent

import (
	"fops/infrastructure/repository/context"
	"time"
)

// ToList Git列表
func ToList() []PO {
	return context.NewContext().Git.Desc("Id").ToList()
}

// ToListByIds Git列表
func ToListByIds(ids []int) []PO {
	return context.NewContext().Git.Where("Id in ?", ids).ToList()
}

// ToInfo Git信息
func ToInfo(id int) PO {
	return context.NewContext().Git.Where("Id = ?", id).ToEntity()
}

// Count Git数量
func Count() int64 {
	return context.NewContext().Git.Count()
}

// Add 添加GIT
func Add(po PO) int {
	context.NewContext().Git.Insert(&po)
	return po.Id
}

// Update 修改GIT
func Update(id int, po PO) {
	context.NewContext().Git.Where("Id = ?", id).Update(po)
}

// UpdatePullAt 修改GIT的拉取时间
func UpdatePullAt(id int, pullAt time.Time) {
	context.NewContext().Git.Where("Id = ?", id).UpdateValue("PullAt", pullAt)
}

// Delete 删除GIT
func Delete(id int) {
	context.NewContext().Git.Where("Id = ?", id).Delete()
}
