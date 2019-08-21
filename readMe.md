# go mod 的使用和总结

`go modules` 是 golang 1.11 新加的特性。

## go mod 的说明

模块是相关 Go 包的集合。modules 是源代码交换和版本控制的单元。 go 命令直接支持使用 modules，包括记录和解析对其他模块的依赖性。
modules 替换旧的基于 GOPATH 的方法来指定在给定构建中使用哪些源文件。

使用 Modules:

1. 把 golang 升级到 1.11 或以上
2. 设置 GO111MODULE, 环境变量, `export GO111MODULE=on`

## go mod 的简单使用

在 GOPATH 目录之外新建一个目录，并使用 go mod init 初始化生成 go.mod 文件

```bash
  mkdir go-mod-demo
  cd go-mod-demo

  # go mod init moduleName, 接受一个参数：moduleName 代表项目在本地的引用名称，在项目中页可以引用
  go mod init demo1

```

### go mod 添加其它包

通过 go mod 添加其它包有几种方式

#### 1. `go run ...`, `go build` 自动添加

在项目文件中 `import` 第三方库后，执行 `go run ...` 会自动去搜索和下载第三方库，并添加到 go.mod, go.sum 文件中

#### 2. go get 命令手动添加

新版 go get 可以在末尾加 @ 符号，用来指定版本

```bash
go get github.com/gorilla/mux    # 匹配最新的一个 tag
go get github.com/gorilla/mux@latest    # 和上面一样
go get github.com/gorilla/mux@v1.6.2    # 匹配 v1.6.2
go get github.com/gorilla/mux@e3702bed2 # 匹配 v1.6.2
go get github.com/gorilla/mux@c856192   # 匹配 c85619274f5d
go get github.com/gorilla/mux@master    # 匹配 master 分支

```

### go replace 替换无法直接获取的 package

并不是所有的 package 都能成功下载, 可用使用 go replace 替换;

modules 可以通过在 go.mod 文件中使用 replace 指令替换成 github 上对应的库，比如：

1. 修改 go.mod 文件， 添加 replace 配置项

```mod
## go.mod 文件中, 使用 replace 替换无法成功下载的库，可以指定版本
module learn1

go 1.12

replace (
  golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
  golang.org/x/sys => github.com/golang/sys v0.0.0-20190726091711-fc99dfbffb4e

  ## latest 自动下载最新版本
  golang.org/x/text => github.com/golang/text latest
)

require github.com/sirupsen/logrus v1.4.2
```

2. 执行 go mod tidy 拉取缺少的模块，移除不用的模块

然后执行 go mod tidy.这个命令会把 latest 自动替换成最新的版本号

## go mod 后项目中的引用改变

使用 go mod 后，引用项目内部的包，需要使用当前项目的包名作为第一个引用

```go
// 1. 外部的包无效改变
// 2. 内部包的互相引用, 这里 go-web 是当前项目的包名
import (
  "flag"
  "fmt"
  "go-web/demo1"
  "go-web/demo10"
)
```

## go mod 常用指令的说明

### 0. go mod int name

使用包名 name 初始化项目

### 1. go get 获取版本，升级版本等，相当于 npm install

`go get ....` 会自动更新 go.mod, go.sum 文件

## 2. go list -m all 查看依赖

查看 module 下的所有依赖

## 3. go mod tidy

清理无用的依赖

### 5. go mod vendor 同步 GOPATH/src

直接使用这个命令就可以把 GOPATH/src 目录下的依赖包同步到当前项目目录中
