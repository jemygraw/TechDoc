在 K8S 的 APIServer 的代码中，依赖了一个叫做 go-restful 的库来构建 HTTP 的 API。
在学习 K8S 的代码过程中，我们要对这个库做些了解，这样才能更加方便地知道 APIServer 的 Restful 服务是如何构建的。
这个项目在 Github 上面的地址是：https://github.com/emicklei/go-restful 。我们也很容易地找到了作者的一篇介绍使用方法的博客。
帖子的地址在：http://ernestmicklei.com/2012/11/go-restful-first-working-example/ 。

这里我们将它简单翻译一下，帮助大家学习。下面的内容是译文。

---

在前面的一个帖子里面，我介绍了用 Google 的 Go 语言开发的，用来构建 REST 风格 WebService 的包 go-restful 的设计。
今天，我完成了这个包的实现，包含如下的功能：

1. 使用 Route 来创建 WebService，Route 是 HTTP Request 到 Go 函数的映射。
2. 每个 Route 都需要的信息包括 HTTP 请求方法（GET，POST，...)，URL 路径(/users)，MimeType以及其绑定的处理函数。
3. 处理函数的输入包括一个 Request 和一个 Response。
4. Request 对象用来获取 Path 和 Query参数，Headers以及 Request Body（XML，JSON，...）。
5. Response 对象用来设置 Status，Headers，以及 Response Body
6. Request 和 Response 对象都可以使用标准库来在对象和XML或JSON之间进行转换。

我们可以使用一个简单的例子来演示上面的过程。一个用来对User 对象进行 CRUD 操作的 WebService。
我们首先在一个 userservice 目录里面创建一个 userservice.go 文件。

```
package userservice

import (
    "github.com/emicklei/go-restful"
    "log"
)

type User struct {
    Id, Name string
}
```

User 类型代表我们要操作的对象。

文件中的下一个部分就是 WebService 的 API 定义。这些 API 是一组Route对象的集合，这些Route定义了如何将进来的 HTTP 请求映射到对应的处理函数。

```
func New() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/users").
        Consumes(restful.MIME_XML, restful.MIME_JSON).
        Produces(restful.MIME_XML, restful.MIME_JSON)
        
    service.Route(service.GET("/{user-id}").To(FindUser))
    service.Route(service.POST("").To(UpdateUser))
    service.Route(service.PUT("/{user-id}").To(CreateUser))
    service.Route(service.DELETE("/{user-id}").To(RemoveUser))
        
    return service
}
```

首先，使用一个 root URL 来初始化所有路径的 service，定义每个 Route 可以接收的MIME类型，以及可以响应的MIME类型。当然这个也可以针对每个Route单独指定。然后，service 指定它可以提供哪些路径。这里的函数调用 `GET("/{user-id}")` 是 `Method("GET").Path("/{user-id}")` 的简单写法，这个方法调用创建一个 RouteBuilder 对象。然后使用这个 RouteBuilder 对象指定对应的处理函数。

下面，就是定义每个 Route 的处理函数了。

```
func FindUser(request *restful.Request, response *restful.Response) {
    id := request.PathParameter("user-id")
    // here you would fetch user from some persistence system
    usr := &User{Id: id, Name: "John Doe"}
    response.WriteEntity(usr)
}
```

Route 的处理函数的方法声明都一样，包含一个 Restful 的 Request 和 Response，两个一对。 Request 是 http Request 对象的封装，提供了一些方便的方法。Response 是对 http ResponseWriter 的封装。这种设计方式可以将底层的 HTTP 结构开放给开发者，同时也为开发者提供了一些通用的 Restful 函数，例如 WriteEntity。WriteEntity 函数会检查请求的 Accept 头部来决定 response 的 Content-Type 头部，同时也决定了使用那种方法来序列化对象（这里就是 User 对象）。

userservice.go 文件的其他内容就是剩下的Route处理函数的定义。

```
func UpdateUser(request *restful.Request, response *restful.Response) {
    usr := new(User)
    err := request.ReadEntity(&usr)
    // here you would update the user with some persistence system
    if err == nil {
        response.WriteEntity(usr)
    } else {
        response.WriteError(http.StatusInternalServerError,err)
    }
}

func CreateUser(request *restful.Request, response *restful.Response) {
    usr := User{Id: request.PathParameter("user-id")}
    err := request.ReadEntity(&usr)
    // here you would create the user with some persistence system
    if err == nil {
        response.WriteEntity(usr)
    } else {
        response.WriteError(http.StatusInternalServerError,err)
    }
}

func RemoveUser(request *restful.Request, response *restful.Response) {
    // here you would delete the user from some persistence system
}
```

现在，我们已经完成了 UserService 的定义和实现。下面的代码段演示了如何在一个应用中使用这个 service。

```
package main

import (
    "github.com/emicklei/go-restful"
    "log"
    "net/http"
    "userservice"
)

func main() {
    restful.Add(userservice.New())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

服务启动之后，我们可以利用下面的方法测试：

1. 默认的请求

```
$ curl http://localhost:8080/users/12
<?xml version="1.0" encoding="UTF-8"?>
 <User>
  <Id>12</Id>
  <Name>John Doe</Name>
 </User>
```

2. 带 Accpet 头部的请求

```
$ curl http://localhost:9090/users/12 -H 'Accept: application/json'
{
 "Id": "12",
 "Name": "John Doe"
}
```

3. 新建一个User对象

```
$ curl http://localhost:9090/users -X POST -d '{"Id":"32","Name":"jemygraw"}' -H 'Content-Type: application/json'
<?xml version="1.0" encoding="UTF-8"?>
 <User>
  <Id>32</Id>
  <Name>jemygraw</Name>
 </User>
```

4. 新的一个User对象，要求返回JSON。

````
$ curl http://localhost:9090/users -X POST -d '{"Id":"32","Name":"jemygraw"}' -H 'Content-Type: application/json' -H 'Accept: application/json'
{
 "Id": "32",
 "Name": "jemygraw"
}
```

---

译者：

本文永久保存地址为：https://github.com/jemygraw/TechDoc/blob/master/Go%E5%BA%93%E5%AD%A6%E4%B9%A0/2018-09-14%20k8s%20%E4%BE%9D%E8%B5%96%E5%BA%93%E4%B8%AD%E7%9A%84go-restful%E6%A1%86%E6%9E%B6.md

