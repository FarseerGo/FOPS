package domainEvent

import (
	"fops/domain/building/build/event"
	"fops/domain/metaData/project"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/eventBus"
)

func init() {
	eventBus.Subscribe(event.BuildFinishedEventName, buildFinishedConsumer)
}

var projectRepository = container.Resolve[project.Repository]()

// buildFinishedConsumer 构建完成事件
func buildFinishedConsumer(message any, _ eventBus.EventArgs) {

	buildFailedEvent := message.(event.BuildFinishedEvent)
	projectDo := projectRepository.ToInfo(buildFailedEvent.ProjectId)
	projectDo.UpdateBuildVer(buildFailedEvent.IsSuccess, buildFailedEvent.ClusterId, buildFailedEvent.BuildId)

	// 更新项目的版本信息
	projectRepository.UpdateClusterVer(buildFailedEvent.ProjectId, projectDo.ClusterVer)
}
