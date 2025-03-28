### Flightsheet

#### POST   

##### /flightsheet

添加 Flightsheet
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

#### DELETE

##### /flightsheet

删除 Flightsheet
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

#### PUT

##### /flightsheet

修改 Flightsheet
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

##### /flightsheet/apply_wallet

应用 Wallet
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

#### GET

##### /flightsheet

获取 Flightsheet
- params
    ```
    ?page_num=1   // 页号
    &page_size=2  // 页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // flightsheet 对象数组
        "msg": "",     // 信息
        "total": "",   // 总数
    }
    ```

##### /flightsheet/:fs_id

获取指定 Flightsheet
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


