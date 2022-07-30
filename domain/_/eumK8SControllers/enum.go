package eumK8SControllers

type Enum int

const (
	// Deployment 无状态应用
	Deployment Enum = iota
	// StatefulSet 有状态应用
	StatefulSet
	// DaemonSet 所有节点都会运行一个实例
	DaemonSet
	// Job 一次性任务
	Job
	// Cronjob 定时任务
	Cronjob
)

// GetName 获取标签名称
func GetName(eum Enum) string {
	switch eum {
	case Deployment:
		return "deployment"
	case StatefulSet:
		return "statefulSet"
	case DaemonSet:
		return "daemonSet"
	case Job:
		return "job"
	case Cronjob:
		return "cronjob"
	}
	return ""
}
