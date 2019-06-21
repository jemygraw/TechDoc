最近在研究如何将 Jenkins 集成到发布环境中，重点研究了一下 Jenkins 中相关的 API。

对于 Jenkins 中的任务，Jenkins 提供了一些 RESTFUL 的API来获取这些任务的信息，或者是来触发新的任务构建等。

## 鉴权

Jenkins 的 API 使用的是 Basic 的鉴权方式，也就是在 Jenkins 中，我们可以使用登陆ID和相应的 API Token 来做鉴权。

假设我们的登陆ID是 `jinxinxin` ，相应的 API Token 为 `611e3be83c538f9bf8b25be0218f0832`。那么如果使用 curl 来模拟请求的话，请求的头部如下所示，就是把登陆ID和 API Token 用冒号 `:` 拼接起来，然后做 Base64 编码，前面再拼接上 `Basic `，构建了 Authorization 头部的值。

## API

### 启动一个新的 Jenkins Build

```
POST {JenkinsHost}/job/{JenkinsJobName}/build

Authorization: Basic <Token>

---

200 OK
```

例如：

```
$ curl --user jinxinxin:11e42a69ae0872b71c013ec3f825f9df43 http://localhost:8080/job/test-echo-java/build -X POST -v
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
* Server auth using Basic with user 'jinxinxin'
> POST /job/test-echo-java/build HTTP/1.1
> Host: localhost:8080
> Authorization: Basic amlueGlueGluOjExZTQyYTY5YWUwODcyYjcxYzAxM2VjM2Y4MjVmOWRmNDM=
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 201 Created
< Date: Wed, 10 Oct 2018 04:17:13 GMT
< X-Content-Type-Options: nosniff
< Location: http://localhost:8080/queue/item/12/
< Content-Length: 0
< Server: Jetty(9.4.z-SNAPSHOT)
<
* Connection #0 to host localhost left intact
```

该请求会触发一个新的构建，但是这个请求并不返回一个Body，而是通过状态码判断请求是否成功，成功为200，失败的可能是404表示job不存在或者其他错误等。
或许你会好奇，这个请求为什么不返回一个新的构建的 BuildNumber，因为这个请求本身触发的是将新的构建请求加入到队列中，这个时候还没有分配给这个构建任务一个 Build Number。

### 获取启动的 Jenkins Build 的 BuildNumber

我们为什么那么执着地一定要获取这个任务的 BuildNumber 呢？因为我们想要获取这个任务的具体执行情况，就必须要有这个 BuildNumber。我们使用下面的 API 来获取这个 BuildNumber。

```
GET {JenkinsHost}/job/{JenkinsJobName}/lastBuild/api/json

Authorization: Basic <Token>

---

200 OK

{
    ...
}
```

例如:

```
$ curl --user 'jinxinxin:611e3be83c538f9bf8b25be0218f0832' http://ci.example.com/job/test-echo-java/lastBuild/api/json
```

这个请求用来获取刚刚触发的构建任务的信息，我们从返回的JSON格式的Body里面，可以解析得到 Build 的 Number。

### 