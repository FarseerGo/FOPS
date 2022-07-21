package build

import "fs/eventBus"

// FinishedEventName 事件名称
const FinishedEventName = "BuildFinishedEvent"

// FinishedEvent 构建完成后，发布事件
type FinishedEvent struct {
	// 项目ID
	ProjectId int
	// 本次失败的构建ID
	BuildId int
	// 构建的集群
	ClusterId int
	// 是否成功
	IsSuccess bool
}

func (receiver FinishedEvent) PublishEvent() {
	eventBus.PublishEvent(FinishedEventName, receiver)
}
