package context

import (
	"fops/infrastructure/repository/model"
	"fs/data"
)

type MysqlContext struct {
	Admin         data.TableSet[model.AdminPO]         `data:"name=admin"`
	Build         data.TableSet[model.BuildPO]         `data:"name=build"`
	Cluster       data.TableSet[model.ClusterPO]       `data:"name=cluster"`
	DockerfileTpl data.TableSet[model.DockerfileTplPO] `data:"name=dockerfile_tpl"`
	DockerHub     data.TableSet[model.DockerHubPO]     `data:"name=docker_hub"`
	Git           data.TableSet[model.GitPO]           `data:"name=basic_git"`
	Project       data.TableSet[model.ProjectPO]       `data:"name=basic_project"`
	ProjectGroup  data.TableSet[model.ProjectGroupPO]  `data:"name=basic_project_group"`
	YamlTpl       data.TableSet[model.YamlTplPO]       `data:"name=k8s_yaml_tpl"`
}

//var Mysql = data.Init[MysqlContext]()
//func NewContext() MysqlContext {
//	context := data.NewDbContext("mysql")
//	return MysqlContext{
//		//DbContext:     context,
//		Admin:         data.NewTableSet(context, "admin", adminAgent.AdminPO{}),
//		Build:         data.NewTableSet(context, "build", buildAgent.AdminPO{}),
//		Cluster:       data.NewTableSet(context, "k8s_cluster", clusterAgent.AdminPO{}),
//		DockerfileTpl: data.NewTableSet(context, "docker_file_tpl", dockerfileTplAgent.AdminPO{}),
//		DockerHub:     data.NewTableSet(context, "docker_hub", dockerHubAgent.AdminPO{}),
//		Git:           data.NewTableSet(context, "basic_git", gitAgent.AdminPO{}),
//		Project:       data.NewTableSet(context, "basic_project", projectAgent.AdminPO{}),
//		ProjectGroup:  data.NewTableSet(context, "basic_project_group", projectGroupAgent.AdminPO{}),
//		YamlTpl:       data.NewTableSet(context, "k8s_yaml_tpl", yamlTplAgent.AdminPO{}),
//	}
//}
