### Wallet

#### POST

##### /wallet

添加 Wallet
- request
    ```json
    {
        "name": "wallet1",   // Wallet Name
        "address": "0x1123", // Wallet 地址
        "coin_type": "c1"    // 代币类型
    }
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": {
            "id": 3,
            "name": "wallet1",
            "address": "0x1123",
            "coin_type": "c1"
        }, // Wallet 对象
        "msg": "create wallet succes" // 信息
    }
    ```

#### DELETE

##### /wallet

删除 Wallet
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

#### PUT

##### /wallet

修改 Wallet
- request
    ```json
    {
        "wallet_id": 1, // Wallet ID
        "update_info": {
            "name": "wallet11",
            "address": "0x321"
            // ...
        } // Wallet 更新对象
    }
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": null,
        "msg": "update wallet success" // 信息
    }
    ```

#### GET

##### /wallet

获取 Wallet
- params
    ```
    ?page_num=1,
    &page_size=10
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": [
            {
                "id": 1,
                "name": "wallet11",
                "address": "0x321",
                "coin_type": "c1"
            },
            {
                "id": 2,
                "name": "wallet1",
                "address": "0x1123",
                "coin_type": "c1"
            }
        ], // Wallet 对象数组
        "msg": "get user all wallet success", // 信息
        "total": 2 // 总数
    }
    ```

##### /wallet/:wallet_id

获取指定 Wallet 
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
