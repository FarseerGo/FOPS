package event

import "fs/eventBus"

// BuildFinishedEventName 事件名称
const BuildFinishedEventName = "BuildFinishedEvent"

// BuildFinishedEvent 构建完成后，发布事件
type BuildFinishedEvent struct {
	// 项目ID
	ProjectId int
	// 本次失败的构建ID
	BuildId int
	// 构建的集群
	ClusterId int
	// 是否成功
	IsSuccess bool
}

// PublishEvent 发布事件
func (receiver BuildFinishedEvent) PublishEvent() {
	eventBus.PublishEvent(BuildFinishedEventName, receiver)
}
