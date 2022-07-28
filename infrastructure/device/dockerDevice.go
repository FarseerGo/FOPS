package device

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	
)

func init() {
	_ = container.Register(func() devicer.IDockerDevice { return &dockerDevice{} })
}

type dockerDevice struct {
}

func (dockerDevice) GetDockerHub(dockerHubAddress string) string {
	var dockerHub = "localhost";
	if dockerHubAddress != ""	{
		dockerHub = dockerHubAddress;
		if (dockerHub.EndsWith("/")) dockerHub.Substring(0, dockerHub.Length - 1);
	}

	return dockerHub;
}

func (dockerDevice) GetDockerImage(dockerHubAddress string, projectName string, buildNumber int) string {
	return $"{GetDockerHub(dockerHubAddress)}:{projectName}-{buildNumber}";
}

func (dockerDevice) Login(dockerHub string, loginName string, loginPwd string, progress chan string, env vo.EnvVO, ctx context.Context) bool {
	receiveOutput.Report("---------------------------------------------------------");
	if (!string.IsNullOrWhiteSpace(dockerHub) && !string.IsNullOrWhiteSpace(loginName))
	{
		var result = await ShellTools.Run("docker", $"login {dockerHub} -u {loginName} -p {loginPwd}", receiveOutput, env, null, cancellationToken);
		if (result != 0)
		{
			receiveOutput.Report("镜像仓库登陆失败。");
			return false;
		}
	}

	receiveOutput.Report("镜像仓库登陆成功。");
	return true;
}

func (dockerDevice) ExistsDockerfile(dockerfilePath string) bool {
	return File.Exists(dockerfilePath);
}

func (dockerDevice) CreateDockerfile(projectName string, dockerfileContent string, ctx context.Context) {
	if (File.Exists(EnvVO.DockerfilePath)) File.Delete(EnvVO.DockerfilePath);
	await File.AppendAllTextAsync(EnvVO.DockerfilePath, dockerfileContent, cancellationToken);
}

func (dockerDevice) Build(env vo.EnvVO, progress chan string, ctx context.Context) bool {
	receiveOutput.Report("---------------------------------------------------------");
	receiveOutput.Report($"开始镜像打包。");

	// 打包
	var result = await ShellTools.Run("docker", $"build -t {env.DockerImage} --network=host -f {EnvVO.DockerfilePath} {EnvVO.DistRoot}", receiveOutput, env, EnvVO.DistRoot, cancellationToken);

	receiveOutput.Report(result == 0 ? $"镜像打包完成。" : $"镜像打包出错了。");
	return result == 0;
}

func (dockerDevice) Push(env vo.EnvVO, progress chan string, ctx context.Context) bool {
	// 上传
	var result = await ShellTools.Run("docker", $"push {env.DockerImage}", receiveOutput, env, null, cancellationToken);

	if (result == 0)
	{
		receiveOutput.Report($"镜像上传完成。");

		// 上传成功后，需要更新项目中的镜像版本属性
		new DockerPushedEvent(env.ProjectId, env.BuildNumber).PublishEvent();
		return true;
	}

	receiveOutput.Report( $"镜像上传出错了。");
	return false;
}
