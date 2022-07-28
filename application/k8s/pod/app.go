package pod

import (
	"fops/domain/k8s/pod"
	"fops/domain/metaData/projectGroup"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/linq"
)

type app struct {
	repository             pod.Repository
	projectGroupRepository projectGroup.Repository
}

func NewApp() *app {
	return &app{
		repository:             container.Resolve[pod.Repository](),
		projectGroupRepository: container.Resolve[projectGroup.Repository](),
	}
}

// ToList 项目列表
func (app *app) ToList(groupId int, clusterId int) []pod.DomainObject {
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
func (app *app) Update(pod pod.DomainObject) {
	app.repository.Update(pod)
}
