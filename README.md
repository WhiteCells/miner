## Miner

### 项目结构

``
├── assets # 资源文件，无关项目
├── cmd # 简化 main，暂不处理
├── common # 分离出来的一些模块，使其他模块专注自身职责
│   ├── dto # 数据对象，request 和 response
│   │   ├── farm.go
│   │   ├── miner.go
│   │   └── user.go
│   ├── perm # 基础权限信息
│   │   └── perm.go
│   ├── points # 积分类型信息
│   │   └── points.go
│   └── status # 状态信息
│       └── status.go
├── controller # 控制器
│   ├── farm.go
│   ├── flightsheet.go
│   └── user.go
├── dao # 数据访问对象
│   ├── mysql # 数据库访问对象
│   │   ├── 2_farm_miner.go
│   │   ├── 2_user_farm.go
│   │   ├── 2_user_miner.go
│   │   ├── farm.go
│   │   ├── flightsheet.go
│   │   ├── miner.go
│   │   ├── oper_log.go
│   │   ├── points_record.go
│   │   ├── user.go
│   │   └── wallet.go
│   └── redis # redis 访问对象
│       ├── farm.go
│       ├── flightsheet.go
│       ├── miner.go
│       └── user.go
├── logs # 日志文件夹
├── middleware # 中间件
│   ├── auth.go
│   ├── ip.go
│   ├── logger.go
│   └── session.go
├── model # 数据库模型
│   ├── 2_farm_miner.go
│   ├── 2_flightsheet_wallet.go
│   ├── 2_miner_flightsheet.go
│   ├── 2_user_farm.go
│   ├── 2_user_miner.go
│   ├── 2_user_wallet.go
│   ├── farm.go
│   ├── flightsheet.go
│   ├── miner.go
│   ├── oper_log.go
│   ├── points_record.go
│   ├── user.go
│   └── wallet.go
├── route # 路由
│   └── user.go
├── service # 业务服务
│   ├── farm.go
│   ├── flightsheet.go
│   ├── miner.go
│   ├── oper_log.go
│   ├── user.go
│   └── wallet.go
├── utils # 其他模块（乱七八糟的模块）
│   ├── captcha.go # 验证码
│   ├── config.go # 读配置，有一个全局 Config 变量
│   ├── invitecode.go # 生成邀请码，暂且使用 UUID
│   ├── jwt.go # jwt token 验证
│   ├── logger.go # 本地日志
│   ├── mysql.go # 全局 DB，单例初始化
│   └── redis.go # 全局 RDB，单例初始化
├── config.yml # 配置文件
├── go.mod
├── main.go
└── README.md
``

模块调用流程：

![](./assets/flow.png)

### 运行

```sh
go mod tidy
go run ./main.go
```

### Q&A

为什么有 model 中的有的文件前有个 2：

> `2_farm_miner.go` 开头带 2 的为中间表模型（不知道肿么想的🤷‍♂️🤷‍♀️，一眼能快速分辨中间表吧
> 原本想用下划线，但是 go 以下划线开始的文件似乎会被忽略

每个对象的 Newxxx() 是否多于？

> 
> ```go
> type FarmMinerDAO struct{}
> 
> func NewFarmMinerDAO() > *FarmMinerDAO {
> 	return &FarmMinerDAO{}
> }
> ```
> 
> 统一构造函数，方便在其他结构中初始化对象

为什么有的关键模块文件这么少：

> 因为没写完（doge

### TODO

- 表结构需要调整

- 中间表后续使用 gorm 的 tag `many2many`

- 登陆日志需要单独出来

- 用户请求存在问题，发起请求的 user_id 应该从上下文获取

- 规范回包，目前只做了简单的错误返回

- redis 持久化

- 单元测试，压力测试

- 连接 Miner

- ...
