Kubernetes 支持多种模式的API访问鉴权方式。包括私钥+证书模式，Basic 用户名密码模式，Bearer Token 模式等。其中最常用的是基于ServiceAccount的私钥+证书模式。不过另外两种模式也在支持范畴，所以我们也了解一下，方便特殊场景下的使用。

# Basic用户鉴权

首先，我们在API服务端的 `/etc/kubernetes/` 目录下新建一个 `users.csv` 文件，内容如下：

```
[root@ksnode1 kubernetes]# cat users.csv
pass123,jemy,1007,"developer"
```

然后在 APIServer 的启动命令行选项中加入选项

```
--basic-auth-file=/etc/kubernetes/users.csv
```

我们在集群外的一个客户端设置下访问集群的Basic账号信息，主要需要用户名，密码，集群ca证书以及集群的 API Server 服务地址。

```
$ kubectl config set-credentials jemy --username jemy --password pass123
User "jemy" set.
$ kubectl config set-cluster k8s-learning --server https://10.8.1.72:6444 --certificate-authority /etc/kubernetes/pki/env-jxx/ca.pem
$ kubectl config set-context k8s-learning-ctx --cluster k8s-learning --user jemy
$ kubectl config use-context k8s-learning-ctx
```

然后，我们尝试直接访问nodes命令，会发现报错如下：

```
$ kubectl get nodes
Error from server (Forbidden): nodes is forbidden: User "jemy" cannot list nodes at the cluster scope
```

这种错误表示权限验证通过了但是没有任何的权限。如果是权限验证不通过，报错如下：

```
$ kubectl get nodes
error: You must be logged in to the server (Unauthorized)
```

上面的这个问题主要是由K8S的访问权限管理模式决定的。K8S的资源访问权限的基本处理流程分为三步：

1. Authentication 权限验证，简单来说就是检查所使用的API的用户权限是否存在，比如用户名密码是否正确，就是一个登陆权限而已
2. Authorization 授权验证，简单来说就是检查这个用户权限是否拥有操作K8S资源的权限，对哪些资源有操作权限，只要存在一个资源的操作权限正确，就允许通过
3. Admission Control 准入控制，简单来说就是检查这个用户对这些资源的具体操作权限是否合法，存在一个不合法则全部拒绝操作

为了让这个用户能够访问集群资源，我们给这个普通的用户jemy授权一个系统内置的ClusterRole：`cluster-admin`

```
$ kubectl create clusterrolebinding cluster-admin-for-jemy --clusterrole cluster-admin --user jemy
clusterrolebinding.rbac.authorization.k8s.io/cluster-admin-for-jemy created
```

然后再运行上面的 `kubectl get nodes` 即可发现节点。

```
$ kubectl get nodes
NAME        STATUS    ROLES     AGE       VERSION
10.8.1.75   Ready     <none>    23d       v1.11.0
10.8.1.76   Ready     <none>    23d       v1.11.0
```

我们使用 `kubectl get clusterrole cluster-admin -o yaml` 看下这个内置的 cluster-admin 定义：

```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  creationTimestamp: 2018-08-27T01:31:58Z
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: cluster-admin
  resourceVersion: "43"
  selfLink: /apis/rbac.authorization.k8s.io/v1/clusterroles/cluster-admin
  uid: fd9b31a4-a998-11e8-811b-569fbff0044d
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - '*'
- nonResourceURLs:
  - '*'
  verbs:
  - '*'
```

这个 ClusterRole 里面定义了所有资源的访问，配置等操作权限。

为了了解权限控制，我们可以创建一个自己的ClusterRole，只允许读取节点和列举节点信息。

先把刚刚的ClusterRoleBinding删除

```
$ kubectl delete  clusterrolebinding cluster-admin-for-jemy 
```

然后先创建我们自己的 ClusterRole

```
$ kubectl create clusterrole my-cluster-admin --resource nodes --verb 'get,list'
clusterrole.rbac.authorization.k8s.io/my-cluster-admin created
```

重新创建绑定

```
$ kubectl create clusterrolebinding cluster-admin-for-jemy --clusterrole my-cluster-admin --user jemy
clusterrolebinding.rbac.authorization.k8s.io/cluster-admin-for-jemy created
```

读取某个节点信息，对应get操作

```
$ kubectl get nodes 10.8.1.75
NAME        STATUS    ROLES     AGE       VERSION
10.8.1.75   Ready     <none>    23d       v1.11.0
```

获取节点的列表，对应list操作

```
$ kubectl get nodes
NAME        STATUS    ROLES     AGE       VERSION
10.8.1.75   Ready     <none>    23d       v1.11.0
10.8.1.76   Ready     <none>    23d       v1.11.0
```

如果获取其他的资源比如 Pods，就会失败

```
$ kubectl get pods
Error from server (Forbidden): pods is forbidden: User "jemy" cannot list pods in the namespace "default"
```

在上面的例子中，基本介绍了普通用户在K8S中的使用方法，这里有个小的问题，如果我们改变这个用户名字，需要同步集群和客户端的哪些信息呢？

1. 改变 /etc/kubernetes/users.csv 中的用户名 jemy 为 jemygraw，并scp到各个API Server的节点
2. 重启APIServer服务器
3. 修改客户端~/.kube/config 里面定义users 的地方，把里面的 username 改为 jemygraw
4. 在集群中使用 kubectl edit clusterrolebinding cluster-admin-for-jemy，把里面定义subjects的地方kind:User的节点里面的name改为 jemygraw 即可
5. 然后客户端使用kubectl get nodes 就可以了

这个过程可以自行验证下，顺便再熟悉下上面的操作流程。

参考文档：https://kubernetes.io/docs/reference/access-authn-authz/authentication/