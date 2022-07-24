package context

import (
	"fops/infrastructure/repository/agent/admin"
	"fops/infrastructure/repository/agent/build"
	"fops/infrastructure/repository/agent/cluster"
	"fops/infrastructure/repository/agent/dockerHub"
	"fops/infrastructure/repository/agent/dockerfileTpl"
	"fops/infrastructure/repository/agent/git"
	"fops/infrastructure/repository/agent/project"
	"fops/infrastructure/repository/agent/projectGroup"
	"fops/infrastructure/repository/agent/yamlTpl"
	"fs/data"
)

type mysqlContext struct {
	*data.DbContext
	Admin         data.TableSet[admin.PO]
	Build         data.TableSet[build.PO]
	Cluster       data.TableSet[cluster.PO]
	DockerfileTpl data.TableSet[dockerfileTpl.PO]
	DockerHub     data.TableSet[dockerHub.PO]
	Git           data.TableSet[git.PO]
	Project       data.TableSet[project.PO]
	ProjectGroup  data.TableSet[projectGroup.PO]
	YamlTpl       data.TableSet[yamlTpl.PO]
}

func NewContext() mysqlContext {
	context := data.NewDbContext("mysql")
	return mysqlContext{
		DbContext:     context,
		Admin:         data.NewTableSet(context, "admin", admin.PO{}),
		Build:         data.NewTableSet(context, "build", build.PO{}),
		Cluster:       data.NewTableSet(context, "k8s_cluster", cluster.PO{}),
		DockerfileTpl: data.NewTableSet(context, "docker_file_tpl", dockerfileTpl.PO{}),
		DockerHub:     data.NewTableSet(context, "docker_hub", dockerHub.PO{}),
		Git:           data.NewTableSet(context, "basic_git", git.PO{}),
		Project:       data.NewTableSet(context, "basic_project", project.PO{}),
		ProjectGroup:  data.NewTableSet(context, "basic_project_group", projectGroup.PO{}),
		YamlTpl:       data.NewTableSet(context, "k8s_yaml_tpl", yamlTpl.PO{}),
	}
}
