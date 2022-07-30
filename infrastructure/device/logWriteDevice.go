package device

import (
	"fops/domain/building/devicer"
	"fops/infrastructure/localQueue"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/mq/queue"
	"github.com/farseernet/farseer.go/utils/file"
	"strconv"
	"time"
)

const SavePath = "/var/lib/fops/log/"

func init() {
	_ = container.Register(func() devicer.ILogWriteDevice { return &logWriteDevice{} })
}

type logWriteDevice struct {
}

func (logWriteDevice) View(buildId int) []string {
	path := SavePath + strconv.Itoa(buildId) + ".txt"
	if file.IsExists(path) {
		return file.ReadAllLines(path)
	}
	return []string{}
}

func (logWriteDevice) CreateProgress(buildId int) chan string {
	// 清除历史记录（正常不会存在，当buildId被重置时，有可能会冲突）
	clear(buildId)

	progress := make(chan string, 1000)
	go sendLog(buildId, progress)
	return progress
}

// GenerateFilename 生成文件名
func GenerateFilename(buildId int) string {
	logfile := SavePath + strconv.Itoa(buildId) + ".txt"
	if !file.IsExists(SavePath) {
		file.CreateDir766(SavePath)
	}
	return logfile
}

// 清除历史记录（正常不会存在，当buildId被重置时，有可能会冲突）
func clear(buildId int) {
	path := SavePath + strconv.Itoa(buildId) + ".txt"
	if file.IsExists(path) {
		file.Delete(path)
	}
}

func sendLog(buildId int, progress chan string) {
	for log := range progress {
		queue.Push("BuildLog", localQueue.BuildLogVO{
			BuildId: buildId,
			Log:     log,
			LogAt:   time.Now(),
		})
	}
}
