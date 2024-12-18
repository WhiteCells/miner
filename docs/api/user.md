### User

#### POST

##### /user/register

注册
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
##### /user/login

登陆
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

##### /user/logout

登出
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
#### GET   

##### /user/oper_logs

获取用户日志
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
##### /user/balance

获取用户积分余额
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
##### /user/points_records

获取用户积分记录
- params
    ```
    ?page_num=1,                       // 页号
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
