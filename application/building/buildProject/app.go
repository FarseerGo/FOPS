package buildProject

import (
	domain "fops/domain/_"
	"fops/domain/metaData/dockerfileTpl"
	"fops/domain/metaData/project"
	"fops/domain/metaData/projectGroup"
	"fs/core/container"
	"fs/linq"
	"fs/mapper"
)

type app struct {
	projectRepository       project.Repository
	projectGroupRepository  projectGroup.Repository
	dockerfileTplRepository dockerfileTpl.Repository
}

func NewApp() *app {
	return &app{
		projectRepository:       container.Resolve[project.Repository](),
		projectGroupRepository:  container.Resolve[projectGroup.Repository](),
		dockerfileTplRepository: container.Resolve[dockerfileTpl.Repository](),
	}
}

// ToList 项目列表
func (app *app) ToList(groupId int, clusterId int) []Dto {
	var lstProject []project.DomainObject
	if groupId > 0 {
		lstProject = app.projectRepository.ToListByGroupId(groupId)
	} else {
		lstProject = app.projectRepository.ToList()
	}

	// 筛选了组ID，则要过滤当前组不支持的集群环境
	if groupId > 0 {
		var lstGroup = app.projectGroupRepository.ToList()
		for _, item := range lstGroup {
			// 如果项目组不包含当前选中的集群，则移除项目
			if !linq.FromC(item.ClusterIds).Contains(clusterId) {
				lstProject = linq.From(lstProject).RemoveAll(func(item project.DomainObject) bool {
					return item.GroupId == item.Id
				})
			}
		}
	}

	// 初始化集群默认值
	lstNoExistsCluster := linq.From(lstProject).FindAll(func(project project.DomainObject) bool {
		return !linq.Dictionary(project.ClusterVer).ExistsKey(clusterId)
	})

	for _, project := range lstNoExistsCluster {
		project.ClusterVer[clusterId] = &domain.ClusterVerVO{
			DockerVer: "0",
		}
	}

	// 设置Docker模板名称
	lstDTO := mapper.Array[Dto](lstProject)
	lstDTO = linq.FromOrder[Dto, int](lstDTO).OrderByDescending(func(item Dto) int {
		return item.Id
	})

	lstDockerfile := app.dockerfileTplRepository.ToList()

	for _, buildProjectDTO := range lstDTO {
		dockerfile := linq.From(lstDockerfile).Find(func(item dockerfileTpl.DomainObject) bool {
			return item.Id == buildProjectDTO.DockerfileTpl
		})
		if dockerfile.Id > 0 {
			buildProjectDTO.DockerfileName = dockerfile.Name
		}
	}

	return lstDTO
}
