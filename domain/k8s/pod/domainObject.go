package pod

import (
	"fops/domain/_/eumK8SKind"
	"fs/utils/lambda"
	"strconv"
	"strings"
)

type DomainObject struct {
	// 主键
	Id int
	// 项目组ID
	GroupId int
	// 项目名称
	Name string
	// 程序入口名称
	EntryPoint string
	// 程序启动端口
	EntryPort int
	// 访问域名
	Domain string
	// 项目路径
	Path string
	// K8SDeployment模板
	K8STplDeployment YamlTplVO
	// K8SDeployment模板列表
	K8STplDeploymentList []YamlTplVO
	// K8SIngress模板
	K8STplIngress YamlTplVO
	// K8STplIngress模板列表
	K8STplIngressList []YamlTplVO
	// K8SService模板
	K8STplService YamlTplVO
	// K8STplIngress模板列表
	K8STplServiceList []YamlTplVO
	// K8SConfig模板
	K8STplConfig YamlTplVO
	// K8STplConfig模板列表
	K8STplConfigList []YamlTplVO
	// K8S模板自定义变量(K1=V1,K2=V2)
	K8STplVariable string
}

func New() DomainObject {
	return DomainObject{}
}

// SetYamlTpl 设置Yaml模板属性
func (do *DomainObject) SetYamlTpl(lstYaml []YamlTplVO, deploymentId int, ingressId int, serviceId int, configId int) {
	do.K8STplDeployment = lambda.By(lstYaml).Find(func(o YamlTplVO) bool { return o.Id == deploymentId })
	do.K8STplIngress = lambda.By(lstYaml).Find(func(o YamlTplVO) bool { return o.Id == ingressId })
	do.K8STplService = lambda.By(lstYaml).Find(func(o YamlTplVO) bool { return o.Id == serviceId })
	do.K8STplConfig = lambda.By(lstYaml).Find(func(o YamlTplVO) bool { return o.Id == configId })

	do.K8STplDeploymentList = lambda.By(lstYaml).FindAll(func(o YamlTplVO) bool { return o.K8SKindType == eumK8SKind.Controllers })
	do.K8STplIngressList = lambda.By(lstYaml).FindAll(func(o YamlTplVO) bool { return o.K8SKindType == eumK8SKind.Ingress })
	do.K8STplServiceList = lambda.By(lstYaml).FindAll(func(o YamlTplVO) bool { return o.K8SKindType == eumK8SKind.Service })
	do.K8STplConfigList = lambda.By(lstYaml).FindAll(func(o YamlTplVO) bool { return o.K8SKindType == eumK8SKind.Config })
}

// 将已使用的模板，合并成一个大的yaml
func (do *DomainObject) MergeTplYaml() string {
	lstYaml := []string{do.K8STplDeployment.Template, do.K8STplService.Template, do.K8STplIngress.Template, do.K8STplConfig.Template}
	//lstYaml=lambda.By(lstYaml).RemoveNil()
	lambda.By(lstYaml).RemoveNil()

	// 替换模板
	for index, _ := range lstYaml {
		lstYaml[index] = do.ReplaceTpl(lstYaml[index])
	}

	return strings.Join(lstYaml, "\r\n---\r\n")
}

// ReplaceTpl 替换模板
func (do *DomainObject) ReplaceTpl(dockerfileTpl string) string {
	// 替换项目名称
	yaml := strings.ReplaceAll(dockerfileTpl, "${project_name}", do.Name)
	yaml = strings.ReplaceAll(yaml, "${domain}", do.Domain)
	yaml = strings.ReplaceAll(yaml, "${entry_point}", do.EntryPoint)
	yaml = strings.ReplaceAll(yaml, "${entry_port}", strconv.Itoa(do.EntryPort))
	yaml = strings.ReplaceAll(yaml, "${project_path}", strings.TrimPrefix(do.Path, "/"))

	// 替换模板变量
	for _, kv := range strings.Split(do.K8STplVariable, ",") {
		kvGroup := strings.Split(kv, "=")
		if len(kvGroup) != 2 {
			continue
		}
		yaml = strings.ReplaceAll(yaml, "${{{kvGroup[0]}}}", kvGroup[1])
	}
	return yaml
}
