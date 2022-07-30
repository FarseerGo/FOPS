package device

import (
	"context"
	"fops/domain/_/eumK8SControllers"
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/exec"
	"github.com/farseernet/farseer.go/utils/file"
)

func init() {
	_ = container.Register(func() devicer.IKubectlDevice { return &kubectlDevice{} })
}

type kubectlDevice struct {
}

func (kubectlDevice) GetConfigFile(clusterName string) string {
	return vo.KubeRoot + clusterName
}

func (device kubectlDevice) CreateConfigFile(clusterName string, clusterConfig string) string {
	configFile := device.GetConfigFile(clusterName)
	// 文件不存在，则创建
	if !file.IsExists(configFile) {
		file.WriteString(configFile, clusterConfig)
	} else {
		// 比对配置是否不一样，不一样则覆盖新的
		var config = file.ReadString(configFile)
		if clusterConfig != config {
			file.WriteString(configFile, clusterConfig)
		}
	}
	return configFile
}

func (device kubectlDevice) SetYaml(clusterName string, projectName string, yamlContent string, progress chan string, ctx context.Context) bool {
	// 将yaml文件写入临时文件
	fileName := "/tmp/" + projectName + ".yaml"
	file.Delete(fileName)
	file.WriteString(fileName, yamlContent)

	configFile := device.GetConfigFile(clusterName)

	// 发布
	var exitCode = exec.RunShellContext("kubectl apply -f "+fileName+" --kubeconfig="+configFile+" --insecure-skip-tls-verify", progress, nil, "", ctx)
	if exitCode != 0 {
		progress <- "K8S更新镜像失败。"
		return false
	}
	progress <- "更新镜像版本完成。"
	return true
}

func (device kubectlDevice) SetImages(cluster vo.ClusterVO, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string, ctx context.Context) bool {
	return device.SetImagesByClusterName(cluster.Name, cluster.Config, projectName, dockerImages, k8SControllersType, progress, ctx)

}

func (device kubectlDevice) SetImagesByClusterName(clusterName string, clusterConfig string, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string, ctx context.Context) bool {

	progress <- "---------------------------------------------------------"
	progress <- "开始更新K8S POD的镜像版本。"

	var configFile = device.CreateConfigFile(clusterName, clusterConfig)
	var exitCode = exec.RunShellContext("kubectl set image "+eumK8SControllers.GetName(k8SControllersType)+"/"+projectName+" "+projectName+"="+dockerImages+" --kubeconfig="+configFile+" --insecure-skip-tls-verify", progress, nil, "", ctx)
	if exitCode != 0 {
		progress <- "K8S更新镜像失败。"
		return false
	}
	progress <- "更新镜像版本完成。"
	return true
}
