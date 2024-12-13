### API

### User

#### POST   /user/register                注册
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
#### POST   /user/login                   登陆
- request
    ```json
    {
        "username": "",       // 用户名
        "password": "",       // 密码
        "captcha_code": "",   // 图形验证码
        "google_code": "",    // google 验证码
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
#### POST   /user/logout                  登出
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
#### GET   /user/oper_logs                  获取用户日志
- params
    ```
    ?page_size=1                      // 每页数量
    &page_num=1                       // 页号
    &action=DELETE                    // 可选类型
    &start_time=2020-12-01T23:59:59Z  // 开始时间
    &end_time=2024-12-13T23:59:59Z    // 结束时间
    ```
- response
    ```json
    {
        "code": 200,         // 状态码
        "data": [{}],        // 操作日志对象数组
        "msg": "",           // 信息
        "total": 111,        // 总数 
    }
    ```
#### GET   /user/balance                  获取用户积分余额
- params
    ```
    ?page_size=1                      // 每页数量
    &page_num=1                       // 页号
    &start_time=2020-12-01T23:59:59Z  // 开始时间
    &end_time=2024-12-13T23:59:59Z    // 结束时间
    ```
- response
    ```json
    {
        "code": 200,         // 状态码
        "data": 1111,        // 用户积分余额
        "msg": "",           // 信息
    }
    ```
#### GET   /user/points_records              获取用户积分记录
- params
    ```
    &page_num=1,                       // 页号
    &page_size=30,                     // 每页数量 
    &start_time=2024-12-01T23:59:59Z   // 开始时间
    &end_time=2024-12-02T23:59:59Z"    // 结束时间
    ```
- response
    ```json
    {
        "code": 200,         // 状态码
        "data": [{}],        // 积分记录对象数组
        "msg": "",           // 信息
        "total": 111,        // 总数
    }
    ```

### Miner

#### POST   /miner                        添加 Miner
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
#### DELETE /miner                        删除 Miner
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
#### PUT    /miner                        修改 Miner
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
#### GET    /miner                        获取 Farm 下的 Miner
- params
    ```
    &farm_id=1    // farm ID
    &page_num=1   // 页号
    &page_size=1  // 每页个数
    &start_time=  // 时间区间
    &end_time=    // 时间区间
    &order=       // 排序
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // 返回 Miner 对象数组
        "msg": "",     // 信息
    }
    ```
#### GET    /miner/:miner_id              获取指定 Miner
- params
    ```

    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": {},    // 返回 Miner 对象
        "msg": "",     // 信息
    }
    ```
#### PUT   /miner/apply-fs               应用 Flightsheet
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

### Farm

#### POST   /farm                         添加 Farm
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
#### DELETE /farm                         删除 Farm
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
#### PUT    /farm                         修改 Farm
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
#### GET    /farm                         获取 Farm
- request
    ```

    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // Farm 对象数组
        "msg": "",     // 信息
    }
    ```
#### GET    /farm/:farm_id              获取指定 Farm
- request
    params
    ```

    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": {},    // Farm 对象
        "msg": "",     // 信息
    }
    ```
#### PUT   /farm/apply-fs                应用 Flightsheet
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

### Flightsheet

#### POST   /flightsheet                  添加 Flightsheet
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
#### DELETE /flightsheet                  删除 Flightsheet
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
#### PUT    /flightsheet                  修改 Flightsheet
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
#### GET    /flightsheet                  获取 Flightsheet
- params
    ```
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // flightsheet 对象数组
        "msg": "",     // 信息
    }
    ```
#### GET    /flightsheet/:fs_id           获取指定 Flightsheet
- params
    ```
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": {},    // flightsheet 对象
        "msg": "",     // 信息
    }
    ```
#### PUT   /flightsheet/apply-wallet     应用 Wallet
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

### Wallet

#### POST   /wallet                       添加 Wallet
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
#### DELETE /wallet                       删除 Wallet
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
#### PUT    /wallet                       修改 Wallet
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
#### GET    /wallet                       获取 Wallet
- params
    ```
    ```
- response
    ```json
    {
        "code": 200,  // 状态码
        "data": [{}], // Wallet 对象数组
        "msg": "",    // 信息
    }
    ```
#### GET    /wallet/:fs_id                获取指定 Wallet 
- params
    ```
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": {},  // Wallet 对象
        "msg": "",   // 信息
    }
    ```

### Admin