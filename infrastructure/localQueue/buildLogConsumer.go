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
	var buildLogVos []BuildLogVO
	linq.From(message).Select(&buildLogVos, func(item any) any {
		return item.(BuildLogVO)
	})

	// 按BuildId分组
	var mapBuildLog map[int][]BuildLogVO
	linq.From(buildLogVos).GroupBy(&mapBuildLog, func(item BuildLogVO) any {
		return item.BuildId
	})

	for buildId, items := range mapBuildLog {
		// 创建日志文件
		logfile := device.GenerateFilename(buildId)
		var logs []string
		linq.From(items).Select(&logs, func(item BuildLogVO) any {
			return str.ToDateTime(item.LogAt) + " " + item.Log
		})
		// 写入日志文件
		file.AppendAllLine(logfile, logs)
	}
}
