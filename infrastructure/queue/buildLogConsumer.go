package queue

import (
	"fmt"
	"github.com/farseernet/farseer.go/linq"
	"github.com/farseernet/farseer.go/mq/queue"
)

func init() {
	queue.Subscribe("BuildLog", "", 1000, buildLogConsumer)
}

// 订阅日志，写入到log文件
func buildLogConsumer(subscribeName string, message []any, remainingCount int) {

	lstBuildLogVO := linq.FromT[any, BuildLogVO](message).Select(func(item any) BuildLogVO {
		return item.(BuildLogVO)
	})
	fmt.Println(lstBuildLogVO)
}
