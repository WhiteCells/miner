### Farm

#### POST

##### /farm

添加 Farm
- request
    ```json
    {
        "name": "user1_farm2",  // Farm 名称
        "time_zone": "t1",      // 时区
    }
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": {
            "id": 14,
            "name": "user1_farm2",
            "time_zone": "t1",
            "hash": ""
        }, // Farm 对象
        "msg": "create farm success" // 信息
    }
    ```

#### DELETE

##### /farm

删除 Farm
- request
    ```json
    {
        "farm_id": 1,  // Farm ID
    }
    ```
- response
    ```json
    {
        "code": 200,  // 状态码
        "data": null,
        "msg": "delete farm success" // 信息
    }
    ```
#### PUT

##### /farm

修改 Farm
- request
    ```json
    {
        "farm_id": 12,   // Farm ID 
        "update_info": {
            "name": "user1_farm2"
        } // 更新数据对象
    }
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": null,
        "msg": "update farm success" // 信息
    }
    ```

##### /farm/apply_fs

应用 Flightsheet

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
        "data": ,      // 
        "msg": "",     // 信息
    }
    ```

##### /farm/transfer

转移矿场
- request
    ```json
    {
        "farm_id": 1,       // Farm ID
        "to_user_id": 1,    // 转换到的用户 ID
    }
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": ,      // 
        "msg": "",     // 信息
    }
    ```

#### GET

##### /farm

获取 Farm
- params
    ```
    ?page_num=1   // 页号
    &page_size=2  // 页数
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": [
            {
                "id": 12,
                "name": "user1_farm2",
                "time_zone": "t1",
                "hash": ""
            },
            {
                "id": 13,
                "name": "user1_farm2",
                "time_zone": "t1",
                "hash": ""
            },
            {
                "id": 14,
                "name": "user1_farm2",
                "time_zone": "t1",
                "hash": ""
            }
        ], // Farm 对象数组
        "msg": "get farm success", // 信息
        "total": 3 // 总量
    }
    ```