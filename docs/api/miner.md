### Miner

#### POST

##### /miner

添加 Miner
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

#### DELETE

##### /miner

删除 Miner
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

#### PUT

##### /miner

修改 Miner
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
##### /miner/apply_fs

应用 Flightsheet
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
#### GET

##### /miner

获取 Farm 下的 Miner
- params
    ```
    ?farm_id=1    // farm ID
    &page_num=1   // 页号
    &page_size=1  // 每页个数
    ```
- response
    ```json
    {
        "code": 200, // 状态码
        "data": [
            {
                "id": 1,
                "name": "user2_farm1_miner2",
                "gpu_info": "",
                "status": 1,
                "ip": "127.1.1.1",
                "ssh_port": 222,
                "ssh_user": "u",
                "ssh_password": "p"
            },
            {
                "id": 2,
                "name": "user2_farm1_miner2",
                "gpu_info": "",
                "status": 1,
                "ip": "127.1.1.1",
                "ssh_port": 222,
                "ssh_user": "u",
                "ssh_password": "p"
            }
        ], // Miner 对象数组
        "msg": "get user all miner success", // 信息
        "total": 3 // 总数
    }
    ```

##### /miner/:miner_id

获取指定 Miner
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
