package event

import "fs/eventBus"

const DockerPushedEventName = "DockerPushedEvent"

type DockerPushedEvent struct {
	// 构建版本号
	BuildNumber int
	// 项目ID
	ProjectId int
}

// PublishEvent 发布事件
func (receiver DockerPushedEvent) PublishEvent() {
	eventBus.PublishEvent(DockerPushedEventName, receiver)
}
