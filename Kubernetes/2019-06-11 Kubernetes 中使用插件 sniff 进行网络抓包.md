# Kubernetes 中使用插件 sniff 进行网络抓包

## 背景

在 Kubernetes 的实际使用中，我们经常需要配合业务调查问题，对于微服务来说，这个问题更多的是查看 API 的调用情况，这些API或者采用 RPC 协议或者是采用 HTTP 的协议。这两种协议都是基于 TCP 的协议，所以一般我们会到容器中使用 tcpdump 工具来抓包，然后就地或者拿出来放到 wireshark 图形化软件里面分析。

这种情况下，需要我们的基础镜像提前把 tcpdump 等排查工具打包进去，否则线上安装 debug 软件，一者违反安全规则，另外如果需要支持的 Pod 过多，安装 debug 工具本身就有不小的工作量。


## krew

在 Kubernetes 中，有一个插件命令叫做 krew，可以通过这个命令来安装一个叫做 sniff 的插件工具来完成这个工作。下面我们先看看如何安装这个 krew 插件。

krew 的项目地址在：[https://github.com/kubernetes-sigs/krew](https://github.com/kubernetes-sigs/krew) 。如果有兴趣可以自行浏览，我们这里介绍下在 Centos 等 Linux 下面如何安装。

首先，需要确认系统安装了 git 。

其次，复制下面的命令到终端软件中，这段命令会去下载和安装这个 krew 插件。

```
$(
  set -x; cd "$(mktemp -d)" &&
  curl -fsSLO "https://storage.googleapis.com/krew/v0.2.1/krew.{tar.gz,yaml}" &&
  tar zxvf krew.tar.gz &&
  ./krew-"$(uname | tr '[:upper:]' '[:lower:]')_amd64" install \
    --manifest=krew.yaml --archive=krew.tar.gz
)
```

安装好的 krew 命令在目录 `~/.krew/bin` 下面，所以我们可以把这个路径加到终端的配置文件中。一般是 `~/.bashrc` 或者是 `~/.zshrc`。

例如，使用如下的命令将 krew 命令的所在路径写入到 $PATH 中。

```
$ echo 'export PATH=$PATH:$HOME/.krew/bin' >> ~/.bashrc
```

需要注意的是，上面设置的这个路径同时也是通过 krew 安装的其他的插件命令所在的目录，所以一次设置后面其他的命令都可以直接使用了。

安装完 krew 之后，我们需要通过 `kubectl krew update` 命令来更新支持的插件命令列表，下载下来的 Plugin 配置文件都存放在 `~/.krew/index/plugins`  下面。

```
$ kubectl krew update
Updated the local copy of plugin index.
```

## sniff

我们可以使用命令 `kubectl krew install sniff` 来安装这个插件命令，安装好的命令位于 `~/.kube/store/sniff` 下面。

```
$ kubectl krew install sniff
Updated the local copy of plugin index.
Installing plugin: sniff
CAVEATS:
\
 |  This plugin needs the following programs:
 |  * wireshark (optional, used for live capture)
/
Installed plugin: sniff
```

我们可以看下 sniff 命令的所在目录下的文件。

```
$ ls -ahl ~/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/
total 36M
drwx------ 2 root root 4.0K Jun 11 17:58 .
drwxr-xr-x 3 root root 4.0K Jun 11 17:58 ..
-rwxr-xr-x 1 root root  33M Jun 11 17:58 kubectl-sniff
-rwxr-xr-x 1 root root 2.6M Jun 11 17:58 static-tcpdump
```

这里面其实下载了两个文件，其中一个是 sniff，另外一个是 tcpdump，稍后我们会看到这个 tcpdump 的用途。

## Pod 抓包

我们首先找个 Pod 来研究下 sniff 的具体抓包操作方法。

```$ kubectl get pods -n devops
NAME                         READY   STATUS    RESTARTS   AGE
echo-go-bdf4bd7ff-v6hml      1/1     Running   0          8h
echo-java-55c5dcbbc9-7dh5c   1/1     Running   1          30h
```

我们拿第一个 Pod 出来测试下抓包。

```
$ kubectl sniff echo-go-bdf4bd7ff-v6hml -n devops
<1> INFO[0000] sniffing method: upload static tcpdump
<2> INFO[0000] using tcpdump path at: '/root/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/static-tcpdump'
<3> INFO[0000] no container specified, taking first container we found in pod.
<4> INFO[0000] selected container: 'echo-go'
<5> INFO[0000] sniffing on pod: 'echo-go-bdf4bd7ff-v6hml' [namespace: 'devops', container: 'echo-go', filter: '', interface: 'any']
<6> INFO[0000] uploading static tcpdump binary from: '/root/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/static-tcpdump' to: '/tmp/static-tcpdump'
<7> INFO[0000] uploading file: '/root/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/static-tcpdump' to '/tmp/static-tcpdump' on container: 'echo-go'
<8> INFO[0000] executing command: '[/bin/sh -c ls -alt /tmp/static-tcpdump]' on container: 'echo-go', pod: 'echo-go-bdf4bd7ff-v6hml', namespace: 'devops'
<9> INFO[0000] command: '[/bin/sh -c ls -alt /tmp/static-tcpdump]' executing successfully exitCode: '0', stdErr :''
<10> INFO[0000] file found: '-rwxr-xr-x 1 root root 2642872 Jan  1  1970 /tmp/static-tcpdump'
<11> INFO[0000] file was already found on remote pod
<12> INFO[0000] tcpdump uploaded successfully
<13> INFO[0000] spawning wireshark!
<14> INFO[0000] starting sniffer cleanup
<15> INFO[0000] sniffer cleanup completed successfully
<16> Error: exec: "wireshark": executable file not found in $PATH
```

为了方便叙述，我们把每一行输出都标上了数字。这些数字不在实际输出的结果中。 从第 <1> - <7> 行我们可以看出，sniff 命令把我们刚刚看到的 static-tcpdump 上传到 Pod 中，我们可以到 Pod 里面在 `/tmp` 目录下发现这个 `static-tcpdump` 文件。然后试图启动 wireshark 进程，但是由于我们服务器一般不安装 wireshark 所以启动失败了。

这种情况下，我们可以使用一个新的选项 `-o` 来将抓包的内容输出到文件中。

```
$ kubectl sniff echo-go-bdf4bd7ff-v6hml -n devops -o test.pcap
kubectl sniff echo-go-bdf4bd7ff-v6hml -n devops -o pcap
INFO[0000] sniffing method: upload static tcpdump
INFO[0000] using tcpdump path at: '/root/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/static-tcpdump'
INFO[0000] no container specified, taking first container we found in pod.
INFO[0000] selected container: 'echo-go'
INFO[0000] sniffing on pod: 'echo-go-bdf4bd7ff-v6hml' [namespace: 'devops', container: 'echo-go', filter: '', interface: 'any']
INFO[0000] uploading static tcpdump binary from: '/root/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/static-tcpdump' to: '/tmp/static-tcpdump'
INFO[0000] uploading file: '/root/.krew/store/sniff/71102253eded8900c8f7b0d0624c65b3c77ecd6bcd28fabc9a200daac502282a/static-tcpdump' to '/tmp/static-tcpdump' on container: 'echo-go'
INFO[0000] executing command: '[/bin/sh -c ls -alt /tmp/static-tcpdump]' on container: 'echo-go', pod: 'echo-go-bdf4bd7ff-v6hml', namespace: 'devops'
INFO[0000] command: '[/bin/sh -c ls -alt /tmp/static-tcpdump]' executing successfully exitCode: '0', stdErr :''
INFO[0000] file found: '-rwxr-xr-x 1 root root 2642872 Jan  1  1970 /tmp/static-tcpdump
'
INFO[0000] file was already found on remote pod
INFO[0000] tcpdump uploaded successfully
INFO[0000] output file option specified, storing output in: 'pcap'
INFO[0000] start sniffing on remote container
INFO[0000] executing command: '[/tmp/static-tcpdump -i any -U -w - ]' on container: 'echo-go', pod: 'echo-go-bdf4bd7ff-v6hml', namespace: 'devops'
```

这种情况下，我们看到容器中的 `/tmp/static-tcpdump` 命令已经启动了，并把输出导向到 stdout，然后我们在 sniff 命令中把它写入到 test.pcap 文件中。

最后，我们就可以用本地的 wireshark 图形化工具打开这个 test.pcap 文件进行分析了。