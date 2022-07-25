package context

import (
	"fops/infrastructure/repository/agent/adminAgent"
	"fops/infrastructure/repository/agent/buildAgent"
	"fops/infrastructure/repository/agent/clusterAgent"
	"fops/infrastructure/repository/agent/dockerHubAgent"
	"fops/infrastructure/repository/agent/dockerfileTplAgent"
	"fops/infrastructure/repository/agent/gitAgent"
	"fops/infrastructure/repository/agent/projectAgent"
	"fops/infrastructure/repository/agent/projectGroupAgent"
	"fops/infrastructure/repository/agent/yamlTplAgent"
	"fs/data"
)

type MysqlContext struct {
	*data.DbContext `data:"name=mysql"`
	//Admin           data.TableSet[adminAgent.PO]         `data:"name=admin"`
	Admin         data.TableSet[adminAgent.PO]         `data:"name=admin"`
	Build         data.TableSet[buildAgent.PO]         `data:"name=build"`
	Cluster       data.TableSet[clusterAgent.PO]       `data:"name=cluster"`
	DockerfileTpl data.TableSet[dockerfileTplAgent.PO] `data:"name=dockerfile_tpl"`
	DockerHub     data.TableSet[dockerHubAgent.PO]     `data:"name=docker_hub"`
	Git           data.TableSet[gitAgent.PO]           `data:"name=basic_git"`
	Project       data.TableSet[projectAgent.PO]       `data:"name=basic_project"`
	ProjectGroup  data.TableSet[projectGroupAgent.PO]  `data:"name=basic_project_group"`
	YamlTpl       data.TableSet[yamlTplAgent.PO]       `data:"name=k8s_yaml_tpl"`
}

var Mysql = data.Init[MysqlContext]()

//func NewContext() MysqlContext {
//	context := data.NewDbContext("mysql")
//	return MysqlContext{
//		//DbContext:     context,
//		Admin:         data.NewTableSet(context, "admin", adminAgent.PO{}),
//		Build:         data.NewTableSet(context, "build", buildAgent.PO{}),
//		Cluster:       data.NewTableSet(context, "k8s_cluster", clusterAgent.PO{}),
//		DockerfileTpl: data.NewTableSet(context, "docker_file_tpl", dockerfileTplAgent.PO{}),
//		DockerHub:     data.NewTableSet(context, "docker_hub", dockerHubAgent.PO{}),
//		Git:           data.NewTableSet(context, "basic_git", gitAgent.PO{}),
//		Project:       data.NewTableSet(context, "basic_project", projectAgent.PO{}),
//		ProjectGroup:  data.NewTableSet(context, "basic_project_group", projectGroupAgent.PO{}),
//		YamlTpl:       data.NewTableSet(context, "k8s_yaml_tpl", yamlTplAgent.PO{}),
//	}
//}
