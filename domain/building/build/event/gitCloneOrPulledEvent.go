package event

import "fs/eventBus"

const GitCloneOrPulledEventName = "GitCloneOrPulledEvent"

type GitCloneOrPulledEvent struct {
	// GitId
	GitId int
}

// PublishEvent 发布事件
func (receiver GitCloneOrPulledEvent) PublishEvent() {
	eventBus.PublishEvent(GitCloneOrPulledEventName, receiver)
}
