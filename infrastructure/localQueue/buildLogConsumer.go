package localQueue

import (
	"fops/infrastructure/device"
	"github.com/farseernet/farseer.go/linq"
	"github.com/farseernet/farseer.go/mq/queue"
	"github.com/farseernet/farseer.go/utils/file"
	"github.com/farseernet/farseer.go/utils/str"
)

func init() {
	queue.Subscribe("BuildLog", "", 1000, buildLogConsumer)
}

// 订阅日志，写入到log文件
func buildLogConsumer(subscribeName string, message []any, remainingCount int) {
	// 转成BuildLogVO数组
	buildLogVos := linq.FromT[any, BuildLogVO](message).Select(func(item any) BuildLogVO {
		return item.(BuildLogVO)
	})

	// 按BuildId分组
	mapBuildLog := linq.FormGroupBy[BuildLogVO, int](buildLogVos).GroupBy(func(item BuildLogVO) int {
		return item.BuildId
	})

	for buildId, items := range mapBuildLog {

		// 创建日志文件
		logfile := device.GenerateFilename(buildId)
		logs := linq.FromT[BuildLogVO, string](items).Select(func(item BuildLogVO) string {
			return str.ToDateTime(item.LogAt) + " " + item.Log
		})
		// 写入日志文件
		file.AppendAllLine(logfile, logs)
	}
}
