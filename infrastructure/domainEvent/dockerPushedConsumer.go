package domainEvent

import (
	"fops/domain/building/build/event"
	"github.com/farseernet/farseer.go/eventBus"
	"github.com/farseernet/farseer.go/utils/parse"
)

func init() {
	eventBus.Subscribe(event.DockerPushedEventName, dockerPushedConsumer)
}

// docker推送完成事件
func dockerPushedConsumer(message any, ea eventBus.EventArgs) {
	dockerPushedEvent := message.(event.DockerPushedEvent)

	// 更新项目的版本信息
	projectRepository.UpdateDockerVer(dockerPushedEvent.ProjectId, parse.Convert(dockerPushedEvent.BuildNumber, ""))
}
