## test
测试环境：

主机：亚太二区-A云主机

操作系统：Ubuntu Server 18.04.5 LTS 64bit

ks集群：All-In-One

镜像仓库：DockerHub公有云仓库

测试条件：安装Kubesphere集群，开启logging组件，并将fluent-bit升级为1.7.3

测试范围：对fluent-bit1.7.3日志收集及其兼容性进行测试

测试内容：测试fluent-bit 具体性能

测试步骤：1.编写测试代码

2.编写dockerfile，打包镜像

3.push到Dockerhub，再在KubeSphere集群中pull下来

4.运行测试镜像，修改fluent-bit的yaml文件，让fb接收信息，测试fb的功能


预期结果：获取fluent-bit接收日志的最大值，了解fluent-bit的具体性能