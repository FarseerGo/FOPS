package projectGroupApp

import (
	"fops/domain/metaData/projectGroup"
	"fs/core/container"
	"fs/mapper"
)

var repository, _ = container.Resolve[projectGroup.Repository]()

// ToList 项目组列表
func ToList() []DTO {
	lstDo := repository.ToList()
	return mapper.MapperArray[[]DTO](lstDo)
}

// ToInfo 项目组信息
func ToInfo(id int) DTO {
	do := repository.ToInfo(id)
	return mapper.AutoMapper[DTO](do)
}

// Count 项目组数量
func Count() int {
	return repository.Count()
}

// Add 添加项目组
func Add(dto DTO) {
	do := mapper.AutoMapper[projectGroup.DomainObject](dto)
	repository.Add(do)
}

// Update 修改项目组
func Update(dto DTO) {
	do := mapper.AutoMapper[projectGroup.DomainObject](dto)
	repository.Update(dto.Id, do)
}

// Delete 删除项目组
func Delete(id int) {
	repository.Delete(id)
}
