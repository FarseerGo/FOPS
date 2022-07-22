package projectGroup

import (
	"fmt"
	_ "fops/infrastructure/repository"
	"testing"
)

func TestToList(t *testing.T) {
	lst := ToList()
	for _, dto := range lst {
		fmt.Println(dto)
	}
}

// ToInfo 项目组信息
func TestToInfo(t *testing.T) {
	info := ToInfo(1)
	fmt.Println(info)
}

// Count 项目组数量
func TestCount(t *testing.T) {
	count := Count()
	fmt.Println(count)
}

// Add 添加项目组
func TestAdd(t *testing.T) {
	Add(DTO{
		Id: 2,
		ClusterIds: []int{
			1, 2,
		},
		Name: "test",
	})
}

// Update 修改项目组
func TestUpdate(t *testing.T) {
	Update(DTO{
		Id: 2,
		ClusterIds: []int{
			1, 2,
		},
		Name: "test",
	})
}

// Delete 删除项目组
func TestDelete(t *testing.T) {
	Delete(1)
}
