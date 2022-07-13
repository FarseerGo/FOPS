package repository

import (
	"fops/domain/metaData/projectGroup"
	"fs/core/container"
)

func init() {
	// 注册项目组仓储
	container.Register(func() projectGroup.Repository { return &projectGroupRepository{} })
}

type projectGroupRepository struct {
}

// New 初始化对象
func new() projectGroup.Repository {
	return &projectGroupRepository{}
}

// ToList 项目组列表
func (repository *projectGroupRepository) ToList() *[]projectGroup.DomainObject {
	return &[]projectGroup.DomainObject{
		{
			Id:         1,
			ClusterIds: []int{1, 2},
			Name:       "自定义",
		},
	}
}

// ToInfo 项目组信息
func (repository *projectGroupRepository) ToInfo(id int) *projectGroup.DomainObject {
	return &projectGroup.DomainObject{
		Id:         1,
		ClusterIds: []int{1, 2},
		Name:       "自定义"}
}

// Count 项目组数量
func (repository *projectGroupRepository) Count() int {
	return 1
}

// Add 添加项目组
func (repository *projectGroupRepository) Add(do projectGroup.DomainObject) int {
	return 2
}

// Update 修改项目组
func (repository *projectGroupRepository) Update(id int, do projectGroup.DomainObject) {
}

// Delete 删除项目组
func (repository *projectGroupRepository) Delete(id int) {
}
