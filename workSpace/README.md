# Go 工作区模式
在 Go1.18 之前，比较流行的「工作区（workspace）模式」为 Module 工作区模式

Go1.18 正式发布后，工作区模式有点变化(工作区模式真的很有用)，工作区模式（Workspace mode），可不是之前 GOPATH 时代的 Workspace，而是希望在本地开发时支持多 Module

## 一.传统模式
### 1.1 操作流程
````shell
$ mkdir mytest example
$ cd mytest/
$ go mod init github.com/will-yinchengxin/mytest
$ cd ../example/
$ go mod init github.com/will-yinchengxin/example
````

> #### 在 mypkg/bar.go 中增加如下示例代码
````go
package mytest

func Bar() {
	println("This is package mytest")
}
````

> #### 在 example/main.go 中新增
````go
package main

import "github.com/will-yinchengxin/mytest"

func main() {
	mytest.Bar()
}
````
这时候，如果我们运行 go mod tidy，肯定会报错，因为我们的 mypkg 包根本没有提交到 github 上，肯定找不到。go run main.go 也就不成功。

我们当然可以提交 mytest 到 github，但我们每修改一次 mytest，就需要提交（而且每次提交之后需要在 example 中 go get 最新版本），否则 example 中就没法使用上最新的。

针对这种情况，目前是建议通过 replace 来解决，即在 example 中的 go.mod 增加如下 replace：（v1.0.0 根据具体情况修改，还未提交，可以使用 v1.0.0）

````
module github.com/will-yinchengxin/example

go 1.18

require gitlab.com/will-yinchengxin/mytest v1.0.5

replace gitlab.com/will-yinchengxin/mytest v1.0.5 => gitlab.com/will-yinchengxin/mytest v1.0.4-0.20221101103853-069e3897c811
````
当都开发完成时，我们需要手动删除 replace，并执行 go mod tidy 后提交，否则别人使用就报错了。 如果本地有多个 module，每一个都得这么处理(有点不方便)。

## 二.工作区模式
针对上面的这个问题，Michael Matloob 提出了 Workspace Mode（工作区模式）。相关 issue 讨论：cmd/go: add a workspace mode，
这里是 Proposal。并在 Go1.18 中发布了（因此，要使用工作区，请确保 Go 版本在 1.18+）。
````shell
$ go version
go version go1.18 darwin/arm64

$ go help work
Work provides access to operations on workspaces.

Note that support for workspaces is built into many other commands, not
just 'go work'.

See 'go help modules' for information about Go's module system of which
workspaces are a part.

See https://go.dev/ref/mod#workspaces for an in-depth reference on
workspaces.

See https://go.dev/doc/tutorial/workspaces for an introductory
tutorial on workspaces.

A workspace is specified by a go.work file that specifies a set of
module directories with the "use" directive. These modules are used as
root modules by the go command for builds and related operations.  A
workspace that does not specify modules to be used cannot be used to do
builds from local modules.

go.work files are line-oriented. Each line holds a single directive,
made up of a keyword followed by arguments. For example:

        go 1.18

        use ../foo/bar
        use ./baz

        replace example.com/foo v1.2.3 => example.com/bar v1.4.5

The leading keyword can be factored out of adjacent lines to create a block,
like in Go imports.

        use (
          ../foo/bar
          ./baz
        )

The use directive specifies a module to be included in the workspace's
set of main modules. The argument to the use directive is the directory
containing the module's go.mod file.

The go directive specifies the version of Go the file was written at. It
is possible there may be future changes in the semantics of workspaces
that could be controlled by this version, but for now the version
specified has no effect.

The replace directive has the same syntax as the replace directive in a
go.mod file and takes precedence over replaces in go.mod files.  It is
primarily intended to override conflicting replaces in different workspace
modules.

To determine whether the go command is operating in workspace mode, use
the "go env GOWORK" command. This will specify the workspace file being
used.

Usage:

        go work <command> [arguments]

The commands are:

        edit        edit go.work from tools or scripts
        init        initialize workspace file
        sync        sync workspace build list to modules
        use         add modules to workspace file

Use "go help work <command>" for more information about a command.

````
注意：上文中提到，工作区不只是 go work 相关命令，Go 其他命令也会涉及工作区内容，比如 go run、go build 等。

根据这个提示，我们初始化 workspace：
````
├── example
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── go.work
└── mytest
    ├── bar.go
    └── go.mod
````    
注意几点：
- 多个子模块应该在一个目录下（或其子目录）。比如这里的 gowork 目录；（这不是必须的，但更好管理，否则 go work init 需要提供正确的子模块路径）
- go work init 需要在 gowork 目录执行；
- go work init 之后跟上需要本地开发的子模块目录名；

打开 go.work 看看长什么样：
````
go 1.18

use (
	./example
	./mytest
)
````
go.work 文件的语法和 go.mod 类似（go.work 优先级高于 go.mod），因此也支持 replace。

注意：实际项目中，多个模块之间可能还依赖其他模块，建议在 go.work 所在目录执行 go work sync。

现在，我们将 example/go.mod 中的 replace 语句删除，再次执行 go run main.go（在 example 目录下），得到了正常的输出。也可以在 gowork 目录下，这么运行：go run example/main.go，也能正常。

注意，go.work 不需要提交到 Git 中，因为它只是你本地开发使用的。

当你开发完成，应该先提交 mytest 包到 GitHub，然后在 example 下面执行 go get：
````
go get -u gitlab.com/will-yinchengxin/mytest@v1.0.5
````

然后禁用 workspace（通过 GOWORK=off 禁用），再次运行 example 模块，是否正确：
````
$ cd ./gowork/example
$ GOWORK=off go run main.go
````

在 GOPATH 年代，多 GOPATH 是一个头疼的问题。当时没有很好的解决，Module 就出现了，多 GOPATH 问题因此消失。但多 Module 问题随之出现。Workspace 方案较好的解决了这个问题

## Tips
#### 关于 git tag
发布版本携带版本号，通过打 tag 的方式实现
````
git tag v1.0.0   // 轻量级tag
git tag -a v1.0.0 -m "有注释"   // 附带注释tag

git tag /  git tag --list  // 查看tag列表(在 .git/refs/tags 即可查看所有内容, 且 tag list 展示结果不受分支影响)

git	tag	-l v1.0.0 // tag 查询

git push -u origin main // 要将标签推送到远程仓库，首先要建立本地仓库与远程仓库的联系，比如可以采用：

git push origin v1.0.0 // 推送指定的本地标签到远程仓库, 例如将本地master分支上的标签v1.0推送到远程仓库

git push origin --tag // 该方法可以一次性推送所有的本地标签到远程仓库

git push origin :v1.0.0  // 删除远程标签

git tag -d v1.0.0     // 删除本地标签

git checkout v1.0.0  // 切换到指定分支
````



