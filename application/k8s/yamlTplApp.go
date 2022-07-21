package k8s

import (
	"fops/domain/k8s/yamlTpl"
	"fs/core/container"
	"fs/mapper"
)

type YamlTplApp struct {
	repository yamlTpl.Repository
}

func NewYamlTplApp() *YamlTplApp {
	return &YamlTplApp{
		repository: container.Resolve[yamlTpl.Repository](),
	}
}

// ToList Yaml模板列表
func (app *YamlTplApp) ToList() []YamlTplDto {
	lst := app.repository.ToList()
	return mapper.Array[YamlTplDto](lst)
}

// Add 添加Yaml模板
func (app *YamlTplApp) Add(dto YamlTplDto) {
	do := mapper.Single[yamlTpl.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改Yaml模板
func (app *YamlTplApp) Update(dto YamlTplDto) {
	do := mapper.Single[yamlTpl.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToInfo Yaml模板信息
func (app *YamlTplApp) ToInfo(id int) YamlTplDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[YamlTplDto](do)
}

// Count Yaml模板数量
func (app *YamlTplApp) Count() int {
	return app.repository.Count()
}
