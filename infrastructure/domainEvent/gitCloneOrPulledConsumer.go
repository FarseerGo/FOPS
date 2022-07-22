package domainEvent

import (
	"fops/domain/building/build/event"
	"fops/domain/metaData/git"
	"fs/core/container"
	"fs/eventBus"
	"time"
)

func init() {
	eventBus.Subscribe(event.GitCloneOrPulledEventName, gitCloneOrPulledConsumer)
}

var gitRepository = container.Resolve[git.Repository]()

func gitCloneOrPulledConsumer(message any, ea eventBus.EventArgs) {
	gitCloneOrPulledEvent := message.(event.GitCloneOrPulledEvent)

	// 更新git拉取时间
	gitRepository.UpdateForTime(gitCloneOrPulledEvent.GitId, time.UnixMicro(ea.CreateAt))
}
