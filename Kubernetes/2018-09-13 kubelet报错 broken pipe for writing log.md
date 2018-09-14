最近发现从kubelet的日志中发现一些 Error 级别的错误。主要的错误信息为 `write: broken pipe when writing log for log file` 和 `write: connection reset by peer when writing log for log file` ，详细信息如下：

```
2018-09-13 13:36 Test Cluster 
1-132.example.net kubelet -- E0913 13:36:00.011285 5203 kuberuntime_logs.go:201] Failed with err write tcp 10.8.1.132:10250->10.8.1.55:62786: write: broken pipe when writing log for log file /var/log/pods/d63e3f59-b715-11e8-811e-427775cbe8d8/shop-pay-server_0.log: &{timestamp:{sec:63672413760 nsec:10835564 loc:<nil>} stream:stdout log:[9 97 116 32 111 114 103 46 115 112 114 105 110 103 102 114 97 109 101 119 111 114 107 46 97 111 112 46 102 114 97 109 101 119 111 114 107 46 82 101 102 108 101 99 116 105 118 101 77 101 116 104 111 100 73 110 118 111 99 97 116 105 111 110 46 112 114 111 99 101 101 100 40 82 101 102 108 101 99 116 105 118 101 77 101 116 104 111 100 73 110 118 111 99 97 116 105 111 110 46 106 97 118 97 58 49 55 57 41 10]}
```

```
2018-09-13 13:59 Test Cluster 
1-132.example.net kubelet -- E0913 13:59:00.003444 5203 kuberuntime_logs.go:201] Failed with err write tcp 10.8.1.132:10250->10.8.1.55:63848: write: connection reset by peer when writing log for log file /var/log/pods/d63e3f59-b715-11e8-811e-427775cbe8d8/shop-pay-server_0.log: &{timestamp:{sec:63672415140 nsec:3320231 loc:<nil>} stream:stdout log:[50 48 49 56 45 48 57 45 49 51 32 49 51 58 53 57 58 48 48 46 48 48 51 32 91 115 104 97 110 100 105 97 110 45 112 97 121 45 115 101 114 118 101 114 45 55 98 52 102 53 56 100 56 56 53 45 99 119 110 99 52 32 124 32 115 104 97 110 100 105 97 110 45 112 97 121 45 115 101 114 118 101 114 32 124 32 45 32 124 32 112 111 111 108 45 55 45 116 104 114 101 97 100 45 49 48 93 32 68 69 66 85 71 32 111 46 115 46 106 100 98 99 46 100 97 116 97 115 111 117 114 99 101 46 68 97 116 97 83 111 117 114 99 101 85 116 105 108 115 32 45 32 70 101 116 99 104 105 110 103 32 74 68 66 67 32 67 111 110 110 101 99 116 105 111 110 32 102 114 111 109 32 68 97 116 97 83 111 117 114 99 101 10]}
```

其中，`10.8.1.132` 是一个 Node 节点，然后 `10.8.1.55` 是 Master 节点，所以这个错误的信息就是 Node 节点上面的 kubelet 在和Master进行通信，尝试往 Master 上面写入数据的时候报错了。而且这个错误看上去是 Master 那边主动断掉了链接（从 connection reset by peer判断）。所以，我们在错误的日志里面找到了一个关键字，就是 Pod 的名字 shop-pay-server ，然后我们在 master上面查  `journalctl |grep 'Sep 13 13:59'  |grep 'shop-pay-server'` 的日志得到如下内容：

```
Sep 13 13:59:00 1-55-k8s-master.example.net kube-apiserver[90382]: I0913 13:59:00.002017   90382 trace.go:76] Trace[1375012443]: "Get /api/v1/namespaces/default/pods/shop-pay-server-7b4f58d885-cwnc4/log" (started: 2018-09-13 13:48:08.166214535 +0800 CST) (total time: 10m51.835742077s):
Sep 13 13:59:00 1-55-k8s-master.example.net kube-apiserver[90382]: Trace[1375012443]: [10m51.835742077s] [10m51.832181896s] END
Sep 13 13:59:00 1-55-k8s-master.example.net kube-apiserver[90382]: I0913 13:59:00.002103   90382 wrap.go:42] GET /api/v1/namespaces/default/pods/shop-pay-server-7b4f58d885-cwnc4/log?follow=true: (10m51.836399608s) 200 [[kubectl/v1.8.9 (linux/amd64) kubernetes/3fb1aaf] 10.8.1.59:31962]
```

从调用的API上面，有一个关键信息 `kubectl/v1.8.9 (linux/amd64)`，据此分析是有人使用kubectl进行日志查询。从源码层面分析的话，也会了解到请求

```
GET /api/v1/namespaces/default/pods/shop-pay-server-7b4f58d885-cwnc4/log?follow=true
```

其实是一个 kubelet 命令触发的。然后我们从 master 的 history 调查执行过的命令发现触发这个操作的命令：

```
1027  2018-09-13 13:48:07 root kubectl logs -f shop-pay-server-7b4f58d885-cwnc4
```

该问题复现方法：命令行输入上面的logs命令，然后等一会儿有日志输出的时候，ctrl+c掉，即会输出上面的报警信息。