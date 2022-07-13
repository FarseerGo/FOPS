package domain

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Error
	Critical
	NoneLevel
)
