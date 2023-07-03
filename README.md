<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">KubeFish</h1>
<h4 align="center">K8s Dashboard</h4>
<hr />

KubeFish 是一个现代化的 K8s 面板

## 功能展示

### 创建资源
<img src="./img/resource/create.png"/>

### 集群

#### 节点
<img src="./img/cluster/node.gif"/>

#### 命名空间
<img src="./img/cluster/namespace.gif"/>

### 工作负载

#### Pod
<img src="./img/workloads/pod.gif"/>

#### ReplicaSet
<img src="./img/workloads/replicaSet.gif"/>

#### Deployment
<img src="./img/workloads/deployment.gif"/>

#### DaemonSet
<img src="./img/workloads/daemonSet.gif"/>

### 配置

#### ConfigMap
<img src="./img/config/configMap.gif"/>

#### Secrets
<img src="./img/config/secrets.gif"/>

### 访问控制

#### ServiceAccount
<img src="./img/accessControl/serviceAccount.gif"/>

#### Role && RoleBinding
<img src="./img/accessControl/role.gif"/>

#### ClusterRole && ClusterRoleBinding
<img src="./img/accessControl/clusterRole.gif"/>

## 安装

### 集群外安装

Docker Compose 版本需要 V2

```bash
git clone https://github.com/pddzl/kubefish
cd kubefish
# 添加k8s config文件 (/etc/kubernetes/admin.conf 重命名为config)
# cp /tmp/config docker-compose/k8s
docker-compose -f docker-compose/docker-compose.yml build
docker-compose -f docker-compose/docker-compose.yml up -d
```

浏览器打开 `http://node节点:35999`

## 致谢
 + 项目脚手架 [td27-admin](https://github.com/pddzl/td27-admin)

## 📄 License

[MIT](./LICENSE)

Copyright (c) 2022-present [pddzl](https://github.com/pddzl)
