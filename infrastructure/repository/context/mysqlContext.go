package context

import (
	"fops/infrastructure/repository/model"
	"github.com/farseernet/farseer.go/data"
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
