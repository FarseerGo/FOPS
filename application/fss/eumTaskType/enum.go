package eumTaskType

type Enum int

const (
	// None 未开始
	None Enum = iota
	// Scheduler 已调度
	Scheduler
	// Working 执行中
	Working
	// Fail 失败
	Fail
	// Success 完成
	Success
	// ReScheduler 重新调度
	ReScheduler
)
