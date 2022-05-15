# Wire 工程化实践

在 wire 工程化中， 我主要参考了以下项目对于该 demo 进行实现，为了便于观看，部分内容做了精简。

**项目布局参考：**

- layout-project：https://github.com/golang-standards/project-layout
- Kratos框架：https://github.com/go-kratos/kratos

**wire实践参考：**

- 基于Kratos框架的微服务项目：https://github.com/go-kratos/beer-shop

## Demo详解

该demo项目布局

```text
_
│  go.mod
│  go.sum
├─cmd
│      main.go
│      wire.go
│      wire_gen.go
├─config
│      app.json
└─internal
    ├─biz
    │      order.go
    ├─data
    │      data.go
    │      order.go
    └─server
        ├─config
        │      config.go
        └─db
                mysql.go
```

### internal 目录详解

该目录存放所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用

**internal/biz**

业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。

**internal/data**

业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。

**internal/server**

config，db，http，grpc 实例的创建和配置

### 工程化实践



