这个帖子主要是讲如何在本地 Docker 里面跑一个 phpmyadmin 来访问安装在本机的 MySQL 8.0 数据库。

注意这里的数据库版本是 MySQL 8.0，因为这个版本的有些命令和之前的版本已经不同了，有微小而合理的调整。

## 拉取 phpmyadmin 镜像到本地

```
$ docker pull phpmyadmin/phpmyadmin
```

这个镜像是 phpmyadmin 的官方镜像。

## 配置本地的 MySQL 允许远程访问

因为 phpmyadmin 是跑在 Docker 里面的，而 MySQL 是跑在本地的物理机器上面的，所以对于 phpmyadmin 来说，这个 MySQL 服务器就是个远程的服务器。
我们需要设置允许 phpmyadmin 远程访问这个数据库服务器。

很多时候我们图开发方便，很可能会直接用下面的命令来开放 MySQL 数据库的远程访问。

```
mysql> grant all privileges on *.* to root@'%' with grant option;
```

但是这个命令在 MySQL 8.0 上面会有如下的报错：

>ERROR 1410 (42000): You are not allowed to create a user with GRANT

最主要是 MySQL 8.0 版本不再允许 root 远程访问了。在这种情况下，我们就要创建自己的用户来配置远程访问。

```
$ create user 'duoke' identified by 'duoke';
$ grant all privileges on clouddeploy.* to 'duoke'@'%' with grant option;
$ flush privileges;
```

上面我们创建了一个名称为 `duoke` 的用户，密码也是 `duoke`，然后允许它远程访问数据库 `clouddeploy`。

## 启动 phpmyadmin 的 Docker 容器

上面的准备工作做好之后，我们就可以启动 phpmyadmin 的容器了，在启动容器之前，我们需要解决一个问题就是 phpmyadmin 的容器是通过环境变量的方式来指定要访问的 MySQL 服务器的地址的，可以指定为域名，主机名或者直接IP地址。我们一般开发的时候电脑会从公司带到家里，再从家里带到公司，一般公司和家里分配的动态IP地址是不一样的，所以选择指定IP肯定不合适。所以这个地方我们可以指定为一个主机名，这个主机名并不一定需要是本地主机名，更不可以是 localhost。为了方便识别，我们就叫它 phpmydmin-mysql-server ，然后在本地的 `/etc/hosts` 文件里面添加主机名到本地机器IP的地址映射。

```
#192.168.1.102 phpmyadmin-mysql-server
10.120.112.39 phpmyadmin-mysql-server
```

然后启动容器的命令如下：

```
$ docker run -d -e PMA_HOST=phpmyadmin-mysql-server -p 9090:80 phpmyadmin/phpmyadmin
```

然后我们就可以通过地址 http://localhost:9090 访问到 phpmyadmin了，使用上面创建的用户名和密码登陆即可。

这样，我们每次到公司就切换 `/etc/hosts` 文件里面的IP映射就可以了。
