package eumBuildType

type Enum int

const (
	// DotnetPublish dotnet publish
	DotnetPublish Enum = iota
	// Shell Shell脚本
	Shell
	// UnBuild 不用构建
	UnBuild
)
