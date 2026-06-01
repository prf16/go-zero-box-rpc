## 概述

### 简介

go-zero-box-rpc 是 go-zero-box 体系中专门面向 rpc 服务的工程模板方案。

它基于 go-zero 官方 rpc 能力（gRPC + zrpc），在此之上进一步沉淀了实战中验证过的 rpc 工程结构。

同样 go-zero-box 拥有的能力依然在 rpc 服务内置了，比如 scheduler、queue、script 服务。

配套的 HTTP API 服务模板请参考 [go-zero-box](https://github.com/prf16/go-zero-box)，两者配合使用构成完整的微服务体系。

如果你正在使用 go-zero 构建多服务架构，并希望 rpc 服务：

 - 结构清晰

 - 职责单一

 - 易于维护和扩展

那么 go-zero-box-rpc 将非常适合你。

### 环境要求
- Golang >= 1.23
- MySQL >= 5.7
- Redis

### 代码结构

```text
.
├── api                                protobuf 描述文件
│   └── user/                          用户服务 proto 定义及生成的 pb 文件
├── app                                包含应用程序的主要代码
│   ├── etc                            静态配置文件目录
│   ├── internal                       内部业务逻辑
│   │   ├── config/                    配置结构定义
│   │   ├── logic/                     业务逻辑层（按模块分子目录）
│   │   │   ├── hello/                 示例逻辑
│   │   │   └── user/                  用户逻辑（登录、注册、信息查询）
│   │   ├── server/                    gRPC Server 实现（按模块分子目录）
│   │   │   ├── hello/                 Hello 服务
│   │   │   └── user/                  User 服务
│   │   └── svc/                       服务上下文与依赖注入
│   │       ├── command/               CLI 脚本命令（可注册为计划任务）
│   │       ├── model/                 数据模型层（usermodel、messagemodel）
│   │       ├── queue/                 异步队列消费者（邮件、短信、微信）
│   │       ├── services/              业务服务层
│   │       ├── utils/                 工具函数（加密、响应、常量等）
│   │       └── service_context.go     服务上下文，聚合所有依赖
│   ├── app.go                         应用程序的入口文件，定义了 rpc、scheduler、queue、script 服务
│   ├── user.go                        gRPC 服务注册入口
│   ├── wire.go                        依赖注入定义文件
│   └── wire_gen.go                    依赖注入生成文件（自动生成，勿手动编辑）
├── build                              项目构建目录
│   ├── app                            应用构建后的目录
│   │   ├── etc                        配置文件目录
│   │   └── app                        应用二进制文件
│   └── app.tar                        应用构建后的打包文件
├── deploy                             部署相关目录
│   ├── access                         示例图片
│   ├── goctl                          goctl 代码生成模板文件
│   └── sql                            初始化数据库SQL
├── pkg                                工具包（可跨项目复用）
│   ├── asynqx/                        Asynq 异步队列封装（客户端、队列、调度器）
│   ├── cobrax/                        Cobra 命令扩展类型定义
│   ├── database/                      数据库连接封装（支持多库）
│   └── redis/                         Redis 连接封装（支持多实例）
├── runtime                            项目运行时目录（日志等）
├── vendor                             项目依赖包
├── .gitignore                         git 忽略文件
├── go.mod                             项目依赖管理文件
├── Makefile                           项目构建文件
└── README.md                          项目说明文件
```

## 安装开发工具

### wire

wire 是一个依赖注入工具，用于解决 Go 语言中依赖注入的问题。通过 wire，您可以轻松地管理应用程序的依赖关系，并确保它们在编译时进行注入。

```shell
# shell 安装
$ go install github.com/google/wire/cmd/wire@latest

# 验证安装
$ wire
```

### goctl 安装

goctl 是 go-zero 微服务框架下的代码生成工具。使用 goctl 可显著提升开发效率，让开发人员将时间重点放在业务开发上，其功能有：api服务生成、rpc服务生成、model代码生成、模板管理。

```shell
# 方式一（推荐）：shell 安装
$ go install github.com/zeromicro/go-zero/tools/goctl@v1.9.2

# 方式二：手动下载安装
https://github.com/zeromicro/go-zero/releases/tag/tools%2Fgoctl%2Fv1.9.2

# 验证安装
$ goctl --version
```

## Make 命令介绍

Makefile 文件描述了项目工程的编译规则，只需要一个 `make build` 命令，整个工程就开始自动构建项目环境，不再需要手动执行大量的 `go build` 命令，Makefile 文件定义了一系列规则，指明了源文件的编译顺序、依赖关系、是否需要重新编译等，可以输入 `make help` 查看命令集。

```shell
# 查看 make 信息
$ make

# 构建并打包应用（根据 env=dev|test|prod 编译，生成 build/app 及 app.tar）
$ make build env=dev

# 根据 app.proto 定义生成 Go RPC 代码
$ make rpc

# 根据 wire.go 生成依赖注入代码（wire_gen.go）
$ make wire

# 根据 MySQL 表结构生成 Go Model 代码
$ make model
```

## 服务启动命令

项目通过 Cobra 实现了多服务管理，支持以下子命令：

```shell
# 启动 RPC 服务
$ ./app server:rpc

# 启动队列消费服务
$ ./app server:queue

# 启动计划任务服务
$ ./app server:scheduler

# 启动全部服务（单体模式，包含 rpc、队列、计划任务）
$ ./app server:all

# 执行自定义脚本命令（示例）
$ ./app hello:world
```

开发阶段可以通过 `-conf` 参数指定配置文件路径：

```shell
$ go run main.go server:rpc
```

## 配置文件说明

配置文件位于 `app/etc/app.yaml`，各配置段说明如下：

| 配置段 | 说明 |
|--------|------|
| `Server` | go-zero RPC 服务配置，包括服务名称、监听地址、超时时间、日志配置等 |
| `Redis` | Redis 连接信息，包括地址、类型（node/cluster）、密码 |
| `Database` | MySQL 数据库连接字符串，支持多库配置 |

## 快速开始示例
参考 [go-zero-box 快速开始示例](https://github.com/prf16/go-zero-box?tab=readme-ov-file#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B%E7%A4%BA%E4%BE%8B)

## 常见问题
### 1. go mod tidy 超时 i/o timeout
```
1. 确认当前shell
echo $SHELL

2. 编辑相应的 Shell 配置文件
a. 如果使用 zsh
vim ~/.zshrc

a. 如果使用 bash
vim ~/.bash_profile

3. 添加配置信息
export GOPROXY=https://proxy.golang.org,https://mirrors.aliyun.com/goproxy/,direct

4. 重载配置
source ~/.zshrc
source ~/.bash_profile

5. 查看更新结果
go env GOPROXY
```

## 许可证

本项目采用 Apache License 2.0 许可证 - 查看 LICENSE 文件了解详情