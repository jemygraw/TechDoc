最近在做部署工具的时候，将应用(Deployment)发布到集群后，发现Pod启动不起来，使用 kubectl describe pod xxx 发现错误如下：

```
----     ------     ----              ----                -------
Normal   Scheduled  1m                default-scheduler   Successfully assigned default/echo-java-5bb694886c-7wkkv to 10.8.1.76
Normal   Pulling    25s (x3 over 1m)  kubelet, 10.8.1.76  pulling image "reg.example.com/java-apps/echo-java:1.0.0"
Warning  Failed     25s (x3 over 1m)  kubelet, 10.8.1.76  Failed to pull image "reg.example.com/java-apps/echo-java:1.0.0": rpc error: code = Unknown desc = Error response from daemon: Get https://reg.example.com/v2/apps/echo-java/manifests/1.0.0: unauthorized: authentication required
Warning  Failed     25s (x3 over 1m)  kubelet, 10.8.1.76  Error: ErrImagePull
Normal   BackOff    11s (x4 over 1m)  kubelet, 10.8.1.76  Back-off pulling image "reg.example.com/java-apps/echo-java:1.0.0"
Warning  Failed     11s (x4 over 1m)  kubelet, 10.8.1.76  Error: ImagePullBackOff
```

从错误描述上看，说的是镜像拉取失败了。提示的错误是：`authentication required`，从这里我们判断出是因为集群里面没有配置拉取镜像用的 Secret 导致的。在Kubernetes中我们可以把镜像中心的登录信息写入到 Secret，然后通过 Deployment 的 YAML 文件来引用，从而实现让 kubelet 取拉取镜像的功能。

我们使用如下的命令创建一个保存镜像中心登录信息的 Secret：

```
$ kubectl create secret docker-registry myregistrykey --docker-server https://reg.example.com/v2/ --docker-username jemygraw --docker-password jemypasword --docker-email jemygraw@example.com
```

其中 `--docker-username` ， `--docker-password` ，`--docker-email` 分别填写登录的账号信息。而 `--docker-server` 填写注册中心的API服务地址。注意上面例子中的最后的部分为 `v2/` ，这个根据各自的注册中心不同可能不同，注意不要把最后的斜杠忘了。

在创建完成 Secret 之后，我们可以用命令查看下：

```
$ kubectl get secret myregistrykey
```

然后我们就可以在 Deployment 的 YAML 配置文件中使用这个名为 myregistrykey 的Secret。

```
kind: Deployment
...
spec:
  ...
  template:
    ...
    spec:
      ...
      imagePullSecrets:
      - name: myregistrykey
```

然后使用 kubectl delete pod xxx 删除那个Pod，让kubelet重建一个Pod就好了。