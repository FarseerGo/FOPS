package kubectlSetYaml

import (
	"context"
	"fops/application/k8s/cluster"
	"fops/domain/building/devicer"
	"fops/domain/k8s/pod"
	"fops/domain/k8s/yamlTpl"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/linq"

	"strings"
)

type app struct {
	repository    yamlTpl.Repository
	kubectlDevice devicer.IKubectlDevice
}

func NewApp() *app {
	return &app{
		repository:    container.Resolve[yamlTpl.Repository](),
		kubectlDevice: container.Resolve[devicer.IKubectlDevice](),
	}
}

// DeployYaml 发布
func (app *app) DeployYaml(cluster cluster.Dto, yaml string, progress chan string) bool {
	app.kubectlDevice.CreateConfigFile(cluster.Name, cluster.Config)
	ctx := context.Background()
	return app.kubectlDevice.SetYaml(cluster.Name, "single", yaml, progress, ctx)
}

// DeployPodBatch 发布
func (app *app) DeployPodBatch(lstProject []pod.DomainObject, cluster cluster.Dto, progress chan string) bool {
	if cluster.Id < 1 {
		panic("请先选择集群环境")
	}

	app.kubectlDevice.CreateConfigFile(cluster.Name, cluster.Config)
	ctx := context.Background()

	// 拼接已经选择的所有脚本
	var lstYaml []string
	linq.From(lstProject).Select(&lstYaml, func(item pod.DomainObject) any {
		return item.MergeTplYaml()
	})
	yaml := strings.Join(lstYaml, "\r\n---\r\n")

	return app.kubectlDevice.SetYaml(cluster.Name, "all", yaml, progress, ctx)
}

// DeployPod 发布
func (app *app) DeployPod(pod pod.DomainObject, cluster cluster.Dto, progress chan string) bool {
	if cluster.Id < 1 {
		panic("请先选择集群环境")
	}
	if pod.Id > 1 {
		panic("项目不存在")
	}

	app.kubectlDevice.CreateConfigFile(cluster.Name, cluster.Config)
	ctx := context.Background()

	// 拼接已经选择的所有脚本
	var yaml = pod.MergeTplYaml()
	return app.kubectlDevice.SetYaml(cluster.Name, "single", yaml, progress, ctx)
}
