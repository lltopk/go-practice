下载解压正确的arch压缩包版本, Windows对应arm64

https://golang.google.cn/dl/

解压并配置系统环境变量
```
GOROOT: C:\Users\hasee\.go\go1.25.5.windows-amd64
PATH: %GOROOT%\bin
```

验证环境变量
```
go 
go version
go env
```

查看go变量
```
go env
```
设置go变量
```
## grammar
go env -w KEY=
## 配置代理
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
## 配置模块支持
go env -w GO111MODULE=on
```

vscode安装go插件, 打开go插件, 安装go工具, 全部安装
- gopls
- gotests
- impl
- goplay
- dlv(delve)调试工具
- staticcheck
```
ctrl shift p
go:install/update tools
```


几个关键的go变量
```
## 你修改的go变量就保存在下面的用户文件中
GOENV=C:\Users\hasee\AppData\Roaming\go\env
## go工具默认安装目录
GOPATH=C:\Users\hasee\go
## go源码安装位置
GOROOT=C:\Users\hasee\.go\go1.25.5.windows-amd64\go
```

调试运行
- F5