package k8s

import (
	"fops/domain/k8s/pod"
	"fops/domain/metaData/projectGroup"
	"fs/core/container"
	"fs/linq"
)

type PodApp struct {
	repository             pod.Repository
	projectGroupRepository projectGroup.Repository
}

func NewPodApp() *PodApp {
	return &PodApp{
		repository:             container.Resolve[pod.Repository](),
		projectGroupRepository: container.Resolve[projectGroup.Repository](),
	}
}

// ToList 项目列表
func (app *PodApp) ToList(groupId int, clusterId int) []pod.DomainObject {
	var lstProject []pod.DomainObject
	if groupId > 0 {
		lstProject = app.repository.ToListByGroupId(groupId)
	} else {
		lstProject = app.repository.ToList()
	}

	// 筛选了组ID，则要过滤当前组不支持的集群环境
	if groupId > 0 {
		var lstGroup = app.projectGroupRepository.ToList()
		for _, projectGroupDo := range lstGroup {
			// 如果项目组不包含当前选中的集群，则移除项目

			if !linq.FromC(projectGroupDo.ClusterIds).Contains(clusterId) {
				lstProject = linq.From(lstProject).RemoveAll(func(item pod.DomainObject) bool {
					return item.GroupId == projectGroupDo.Id
				})
			}
		}
	}
	return linq.FromOrder[pod.DomainObject, int](lstProject).OrderByDescending(func(item pod.DomainObject) int { return item.Id })
}

// Update 更新模板ID
func (app *PodApp) Update(pod pod.DomainObject) {
	app.repository.Update(pod)
}
