package device

import (
	"context"
	"fops/domain/building/build/event"
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/exec"
	"github.com/farseernet/farseer.go/utils/file"
	"github.com/farseernet/farseer.go/utils/str"
	"path/filepath"
)

func init() {
	_ = container.Register(func() devicer.IGitDevice { return &gitDevice{} })
}

type gitDevice struct {
}

func (device gitDevice) GetGitPath(gitHub string) string {
	if gitHub == "" {
		return ""
	}
	var gitName = device.GetName(gitHub)
	return vo.GitRoot + gitName + "/"
}

func (gitDevice) GetName(gitHub string) string {
	if gitHub == "" {
		return ""
	}
	git := filepath.Base(gitHub)
	return str.CutRight(git, ".git")
}

func (gitDevice) RememberPassword(env vo.EnvVO, progress chan string) {
	exec.RunShell("git config --global credential.helper store", progress, env.ToMap(), "")
}

func (gitDevice) ExistsGitProject(gitPath string) bool {
	// 如果Git存放的目录不存在，则创建
	if !file.IsExists(vo.GitRoot) {
		file.CreateDir766(vo.GitRoot)
	}
	return file.IsExists(gitPath)
}

func (device gitDevice) Clear(gitHub string, progress chan string) bool {
	// 获取Git存放的路径
	gitPath := device.GetGitPath(gitHub)
	exitCode := exec.RunShell("rm -rf "+gitPath, progress, nil, "")
	if exitCode != 0 {
		progress <- "Git清除失败"
		return false
	}
	return true
}

func (device gitDevice) CloneOrPull(git vo.GitVO, progress chan string, ctx context.Context) bool {
	if progress == nil {
		progress = make(chan string, 100)
	}
	progress <- "---------------------------------------------------------"

	// 先得到项目Git存放的物理路径
	gitPath := device.GetGitPath(git.Hub)
	var execSuccess bool

	// 存在则使用pull
	if device.ExistsGitProject(gitPath) {
		progress <- "开始拉取git " + git.Name + " 分支：" + git.Branch + " 仓库：" + git.Hub + "。"
		execSuccess = pull(gitPath, progress, ctx)
	} else {
		progress <- "开始克隆git " + git.Name + " 分支：" + git.Branch + " 仓库：" + git.Hub + "。"
		execSuccess = device.clone(git.Hub, git.Branch, progress, ctx)
	}

	if execSuccess {
		// 通知更新git拉取时间
		event.GitCloneOrPulledEvent{GitId: git.Id}.PublishEvent()
	} else {
		progress <- "拉取出错了。"
	}
	return execSuccess
}

func (device gitDevice) CloneOrPullAndDependent(lstGit []vo.GitVO, progress chan string, ctx context.Context) bool {
	for _, git := range lstGit {
		if !device.CloneOrPull(git, progress, ctx) {
			return false
		}
	}
	progress <- "拉取完成。"
	return true
}

func pull(savePath string, progress chan string, ctx context.Context) bool {
	exitCode := exec.RunShellContext("git -C "+savePath+" pull --rebase", progress, nil, "", ctx)
	if exitCode != 0 {
		progress <- "Git拉取失败"
		return false
	}
	return true
}

func (device gitDevice) clone(github string, branch string, progress chan string, ctx context.Context) bool {
	gitPath := device.GetGitPath(github)
	exitCode := exec.RunShellContext("git clone -b "+branch+" "+github+" "+gitPath, progress, nil, "", ctx)
	if exitCode != 0 {
		progress <- "Git克隆失败"
		return false
	}
	return true
}
