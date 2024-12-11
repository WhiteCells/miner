## Miner

### 项目结构

```sh
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
```

模块调用流程：

![](./assets/flow.png)

### 运行

```sh
go mod tidy
go run ./main.go
```

### 接口

- **User**
  - POST   /user/register                注册
    - request
        ```json
        {
            "username": "",   // 3～32 长度
            "password": "",   // 6～32 长度
            "email": "",      // email 格式
            "invite_code": "" // 可选
        }
        ```
    - response
        ```json
        {
            "code": 200,
        }
        ```
  - POST   /user/login                   登陆
    - request
        ```json
        {
            "username": "",   // 用户名
            "password": "",   // 密码
            <!-- "captcha_code": "", -->
            "google_code": "", // google 验证码
        }
        ```
    - response
        ```json
        {
            "asscess_token": "", // token
            "msg": "",           // 信息
            "user": {},          // User 对象
        }
        ```
  - POST   /user/logout                  登出
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200,         // 状态码
            "msg": "",           // 信息
        }
        ```
- **Miner**
  - POST   /miner                        添加 Miner
    - request
        ```json
        {
            "farm_id": 1,       // farm ID
            "name": "",         // miner 名称
            "model": "",        // model
            "ip": "",           // ip
            "ssh_port": "",     // ssh port
            "ssh_user": "",     // ssh user
            "ssh_passwrod": "", // ssh password
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // 返回 Miner 对象
            "msg": "",     // 信息
        }
        ```
  - DELETE /miner                        删除 Miner
    - request
        ```json
        {
            "farm_id": 1,  // Farm ID
            "miner_id": 1, // Miner ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```
  - PUT    /miner                        修改 Miner
    - request
        ```json
        {
            "miner_id": 1,    // Miner ID
            "update_info": {} // 更新 Miner 数据对象
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // 返回 Miner 对象
            "msg": "",     // 信息
        }
        ```
  - GET    /miner                        获取 Farm 下的 Miner
    - request
        ```json
        {
            "farm_id": 1, // 矿场
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": [{}],  // 返回 Miner 对象数组
            "msg": "",     // 信息
        }
        ```
  - GET    /miner/:miner_id              获取指定 Miner
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // 返回 Miner 对象
            "msg": "",     // 信息
        }
        ```
  - POST   /miner/apply-fs               应用 Flightsheet
    - request
        ```json
        {
            "fs_id": 1,     // flightsheet ID
            "miner_id": 1,  // miner ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```

- **Farm**
  - POST   /farm                         添加 Farm
    - request
        ```json
        {
            "name": "",      // Farm 名称
            "time_zone": "", // 时区
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // Farm 对象
            "msg": "",     // 信息
        }
        ```
  - DELETE /farm                         删除 Farm
    - request
        ```json
        {
            "farm_id": 1,  // Farm ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```
  - PUT    /farm                         修改 Farm
    - request
        ```json
        {
            "farm_id": 1,     // Farm ID 
            "update_info": {} // 更新数据对象
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```
  - GET    /farm                         获取 Farm
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": [{}],  // Farm 对象数组
            "msg": "",     // 信息
        }
        ```
  - GET    /farm/:farm_id              获取指定 Farm
    - request
        ```json
        {
            "farm_id": 1, // Farm ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // Farm 对象
            "msg": "",     // 信息
        }
        ```
  - POST   /farm/apply-fs                应用 Flightsheet
    - request
        ```json
        {
            "farm_id": 1,  // Farm ID
            "fs_id": 1,    // Flightsheet ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```

- **Flightsheet**
  - POST   /flightsheet                  添加 Flightsheet
    - request
        ```json
        {
            "name": "",      // Flightsheet Name
            "coin_type": "", // 货币类型
            "wallet_id": "", // 应用的 Wallet ID
            "mine_pool": "", // 矿池
            "mine_soft": "", // 挖矿软件
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // flightsheet 对象
            "msg": "",     // 信息
        }
        ```
  - DELETE /flightsheet                  删除 Flightsheet
    - request
        ```json
        {
            "fs_id": 1,   // Flightsheet ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```
  - PUT    /flightsheet                  修改 Flightsheet
    - request
        ```json
        {
            "fs_id": 1,       // Flightsheet ID
            "update_info": {} // Flightsheet 更新信息对象
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```
  - GET    /flightsheet                  获取 Flightsheet
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": [{}],  // flightsheet 对象数组
            "msg": "",     // 信息
        }
        ```
  - GET    /flightsheet/:fs_id           获取指定 Flightsheet
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // flightsheet 对象
            "msg": "",     // 信息
        }
        ```
  - POST   /flightsheet/apply-wallet     应用 Wallet
    - request
        ```json
        {
            "fs_id": 1,       // Flightsheet ID
            "wallet_id": 1,   // Wallet ID
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "msg": "",     // 信息
        }
        ```

- **Wallet**
  - POST   /wallet                       添加 Wallet
    - request
        ```json
        {
            "name": "",      // Wallet Name
            "address": "",   // Wallet 地址
            "coin_type": "", // 代币类型
        }
        ```
    - response
        ```json
        {
            "code": 200,   // 状态码
            "data": {},    // Wallet 对象
            "msg": "",     // 信息
        }
        ```
  - DELETE /wallet                       删除 Wallet
    - request
        ```json
        {
            "wallet_id": 1, // Wallet ID
        }
        ```
    - response
        ```json
        {
            "code": 200, // 状态码
            "msg": "",   // 信息
        }
        ```
  - PUT    /wallet                       修改 Wallet
    - request
        ```json
        {
            "wallet_id": 1,   // Wallet ID
            "update_info": {} // Wallet 更新对象
        }
        ```
    - response
        ```json
        {
            "code": 200, // 状态码
            "msg": "",   // 信息
        }
        ```
  - GET    /wallet                       获取 Wallet
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200,  // 状态码
            "data": [{}], // Wallet 对象数组
            "msg": "",    // 信息
        }
        ```
  - GET    /wallet/:fs_id                获取指定 Wallet 
    - request
        ```json
        {
        }
        ```
    - response
        ```json
        {
            "code": 200, // 状态码
            "data": {},  // Wallet 对象
            "msg": "",   // 信息
        }
        ```

### TODO

- 表结构需要调整

- 中间表后续使用 gorm 的 tag `many2many`

- 登陆日志需要单独出来

- 用户请求存在问题，发起请求的 user_id 应该从上下文获取

- 规范回包，目前只做了简单的错误返回

- redis 持久化

- 单元测试，压力测试

- 连接 Miner

- 数据关联在 dao 进行事务操作

- 传入参数没有检查

- 创建系统表，用控制系统配置，如积分奖励等

- ...
