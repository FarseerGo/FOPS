package eumRuntimeEnv

type Enum int

const (
	// Dev 开发环境
	Dev Enum = iota
	// Test 测试环境
	Test
	// PreRelease 预发布
	PreRelease
	// Prod 生产环境
	Prod
)
