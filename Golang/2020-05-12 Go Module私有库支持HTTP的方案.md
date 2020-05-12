## 背景

假设有个 ABC 公司，在内网部署了一个 gitlab 的服务，域名是 http://git.abc.com 。他们的程序员需要使用 Go Module 的功能。
由于有些公共代码是公司内部的，而有些依赖则是 Github 上开源的，另外为了提高 Go Module 下载速度，他们还打算用阿里云的加速服务。

问题来了，Go Module 默认支持的是 https 的域名，所以内网的 gitlab 服务就比较麻烦，也就是说内网的 gitlab 的库下载的时候要满足两点需求，一个是支持 http 协议，另外一个是不使用阿里云的代理。


## 方案

首先升级 Git 版本，当然也可以试试下面的步骤先，如果编译过程下载私有库的时候没有出现 `unknown revision` 的错误，就不需要升级 Git。升级 Git 肯定都会，贴一个 Centos 7 的升级脚本。

```
wget https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.26.2.tar.gz
yum install libcurl-devel
yum install libcom_err-devel
yum install libopenssl-devel
CPPFLAGS="-I/usr/include/" LDFLAGS="-L/usr/lib64" ./configure --prefix /usr/local/git/
make && make install
```

其次，我们需要引入两个环境变量来支持私有库的 http 协议和忽略使用阿里云的代理。

```
export GOINSECURE=git.abc.com
export GONOPROXY=git.abc.com
```

这样的话，整个配置如下：

```
# set GOPATH
export GOPATH=/home/jenkins/go/GoProjects
 
# enable Go module
export GO111MODULE=on
export GOINSECURE=git.abc.com
export GONOPROXY=git.abc.com
export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

上面的内容搞定之后，如果你下载私有库的时候还会出现下面的错误，要求提供鉴权。

```
go: finding module for package git.abc.com/devops/base/wechat
go: finding module for package git.abc.com/devops/base/storage
conf/config.go:6:2: module git.abc.com/devops/base/storage: git ls-remote -q origin in /home/jenkins/go/GoProjects/pkg/mod/cache/vcs/c93a7972976a1967065d0c6e6991f8c13a161f1a51c5aa2079b1724afad589a1: exit status 128:
	fatal: could not read Username for 'http://git.abc.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
conf/config.go:7:2: module git.abc.com/devops/base/wechat: git ls-remote -q origin in /home/jenkins/go/GoProjects/pkg/mod/cache/vcs/c93a7972976a1967065d0c6e6991f8c13a161f1a51c5aa2079b1724afad589a1: exit status 128:
	fatal: could not read Username for 'http://git.abc.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
Build step 'Execute shell' marked build as failure
```

那么，需要设置下面的内容。

```
git config --global http.extraheader 'PRIVATE-TOKEN: <AccessToken>'
git config --global url."git@git.abc.com:".insteadOf "http://git.abc.com/"
```

其中那个 <AccessToken> 替换为 Gitlab 里面的访问 AccessToken即可，自己生成一个。

最后，就可以使用 go build 啦。