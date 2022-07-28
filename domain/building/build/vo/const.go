package vo

const (
	// FopsRoot Fops根目录
	FopsRoot = "/var/lib/fops/"
	// KubeRoot kubectlConfig配置
	KubeRoot = "/var/lib/fops/kube/"
	// NpmModulesRoot NpmModules
	NpmModulesRoot = "/var/lib/fops/npm"
	// DistRoot 编译保存的根目录
	DistRoot = "/var/lib/fops/dist/"
	// GitRoot GIT根目录
	GitRoot = "/var/lib/fops/git/"
	// DockerfilePath Dockerfile文件地址
	DockerfilePath = "/var/lib/fops/dist/Dockerfile"
	// DockerIgnorePath DockerIgnore文件地址
	DockerIgnorePath = "/var/lib/fops/dist/.dockerignore"
	// ShellRoot 生成Shell脚本的存放路径
	ShellRoot = "/var/lib/fops/shell/"
)
