package repository

import (
	"fops/domain/k8s/pod"
	"fops/infrastructure/repository/agent/projectAgent"
	"fops/infrastructure/repository/agent/yamlTplAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() pod.Repository { return &podRepository{} })
}

type podRepository struct {
}

func (repository podRepository) ToListByGroupId(groupId int) []pod.DomainObject {
	lstYaml := mapper.Array[pod.YamlTplVO](yamlTplAgent.ToList())
	lstProject := projectAgent.ToListByGroupId(groupId)
	var lst []pod.DomainObject

	for _, projectPO := range lstProject {
		var podDO = mapper.Single[pod.DomainObject](projectPO)
		podDO.SetYamlTpl(lstYaml, projectPO.K8STplDeployment, projectPO.K8STplIngress, projectPO.K8STplService, projectPO.K8STplConfig)
		lst = append(lst, podDO)
	}
	return lst
}

func (repository podRepository) ToList() []pod.DomainObject {
	lstYaml := mapper.Array[pod.YamlTplVO](yamlTplAgent.ToList())
	lstProject := projectAgent.ToList()
	var lst []pod.DomainObject

	for _, projectPO := range lstProject {
		var podDO = mapper.Single[pod.DomainObject](projectPO)
		podDO.SetYamlTpl(lstYaml, projectPO.K8STplDeployment, projectPO.K8STplIngress, projectPO.K8STplService, projectPO.K8STplConfig)
		lst = append(lst, podDO)
	}
	return lst
}

func (repository podRepository) Update(pod pod.DomainObject) {
	projectAgent.Update(pod.Id, projectAgent.PO{
		K8STplDeployment: pod.K8STplDeployment.Id,
		K8STplService:    pod.K8STplService.Id,
		K8STplIngress:    pod.K8STplIngress.Id,
		K8STplConfig:     pod.K8STplConfig.Id,
	}, "K8STplDeployment", "K8STplService", "K8STplIngress", "K8STplConfig")
}
