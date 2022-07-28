package build

import (
	"fops/application/building/buildProject"
	"fops/domain/building/build"
	"fops/domain/building/build/event"
	"fops/domain/building/build/vo"
	"fops/domain/k8s/cluster"
	"fops/domain/metaData/project"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/mapper"

	"strconv"
)

type app struct {
	repository        build.Repository
	projectRepository project.Repository
	clusterRepository cluster.Repository
}

func NewApp() *app {
	return &app{
		repository:        container.Resolve[build.Repository](),
		projectRepository: container.Resolve[project.Repository](),
		clusterRepository: container.Resolve[cluster.Repository]()}
}

// Build 构建
func (app *app) Build() {
	//BuildService.Build()
}

// Success 设置任务成功
func (app *app) Success(clusterId int, project buildProject.Dto) {
	// 构建ID没有传的时候，通过版本号获取
	buildNumber, _ := strconv.Atoi(project.DockerVer)
	buildId := app.repository.GetBuildId(project.Id, buildNumber)

	// 发布事件
	event.BuildFinishedEvent{ProjectId: project.Id, BuildId: buildId, ClusterId: clusterId, IsSuccess: true}.PublishEvent()
}

// Cancel 主动取消任务
func (app *app) Cancel(id int) {
	app.repository.Cancel(id)
}

// ToBuildingList 获取构建队列前30
func (app *app) ToBuildingList(pageSize int, pageIndex int) []Dto {
	lstDo := app.repository.ToBuildingList(pageSize, pageIndex)
	return mapper.Array[Dto](lstDo)
}

// Count 当前构建的队列数量
func (app *app) Count() int64 {
	return app.repository.Count()
}

// Add 添加构建任务
// todo [Transaction("fops")]
func (app *app) Add(projectId int, clusterId int) int {
	// 获取最后一个编译版本号
	buildNumber := app.repository.GetBuildNumber(projectId)

	// 项目
	projectDO := app.projectRepository.ToInfo(projectId)
	if projectDO.Id == 0 {
		panic("项目ID={projectId}，不存在")
	}
	projectVO := mapper.Single[vo.ProjectVO](projectDO)

	// 集群
	clusterDO := app.clusterRepository.ToInfo(clusterId)
	clusterVO := mapper.Single[vo.ClusterVO](clusterDO)
	if clusterVO.Id < 1 {
		panic("集群ID={clusterId}，不存在")
	}
	buildDO := build.NewDO2(buildNumber, projectVO, clusterVO)

	return app.repository.Add(buildDO)
}

// AddBatch 添加构建任务
func (app *app) AddBatch(lst []Dto, clusterId int) {
	// todo 事务
	//using (var transaction = repository.BeginTransaction())
	//{
	for _, projectDTO := range lst {
		app.Add(projectDTO.Id, clusterId)
	}
	//transaction.SaveChanges()
	//}
}

// ToInfo 查看构建信息
func (app *app) ToInfo(id int) Dto {
	info := app.repository.ToInfo(id)
	return mapper.Single[Dto](info)
}
