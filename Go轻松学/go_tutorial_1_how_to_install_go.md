# Go语言环境安装与测试
## 安装
现在来谈谈Go语言的安装，要使用Go来编写程序首先得把环境搭建起来。
Go的语言环境搭建还是比较简单的👌。Google提供了Windows和Mac的安装包，所以去下载一下安装就可以了。
对于Linux的系统，可以使用系统提供的包安装工具来安装。

**Go的下载地址**  

[https://code.google.com/p/go/downloads/list](https://code.google.com/p/go/downloads/list)

**Windows**

对于Windows系统，Go提供了两种不同的安装包，分别对应32位的系统和64位的系统，安装的时候根据自己的系统实际情况选择下载包。Windows下面提供的是msi格式的安装包，这种包是可执行文件，直接双击安装就可以了。安装完成之后，安装程序会自动地将安装完的Go的根目录下的bin目录加入系统的PATH环境变量里面。所以直接打开命令行，输入go，就可以看到一些提示信息了。

**Mac**

如果是新买的Mac，里面可能自带了一个go的可执行文件，在路径`/etc/paths.d/`下面，就是一个go可执行文件。如果我们需要安装从官网下载的dmg安装包，先要把这个文件删除掉。可以用`sudo rm /etc/paths.d/go`来删除。然后自动安装dmg之后，要使用`export PATH`的方法将安装好的Go目录下面的bin目录加入PATH中。一般安装完之后路径为`/usr/local/go`，所以你可以用下面的方法：
首先切换到自己的用户目录

	cd ~
	
然后

	vim .profile

加入一行

	export PATH=/usr/local/go/bin:$PATH

就可以了。


**Linux**

Linux的发行版有很多，可以根据不同系统提供的包管理工具来安装Go，不过可能系统包管理工具提供的不是最新的Go版本。在这种情况下，你可以去下载最新的tar包。
然后使用下面的方法

	sudo tar -C /usr/local -xzf go1.2.linux-386.tar.gz

如果是64位的系统，用下面的方法

	sudo tar -C /usr/local -xzf go1.2.linux-amd64.tar.gz

当然，这样的方式只是将安装包解压拷贝到`/usr/local/`下面。你还需要使用`export PATH`的方式将Go的bin目录加入PATH。
方法和上面Mac介绍的一样。
另外如果你不是将Go安装到`/usr/local`目录下面，你还需要设置一个GOROOT环境变量。比如你安装到你自己的文件夹下面，比如叫jemy的用户的路径是`/home/jemy`，那么你安装到这个目录的Go路径为`/home/jemy/go`，那么在`export PATH`之前，你还需要使用下面的命令。

	export GOROOT=/home/jemy/go

总结一下，如果你默认安装路径为`/usr/local/go`，那么只需要用

	export PATH=$PATH:/usr/local/go/bin

就可以了。
如果不是默认路径则需要这样

	export GOROOT=/home/jemy/go
	export PATH=$PATH:/$GROOT/bin

上面的`/home/jemy`是根据实际安装的路径情况来确定。

最后说一下go的最基本的三个命令

1.查看版本号

	go version 

结果为

	duokr:~ jemy$ go version
	go version go1.2 darwin/386

2.格式化go代码文件

	go fmt file_name.go

3.运行单个go代码文件

	go run file_name.go

## 测试
 
`生` `死` `hello world`

学习计算机的， 绕不开的三件事。

有谁安装好语言环境，不试一下hello world的？


	//main包, 凡是标注为main包的go文件都会被编译为可执行文件
	package main

	//导入需要使用的包
	import (
		"fmt" //支持格式化输出的包,就是format的简写
	)

	//主函数,程序执行入口
	func main() {
    	/*
        	输出一行hello world
        	Println函数就是print line的意思
    	*/
    	fmt.Println("hello world")
	}
然后使用`go run helloworld.go`来运行这个例子。如果安装成功，那么会输出一行`hello world`。

*PS*

`Windows7可以在文件所在目录下面使用Shift+右键，快速打开已定位到所在目录的命令行窗口。直接输入上面命令即可。`
