# Kubernetes 中如何开发一个 kubectl 的插件命令

## 背景

在日常使用中，Kubectl 作为和 Kubernetes 集群进行交互的工具，提供了丰富的功能。但是偶尔也有时候，你想做一些 Kubectl 暂时还不支持的功能。那么在这种情况下，如何不改变 Kubectl 的代码并且重新编译就能引入新的功能呢？ 这个问题的答案就是采用 Kubectl 的 Plugin 机制。

Kubectl 的 Plugin 机制在 v1.8.0 版本的时候就引入了，并且在 v1.12.0 版本中进行了大规模的重构以适应更加复杂多样的场景，并且最终在 v1.14.0 版本中稳定下来。所以你必须使用 Kubectl v1.12.0 及以上版本才可以支持当前的插件命令。

## 插件命令

所谓的插件命令其实很简单，只要符合以下几个特点即可：

(1) 该命令是一个可执行的文件；  
(2) 该命令能够通过 `$PATH` 搜索到，也就是说如果需要，你必须把这个命令加入到 `$PATH` 中；  
(3) 该命令必须以 `kubectl-` 开头，例如 `kubectl-echo` 就是一个合法的插件命令名称。  

基于以上的要求，我们可以用任何语言去编写这个命令，比如我们最简单的用 C 语言写一个 `kubectl-hello` 的插件命令尝试下。

```c
#include <stdio.h>
int
main(int argc, char *argv[])
{
	printf("hello, i am a kubelet plugin command\n");
}
```

然后我们编译一下：

```
$ gcc -o kubectl-hello kubectl-hello.c
```

然后我们把这个命令所在的目录放到系统的 $PATH 变量中，最后通过 kubectl 命令尝试下。

```
$  kubectl hello
hello, i am a kubelet plugin command
```

通过上面的输出我们可以看到，这个插件命令已经成功完成了，那么剩下来就是利用你熟悉的语言来编写二进制工具来满足你的需求了。

## 发现插件

Kubectl 提供了一个 plugin 的命令，该命令可以使用子命令 list 来列举当前系统中的插件命令。具体的搜索方法如下：

(1) 搜索系统的 $PATH 中指定的所有的目录，查找所有以 `kubectl-` 开头的文件；  
(2) 如果搜索到的匹配以 `kubectl-` 开头的文件是可执行文件，那么会按照顺序作为插件命令输出；如果不是可执行文件，也会输出，但是同时会输出一个 `Warning` 的信息；  

## 当前限制

虽然我们可以自定义插件命令，但是有个限制就是**你无法定义一个 kubectl 已经存在的命令去试图覆盖原命令的行为**。例如 `kubectl-version` 这样的命令永远不会被执行，因为 kubectl 会优先执行内置的 version 命令。基于这样的原因，你也无法给已有的命令增加额外的子命令。

## 使用插件

插件命令不需要安装，也不需要预加载任何东西。它继承 kubectl 命令的执行环境。kubectl 通过插件命令的名称来执行它。例如对于上面的名为 `kubectl-hello` 的命令，kubectl 就通过 `$ kubectl hello` 来执行它。

对于插件命令来讲，它接收到的第一个参数总是它文件所在的全路径。对于上面的 `kubectl-hello` 命令，我们稍作修改，用来打印所有的参数。

```c
#include <stdio.h>
int
main(int argc, char *argv[])
{
	int		i = 0;
	printf("hello, i am a kubelet plugin command\n");
	printf("\n");
	for (; i < argc; i++) {
		printf("%s\n", argv[i]);
	}
}
```

输出如下：  

```
$ kubectl hello kubernetes
hello, i am a kubelet plugin command

/Users/jemy/Bin/k8s-plugins/kubectl-hello
kubernetes
```

## 插件命名

对于插件的命令，必须了解的两点如下：

(1) 插件命令支持子命令，其格式必须为 `kubectl-cmd-cmd1-cmd11` ，也就是每个命令通过 `-` 分隔。这样在调用的时候可以使用 `$ kubectl cmd cmd1 cmd11` 这样的方式来调用。  
(2) 如果要在插件命令中使用多个单词构成一个命令，那么多个单词必须用 `_` 进行分隔，例如对于 `kubectl-hello_world` 命令，可以通过 `$ kubectl hello_world` 这样的方式来调用。  
(3) 插件命令必须自行解析所有传给该命令的选项参数，并进行相应的处理。

## 插件管理

鉴于 kubernetes 本身并没有提供插件命令的包管理器用来安装和更新插件命令，我们可以使用 Kubernetes-sigs 项目中的 krew 来完成相关工作。 


参考文档：[https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)  

