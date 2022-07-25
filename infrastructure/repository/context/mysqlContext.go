package context

import (
	"fs/data"
)

type mysqlContext struct {
	*data.DbContext
	Admin         data.TableSet[adminAgent.PO]
	Build         data.TableSet[buildAgent.PO]
	Cluster       data.TableSet[clusterAgent.PO]
	DockerfileTpl data.TableSet[dockerfileTplAgent.PO]
	DockerHub     data.TableSet[dockerHubAgent.PO]
	Git           data.TableSet[gitAgent.PO]
	Project       data.TableSet[projectAgent.PO]
	ProjectGroup  data.TableSet[projectGroupAgent.PO]
	YamlTpl       data.TableSet[yamlTplAgent.PO]
}

func NewContext() mysqlContext {
	context := data.NewDbContext("mysql")
	return mysqlContext{
		DbContext:     context,
		Admin:         data.NewTableSet(context, "admin", adminAgent.PO{}),
		Build:         data.NewTableSet(context, "build", buildAgent.PO{}),
		Cluster:       data.NewTableSet(context, "k8s_cluster", clusterAgent.PO{}),
		DockerfileTpl: data.NewTableSet(context, "docker_file_tpl", dockerfileTplAgent.PO{}),
		DockerHub:     data.NewTableSet(context, "docker_hub", dockerHubAgent.PO{}),
		Git:           data.NewTableSet(context, "basic_git", gitAgent.PO{}),
		Project:       data.NewTableSet(context, "basic_project", projectAgent.PO{}),
		ProjectGroup:  data.NewTableSet(context, "basic_project_group", projectGroupAgent.PO{}),
		YamlTpl:       data.NewTableSet(context, "k8s_yaml_tpl", yamlTplAgent.PO{}),
	}
}
