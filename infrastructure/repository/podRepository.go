package repository

import (
	"fops/domain/k8s/pod"
	"fops/domain/k8s/yamlTpl"
	"fops/domain/metaData/project"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() pod.Repository { return &podRepository{} })
}

type podRepository struct {
}

func (repository podRepository) ToListByGroupId(groupId int) []pod.DomainObject {
	lstYaml := mapper.Array[pod.YamlTplVO](container.Resolve[yamlTpl.Repository]().ToList())
	lstProject := container.Resolve[project.Repository]().ToListByGroupId(groupId)
	var lst []pod.DomainObject

	for _, projectPO := range lstProject {
		var podDO = mapper.Single[pod.DomainObject](projectPO)
		podDO.SetYamlTpl(lstYaml, projectPO.K8STplDeployment, projectPO.K8STplIngress, projectPO.K8STplService, projectPO.K8STplConfig)
		lst = append(lst, podDO)
	}
	return lst
}

func (repository podRepository) ToList() []pod.DomainObject {
	lstYaml := mapper.Array[pod.YamlTplVO](container.Resolve[yamlTpl.Repository]().ToList())
	lstProject := container.Resolve[project.Repository]().ToList()
	var lst []pod.DomainObject

	for _, projectPO := range lstProject {
		var podDO = mapper.Single[pod.DomainObject](projectPO)
		podDO.SetYamlTpl(lstYaml, projectPO.K8STplDeployment, projectPO.K8STplIngress, projectPO.K8STplService, projectPO.K8STplConfig)
		lst = append(lst, podDO)
	}
	return lst
}

func (repository podRepository) Update(pod pod.DomainObject) {
	container.Resolve[project.Repository]().UpdateYamlId(pod.Id, pod.K8STplDeployment.Id, pod.K8STplService.Id, pod.K8STplIngress.Id, pod.K8STplConfig.Id)
}
