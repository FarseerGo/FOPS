package eumBuildStatus

type Enum int

const (
	// None 未开始
	None Enum = iota
	// Building 构建中
	Building
	// Finish 完成
	Finish
)
