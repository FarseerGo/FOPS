package repository

import (
	"fops/domain/_/eumBuildStatus"
	"fops/domain/building/build"
	"fops/domain/building/build/vo"
	"fops/domain/building/device"
	"fops/domain/k8s/cluster"
	"fops/domain/metaData/dockerHub"
	"fops/domain/metaData/dockerfileTpl"
	"fops/domain/metaData/git"
	"fops/domain/metaData/project"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
	"time"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() build.Repository { return &buildRepository{data.Init[context.MysqlContext]("fops").Build} })
}

type buildRepository struct {
	data.TableSet[model.BuildPO]
}

// GetBuildNumber 获取构建的编号
func (repository buildRepository) GetBuildNumber(projectId int) int {
	return repository.Where("ProjectId = ?", projectId).Order("Id desc").GetInt("BuildNumber")
}

// GetBuildId 获取构建任务的主键
func (repository buildRepository) GetBuildId(projectId int, buildNumber int) int {
	return repository.Where("BuildNumber = ? and ProjectId = ?", buildNumber, projectId).GetInt("Id")
}

// Add 添加构建任务
func (repository buildRepository) Add(do build.DomainObject) int {
	po := mapper.Single[model.BuildPO](do)
	repository.Insert(&po)
	return po.Id
}

// Count 当前构建的队列数量
func (repository buildRepository) Count() int64 {
	return repository.Where("Status <> ?", eumBuildStatus.Finish).Count()
}

// ToBuildingList 获取构建队列前30
func (repository buildRepository) ToBuildingList(pageSize int, pageIndex int) []build.DomainObject {
	pageList := repository.Select("Id", "Status", "BuildNumber", "IsSuccess", "ProjectId", "ProjectName", "CreateAt", "FinishAt", "ClusterId").Order("Id desc").ToPageList(pageSize, pageIndex)
	return mapper.Array[build.DomainObject](pageList.List)
}

// ToInfo 查看构建信息 [Transaction(typeof(MysqlContext))]
func (repository buildRepository) ToInfo(id int) build.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	do := mapper.Single[build.DomainObject](po)
	if do.Id > 0 {

		projectPO := container.Resolve[project.Repository]().ToInfo(do.Project.Id)
		clusterPO := container.Resolve[cluster.Repository]().ToInfo(do.Cluster.Id)
		dockerPO := container.Resolve[dockerHub.Repository]().ToInfo(projectPO.DockerHub)
		projectVO := mapper.Single[vo.ProjectVO](projectPO)
		clusterVO := mapper.Single[vo.ClusterVO](clusterPO)
		dockerVO := mapper.Single[vo.DockerVO](dockerPO)
		do.Set(projectVO, dockerVO, clusterVO, nil)
	}
	return do
}

// GetUnBuildInfo 获取未构建的任务
func (repository buildRepository) GetUnBuildInfo() build.DomainObject {
	po := repository.Where("Status = ? and BuildServerId = ?", eumBuildStatus.None, fs.AppId).Asc("Id").ToEntity()
	do := mapper.Single[build.DomainObject](po)

	if do.Id > 0 {
		projectPO := container.Resolve[project.Repository]().ToInfo(do.Project.Id)
		clusterPO := container.Resolve[cluster.Repository]().ToInfo(do.Cluster.Id)
		dockerPO := container.Resolve[dockerHub.Repository]().ToInfo(projectPO.DockerHub)
		projectVO := mapper.Single[vo.ProjectVO](projectPO)
		clusterVO := mapper.Single[vo.ClusterVO](clusterPO)
		dockerVO := mapper.Single[vo.DockerVO](dockerPO)

		if projectPO.DockerfileTpl > 0 {
			dockerfileTplPO := container.Resolve[dockerfileTpl.Repository]().ToInfo(projectPO.DockerfileTpl)
			dockerVO.DockerfileContent = dockerfileTplPO.Template
		}

		lstGitIds := append([]int{projectPO.GitId}, projectPO.DependentGitIds...)
		lstGitPO := container.Resolve[git.Repository]().ToListByIds(lstGitIds)
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
	return repository.Where("Id = ?", id).Select("Status", "CreateAt").Update(model.BuildPO{
		Status:   eumBuildStatus.Building,
		CreateAt: time.Now(),
	})
}

// Success 任务完成
func (repository buildRepository) Success(id int) {
	repository.Where("Id = ?", id).Select("Status", "IsSuccess", "CreateAt").Update(model.BuildPO{
		Status:    eumBuildStatus.Finish,
		IsSuccess: true,
		FinishAt:  time.Now(),
	})
}

// Cancel 主动取消任务
func (repository buildRepository) Cancel(id int) {
	repository.Where("Id = ?", id).Select("Status", "IsSuccess", "CreateAt").Update(model.BuildPO{
		Status:    eumBuildStatus.Finish,
		IsSuccess: false,
		FinishAt:  time.Now(),
	})
}

// GetStatus 获取构建状态
func (repository buildRepository) GetStatus(id int) eumBuildStatus.Enum {
	return eumBuildStatus.Enum(repository.Where("Id = ?", id).GetInt("Status"))
}
