package repository

import (
	"fops/domain/_/eumBuildStatus"
	"fops/domain/building/build"
	"fops/domain/building/build/vo"
	"fops/domain/building/device"
	"fops/infrastructure/repository/agent/buildAgent"
	"fops/infrastructure/repository/agent/clusterAgent"
	"fops/infrastructure/repository/agent/dockerHubAgent"
	"fops/infrastructure/repository/agent/dockerfileTplAgent"
	"fops/infrastructure/repository/agent/gitAgent"
	"fops/infrastructure/repository/agent/projectAgent"
	"fs/core/container"
	"fs/mapper"
	"time"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() build.Repository { return &buildRepository{} })
}

type buildRepository struct {
}

// GetBuildNumber 获取构建的编号
func (repository buildRepository) GetBuildNumber(projectId int) int {
	return buildAgent.GetBuildNumber(projectId)
}

// GetBuildId 获取构建任务的主键
func (repository buildRepository) GetBuildId(projectId int, buildNumber int) int {
	return buildAgent.GetBuildId(projectId, buildNumber)
}

// Add 添加构建任务
func (repository buildRepository) Add(do build.DomainObject) int {
	po := mapper.Single[buildAgent.PO](do)
	return buildAgent.Add(po)
}

// Cancel 主动取消任务
func (repository buildRepository) Cancel(id int) {
	buildAgent.Update(id, buildAgent.PO{
		Status:    eumBuildStatus.Finish,
		IsSuccess: false,
		FinishAt:  time.Now(),
	})
}

// Count 当前构建的队列数量
func (repository buildRepository) Count() int64 {
	return buildAgent.Count()
}

// ToBuildingList 获取构建队列前30
func (repository buildRepository) ToBuildingList(pageSize int, pageIndex int) []build.DomainObject {
	pageList := buildAgent.ToBuildingList(pageSize, pageIndex)
	return mapper.Array[build.DomainObject](pageList.List)
}

// ToInfo 查看构建信息 [Transaction(typeof(MysqlContext))]
func (repository buildRepository) ToInfo(id int) build.DomainObject {

	po := buildAgent.ToInfo(id)
	do := mapper.Single[build.DomainObject](po)
	if do.Id > 0 {

		projectPO := projectAgent.ToInfo(do.Project.Id)
		clusterPO := clusterAgent.ToInfo(do.Cluster.Id)
		dockerPO := dockerHubAgent.ToInfo(projectPO.DockerHub)
		projectVO := mapper.Single[vo.ProjectVO](projectPO)
		clusterVO := mapper.Single[vo.ClusterVO](clusterPO)
		dockerVO := mapper.Single[vo.DockerVO](dockerPO)
		do.Set(projectVO, dockerVO, clusterVO, nil)
	}
	return do
}

// GetUnBuildInfo 获取未构建的任务
func (repository buildRepository) GetUnBuildInfo() build.DomainObject {
	po := buildAgent.ToUnBuildInfo()
	do := mapper.Single[build.DomainObject](po)

	if do.Id > 0 {
		projectPO := projectAgent.ToInfo(do.Project.Id)
		clusterPO := clusterAgent.ToInfo(do.Cluster.Id)
		dockerPO := dockerHubAgent.ToInfo(projectPO.DockerHub)
		projectVO := mapper.Single[vo.ProjectVO](projectPO)
		clusterVO := mapper.Single[vo.ClusterVO](clusterPO)
		dockerVO := mapper.Single[vo.DockerVO](dockerPO)

		if projectPO.DockerfileTpl > 0 {
			dockerfileTplPO := dockerfileTplAgent.ToInfo(projectPO.DockerfileTpl)
			dockerVO.DockerfileContent = dockerfileTplPO.Template
		}

		lstGitIds := append([]int{projectPO.GitId}, projectPO.DependentGitIds...)
		lstGitPO := gitAgent.ToListByIds(lstGitIds)
		lstGit := mapper.Array[vo.GitVO](lstGitPO)

		gitDevice := container.Resolve[device.IGitDevice]()
		for index := 0; index < len(lstGit); index++ {
			lstGit[index].ProjectPath = gitDevice.GetGitPath(lstGit[index].Hub)
			lstGit[index].IsMaster = lstGit[index].Id == projectPO.GitId // 标记是否为主仓库
		}

		do.Set(projectVO, dockerVO, clusterVO, lstGit)
	}

	return do
}

// SetBuilding 设置任务为构建中
func (repository buildRepository) SetBuilding(id int) int64 {
	return buildAgent.Update(id, buildAgent.PO{
		Status:   eumBuildStatus.Building,
		CreateAt: time.Now(),
	})
}

// Success 任务完成
func (repository buildRepository) Success(id int) {
	buildAgent.Update(id, buildAgent.PO{
		Status:    eumBuildStatus.Finish,
		IsSuccess: true,
		FinishAt:  time.Now(),
	})
}

// GetStatus 获取构建状态
func (repository buildRepository) GetStatus(id int) eumBuildStatus.Enum {
	return buildAgent.GetStatus(id)
}
