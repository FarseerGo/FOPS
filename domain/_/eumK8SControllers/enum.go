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
