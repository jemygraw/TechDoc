# 基于 Container 网络的容器抓包

## 背景

假设存在一个容器，提供的服务是 HTTP 或者 RPC 的服务。由于出于简单可维护的目的，这个容器的基础镜像里面没有带上任何和网络抓包相关的功能。那么如何能搞对这样的容器进行抓包，以分析业务上面可能存在的问题呢？

## 共享网络

Docker 的容器之间可以通过共享网络空间的方式，来让多个容器实现网络互通。这个意思直接一点就是如果你到每个容器内部去访问 localhost 监听的服务，无论这个服务在哪个容器里面，都能够访问成功。这就为网络抓包提供了基础。

## 专用镜像

我们可以使用一个专用的 tcpdump 的镜像来进行抓包。镜像名称为 corfr/tcpdump:latest ，可以使用 docker pull 直接下载。

```
$ docker pull corfr/tcpdump:latest
```

## 抓包实践

我们现在用一个提供简单 HTTP 服务的镜像来进行测试。

下载测试镜像

```
$ docker pull jemygraw/echo-go:1.0
```

启动测试容器

```
$ docker run -p 8080:8080 jemygraw/echo-go:1.0 /home/app/echo-go -port 8080
```

查看测试容器ID

```
$ docker ps
5de30e950459        jemygraw/echo-go:1.0   "/home/app/echo-go -…"   10 minutes ago      Up 10 minutes       0.0.0.0:8080->8080/tcp   dreamy_benz
```

启动抓包镜像，注意使用 `--network` 参数来共享测试容器的网络。

```
$ docker run --network container:5de30e950459  corfr/tcpdump:latest  -i any -U -w -
```

由于 `corfr/tcpdump:latest` 镜像构建的时候使用了 `ENTRYPOINT` 指定了入口命令为 `/usr/sbin/tcpdump` ，所以这里我们指定命令的选项参数即可。

这个时候我们可以尝试访问 `http://localhost:8080/` 就能够在抓包容器的输出中看到抓包结果了。

## 真实案例

Kubernetes 的插件命令 sniff 就是使用了上面的技术来实现特权模式下通过旁观的抓包容器共享目标容器的网络来实现抓包的。该项目地址在：[https://github.com/eldadru/ksniff](https://github.com/eldadru/ksniff) ，有兴趣可以研究。

