package projectGroupAgent

import "fops/infrastructure/repository/context"

// ToList 项目组列表
func ToList() []PO {
	return context.NewContext().ProjectGroup.Desc("Id").ToList()
}

// ToInfo 项目组信息
func ToInfo(id int) PO {
	return context.NewContext().ProjectGroup.Where("Id = ?", id).ToEntity()
}

// Count 项目组数量
func Count() int64 {
	return context.NewContext().ProjectGroup.Count()
}

// Add 添加项目组
func Add(po PO) int {
	context.NewContext().ProjectGroup.Insert(&po)
	return po.Id
}

// Update 修改项目组
func Update(id int, po PO) {
	context.NewContext().ProjectGroup.Where("Id = ?", id).Update(po)
}

// Delete 删除项目组
func Delete(id int) {
	context.NewContext().ProjectGroup.Where("Id = ?", id).Delete()
}
