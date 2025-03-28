# API 文档

基础信息
- Base URL: `https://127.0.0.1:9090`  
- 返回格式: `JSON`  
- 认证方式: `Bearer Token`  
- 错误响应

| 状态码 | 说明 |
|--------|------|
| `400` | 请求参数错误 |
| `401` | 未授权（Token 失效） |
| `500` | 服务器内部错误 |
---

## User

### login
- 接口: `POST /user/login`  
- 描述: 登陆

#### 请求参数

| 参数        | 位置 | 类型   | 必填 | 限制 | 说明                   |
| ----------- | ---- | ------ | ---- | ---- | ---------------------- |
| email       | body | string | 是   |      | 邮箱（必须是邮箱格式） |
| password    | body | string | 是   |      | 密码（6~32位）         |
| google_code | body | string | 是   |      | google 验证码          |

#### 请求示例
```http
POST /user/login HTTP/1.1
Host: 127.0.0.1:9090
Content-Type: application/json
{
	"email": "xxx@xxx.com",
	"password": "12345678",
	"google_code": "634327"
}
```

#### 响应参数

| 参数         | 类型          | 说明      |
| ------------ | ------------- | --------- |
| access_token | string        | jwt token |
| code         | int           | 状态      |
| data         | object        | 用户信息  |
| msg          | string        | 响应消息  |
| permissions  | null 或 array | 权限      |

#### 响应示例

```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJ1c2VyMSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQzMjEzMjIxLCJpYXQiOjE3NDMxMjY4MjF9.rNrIGQd0UhTVMnmbAzE6BaC1EffpMwQDp3Bqlj6ay4M",
    "code": 200,
    "data": {
        "id": 1,
        "name": "user1",
        "email": "xxx@xxx.com",
        "password": "$2a$10$GnrLy5oE3vt1XyTWyfFXwOCckBin9Gv/DYUdYJVTeX4ggz2Lxm5/m",
        "secret": "XKC5HJRUDZDVVWPTG6CMRGDEA4WIISDMSMUEPR2HUALU27P4ZBF4PEOW5B3Q7SDU",
        "address": "",
        "role": "user",
        "invite_points": 0,
        "recharge_points": 0,
        "last_balance": 0,
        "status": "1",
        "uid": "1904443575978954752",
        "last_check_at": "0001-01-01T00:00:00Z",
        "last_login_at": "2025-03-28T09:53:41.486573517+08:00",
        "last_login_ip": "172.16.20.232",
        "invite_code": "1904443575978954752",
        "invited_by": "",
        "key": ""
    },
    "msg": "login success",
    "permissions": null
}
```
---
### logout
- 接口: `POST /user/logout`
- 描述: 登出

#### 请求参数

无请求参数

#### 请求示例

```http
POST /user/logout HTTP/1.1
Host: 127.0.0.1:9090
Authorization: Bearer access_token
```

#### 响应参数

| 参数 | 类型   | 说明   |
| ---- | ------ | ------ |
| code | int    | 状态码 |
| data | object | 数据   |
| msg  | string | 信息   |

#### 响应示例

```json
{
    "code": 200,
    "data": null,
    "msg": "logout success"
}
```
---
### register
- 接口: `POST /user/register`
- 描述: 注册

#### 请求参数

| 参数        | 位置 | 类型   | 必填 | 限制 | 说明   |
| ----------- | ---- | ------ | ---- | ---- | ------ |
| username    | body | string | 是   |      | 用户名 |
| password    | body | string | 是   |      | 密码   |
| email       | body | string | 是   |      | 邮箱   |
| invite_code | body | string | 否   |      | 邀请码 |

#### 请求示例
```http
POST /user/register HTTP/1.1
Host: 127.0.0.1:9090
Content-Type: application/json
{
	"username": "user1",
	"password": "xxx1111",
}
```

#### 相应参数

| 参数 | 类型          | 说明              |
| ---- | ------------- | ----------------- |
| code | int           | 状态码            |
| data | google secret | google 验证码密钥 |
| msg  | string        | 信息              |

#### 响应示例

```json
{
    "code": 200,
    "data": "HH5464UXYF3CSGFJD4APTJMYNLCKYSA3PFP52A5JMFWZFU4HE5UCH42EIV2GNOD6",
    "msg": "register success"
}
```
---

### balance

- 接口：`GET /user/balance`
- 描述：获取用户余额，同时也更新用户余额

#### 请求参数

无请求参数

#### 请求示例

```http
POST /user/balance HTTP/1.1
Host: 127.0.0.1:9090
```

#### 响应参数

| 参数 | 类型   | 说明   |
| ---- | ------ | ------ |
| code | int    | 状态码 |
| data | int    | 余额   |
| msg  | string | 信息   |

#### 响应示例

```json
{
    "code": 200,
    "data": 0,
    "msg": "get points balance success"
}
```

---

### address

- 接口：`GET /user/address`
- 描述：



