package device

type IDirectoryDevice interface {
	// Check 检查目录
	Check(progress chan string)
}
