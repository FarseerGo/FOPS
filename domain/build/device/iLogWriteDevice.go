package device

type ILogWriteDevice interface {
	// View 查看日志
	View(buildId int) []string
	// CreateProgress 生成Progress类，并自动输出日志
	CreateProgress(buildId int) chan string
}
