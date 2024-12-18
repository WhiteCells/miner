### Admin

#### GET

##### /admin/all_users

获取所有用户信息
- params
    ```
    ?page_num=1   // 请求的页号
    &page_size=2  // 请求的一页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // User 对象
        "msg": "",     // 信息
        "total": 100,  // 总数
    }
    ```
    ```json
    {
        "code": 200, // 状态码
        "data": [
            {
                "id": 1,
                "name": "user1",
                "password": "$2a$10$py00YITltnsrubpOWKD8WeAdd76/o3fkRlciWWYjuYDt6OOI/xwRu",
                "secret": "PJF4YWVXV6PPGCL25SKLDCWT5364L6WKCI52IIZHWXOG63EJP5LKSHRM5KKZ4N3X",
                "email": "user2@user2.com",
                "role": "admin",
                "points": 0,
                "status": 1,
                "last_login_at": "2024-12-18T09:59:30+08:00",
                "last_login_ip": "127.0.0.1",
                "invite_code": "43e00220-aa49-4a0d-a512-5b717d320104",
                "invited_by": 0
            },
            {
                "id": 2,
                "name": "user3",
                "password": "$2a$10$Ph3AgeQiOA2QPql0TBieOOnSiI8HyDQnd6v4WgHZcUlshQ.b/Sz6O",
                "secret": "4UT5LTPV3LFFDFRFRX24PFK6QQC2AD54ZJ7HBRBLBDNEJKVPJAAYW3BG2PKEQCEL",
                "email": "user3@user2.com",
                "role": "user",
                "points": 0,
                "status": 0,
                "last_login_at": "0001-01-01T00:00:00Z", // 之前没有登陆
                "last_login_ip": "",
                "invite_code": "0b3989f0-f654-43ac-9dba-a269faf93e00",
                "invited_by": 0
            }
        ], // User 对象
        "msg": "admin get all user success", // 信息
        "total": 4 // 总数
    }
    ```
##### /admin/user_oper_logs

获取所有用户操作日志
- params
    ```
    ?page_num=1   // 请求的页号
    &page_size=10 // 请求的一页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // OperLog 对象
        "msg": "",     // 信息
        "total": 100,  // 总数
    }
    ```

##### /admin/user_login_logs

获取所有用户登陆日志
- params
    ```
    ?page_num=1   // 请求的页号
    &page_size=10 // 请求的一页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // OperLog 对象
        "msg": "",     // 信息
        "total": 100,  // 总数
    }
    ```

##### /admin/user_points_records

获取所有用户积分记录
- params
    ```
    ?page_num=1   // 请求的页号
    &page_size=10 // 请求的一页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // OperLog 对象
        "msg": "",     // 信息
        "total": 100,  // 总数
    }
    ```

#### POST   

#####  /admin/invite_reward

设置邀请积分
- params
    ```
    &page_num=1   // 请求的页号
    &page_size=10 // 请求的一页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // User 对象
        "msg": "",     // 信息
        "total": 100,  // 总数
    }
    ```

##### /admin/recharge_reward

设置充值返现积分
- params
    ```
    &page_num=1   // 请求的页号
    &page_size=10 // 请求的一页大小
    ```
- response
    ```json
    {
        "code": 200,   // 状态码
        "data": [{}],  // User 对象
        "msg": "",     // 信息
        "total": 100,  // 总数
    }
    ```
