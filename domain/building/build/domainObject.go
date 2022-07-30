package build

import (
	"fmt"
	"fops/domain/_/eumBuildStatus"
	"fops/domain/building/build/vo"
	"github.com/farseernet/farseer.go/fsApp"
	"github.com/farseernet/farseer.go/utils/file"
	"strconv"
	"strings"
	"time"
)

type DomainObject struct {
	// 主键
	Id int
	// 构建号
	BuildNumber int
	// 状态
	Status eumBuildStatus.Enum
	// 是否成功
	IsSuccess bool
	// 开始时间
	CreateAt time.Time
	// 完成时间
	FinishAt time.Time
	// 构建的服务端id
	BuildServerId int64
	// 构建日志
	Log vo.LogVO
	// 环境变量
	Env vo.EnvVO
	// 项目
	Project vo.ProjectVO
	// Git
	Gits []vo.GitVO
	// Docker
	Docker vo.DockerVO
	// 集群信息
	Cluster vo.ClusterVO
}

// NewDO 用于Map
func NewDO() DomainObject {
	return DomainObject{}
}

// NewDO1 NewDO 添加新的构建
func NewDO1(buildNumber int, project vo.ProjectVO, gits []vo.GitVO, docker vo.DockerVO, cluster vo.ClusterVO) DomainObject {
	return DomainObject{
		BuildNumber:   buildNumber + 1,
		Status:        0,
		IsSuccess:     false,
		CreateAt:      time.Now(),
		FinishAt:      time.Now(),
		BuildServerId: fsApp.AppId,
		Project:       project,
		Gits:          gits,
		Docker:        docker,
		Cluster:       cluster,
	}
}

// NewDO 添加新的构建
func NewDO2(buildNumber int, project vo.ProjectVO, cluster vo.ClusterVO) DomainObject {
	return DomainObject{
		BuildNumber: buildNumber + 1,
		Project:     project,
		Cluster:     cluster,

		Status:        0,
		IsSuccess:     false,
		CreateAt:      time.Now(),
		FinishAt:      time.Now(),
		BuildServerId: fsApp.AppId,
	}
}

// GetGitMaster 获取Git主仓库
func (do *DomainObject) GetGitMaster() vo.GitVO {
	for _, git := range do.Gits {
		if git.IsMaster {
			return git
		}
	}
	panic("Git主仓库未设置")
}

// GenerateEnv 生成环境变量
func (do *DomainObject) GenerateEnv(projectGitRoot string, dockerHub string, dockerImage string, gitName string) {
	do.Env = vo.EnvVO{
		BuildId:           do.Id,
		BuildNumber:       do.BuildNumber,
		ProjectId:         do.Project.Id,
		ProjectName:       do.Project.Name,
		ProjectDomain:     do.Project.Domain,
		ProjectEntryPoint: do.Project.EntryPoint,
		ProjectEntryPort:  do.Project.EntryPort,
		DockerHub:         dockerHub,
		DockerImage:       dockerImage,
		ProjectGitRoot:    projectGitRoot,
		GitHub:            do.GetGitMaster().Hub,
		GitName:           gitName,
	}
}

// PrintEnv 打印环境变量
func (do *DomainObject) PrintEnv(progress chan string) {
	do.Env.Print(progress)
}

// GenerateDockerfileContent 替换模板
func (do *DomainObject) GenerateDockerfileContent() {
	// 替换项目名称
	var dockerfile = strings.ReplaceAll(do.Docker.DockerfileContent, "${project_name}", do.Project.Name)
	dockerfile = strings.ReplaceAll(dockerfile, "${domain}", do.Project.Domain)
	dockerfile = strings.ReplaceAll(dockerfile, "${entry_point}", do.Project.EntryPoint)
	dockerfile = strings.ReplaceAll(dockerfile, "${entry_port}", strconv.Itoa(do.Project.EntryPort))
	dockerfile = strings.ReplaceAll(dockerfile, "${git_name}", do.Env.GitName)
	dockerfile = strings.ReplaceAll(dockerfile, "${project_path}", strings.TrimPrefix(do.Project.Path, "/"))

	// 替换模板变量
	for _, kv := range strings.Split(do.Project.K8STplVariable, ",") {
		var kvGroup = strings.Split(kv, "=")
		if len(kvGroup) != 2 {
			continue
		}
		dockerfile = strings.ReplaceAll(dockerfile, "${{{kvGroup[0]}}}", kvGroup[1])
	}

	// 如果.net 应用，则自动实现csproj的递归复制并运行dotnet restore
	var csproj = file.GetFiles(vo.DistRoot, ".csproj", true)
	if len(csproj) > 0 {
		var lstCopyCmd []string

		for _, file := range csproj {
			filePath := file[len(vo.DistRoot):]
			fileDir := filePath[:strings.LastIndex(filePath, "/")+1]
			cmd := fmt.Sprintf("COPY [\"%s\",\"%s\"]", filePath, fileDir)
			lstCopyCmd = append(lstCopyCmd, cmd)
		}

		cmd := fmt.Sprintf("RUN dotnet restore %s/%s/ -s https://nuget.cdn.azure.cn/v3/index.json", do.Env.GitName, do.Project.Path)
		lstCopyCmd = append(lstCopyCmd, cmd)

		dockerfile = strings.ReplaceAll(dockerfile, "${dotnet_restore}", strings.Join(lstCopyCmd, "\r\n"))
		do.Docker = vo.NewDocker(do.Docker, dockerfile)
	}
}

func (do *DomainObject) Set(project vo.ProjectVO, docker vo.DockerVO, cluster vo.ClusterVO, gits []vo.GitVO) {
	do.Project = project
	do.Gits = gits
	do.Docker = docker
	do.Cluster = cluster
}
