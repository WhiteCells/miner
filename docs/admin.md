### 获取所有用户信息 get /all_user

```py
import requests

url = "127.0.0.1:9090/admin/all_users"

headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3OTI5OTQ2NyIsInVzZXJfbmFtZSI6IjFAMS5jb20iLCJpc3MiOiJtaW5lciIsImV4cCI6MTc0MTM1NDk0NiwiaWF0IjoxNzQxMjY4NTQ2fQ.1Trlb4rUWyBte8wXmzsve-31FEqP0eQPvTmW8NaEKgA'
}

response = requests.request("GET", url, headers=headers)

print(response.text)
```

### 获取所有用户操作日志 get /user_oper_logs

```py
import requests

url = "127.0.0.1:9090/admin/user_oper_logs?page_num=1&page_size=10"

headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3OTI5OTQ2NyIsInVzZXJfbmFtZSI6IjFAMS5jb20iLCJpc3MiOiJtaW5lciIsImV4cCI6MTc0MTM1NDk0NiwiaWF0IjoxNzQxMjY4NTQ2fQ.1Trlb4rUWyBte8wXmzsve-31FEqP0eQPvTmW8NaEKgA'
}

response = requests.request("GET", url, headers=headers)

print(response.text)
```

### 获取所有用户登录日志 get /user_login_logs

```py
import requests

url = "127.0.0.1:9090/admin/user_login_logs?page_num=1&page_size=10"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3OTI5OTQ2NyIsInVzZXJfbmFtZSI6IjFAMS5jb20iLCJpc3MiOiJtaW5lciIsImV4cCI6MTc0MTM1NDk0NiwiaWF0IjoxNzQxMjY4NTQ2fQ.1Trlb4rUWyBte8wXmzsve-31FEqP0eQPvTmW8NaEKgA'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)

```

### 获取所有用户积分记录 get /user_points_records

```py
import requests

url = "127.0.0.1:9090/admin/user_points_records?page_num=1&page_size=10"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3OTI5OTQ2NyIsInVzZXJfbmFtZSI6IjFAMS5jb20iLCJpc3MiOiJtaW5lciIsImV4cCI6MTc0MTM1NDk0NiwiaWF0IjoxNzQxMjY4NTQ2fQ.1Trlb4rUWyBte8wXmzsve-31FEqP0eQPvTmW8NaEKgA'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 获取所有用户矿场 get /user_farms

```py
import requests

url = "127.0.0.1:9090/admin/user_farms"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3OTI5OTQ2NyIsInVzZXJfbmFtZSI6IjFAMS5jb20iLCJpc3MiOiJtaW5lciIsImV4cCI6MTc0MTM1NDk0NiwiaWF0IjoxNzQxMjY4NTQ2fQ.1Trlb4rUWyBte8wXmzsve-31FEqP0eQPvTmW8NaEKgA'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 获取指定矿场的所有矿机 get /user_miners

```py
import requests

url = "127.0.0.1:9090/admin/user_miners?farm_id=1896209052120125440"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjc3OTI5OTQ2NyIsInVzZXJfbmFtZSI6IjFAMS5jb20iLCJpc3MiOiJtaW5lciIsImV4cCI6MTc0MTM1NDk0NiwiaWF0IjoxNzQxMjY4NTQ2fQ.1Trlb4rUWyBte8wXmzsve-31FEqP0eQPvTmW8NaEKgA'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 设置注册开关 post /switch_register

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/switch_register"

payload = json.dumps({
  "status": "0"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 获取注册开关 get /switch_register

```py
import requests

url = "http://127.0.0.1:9090/admin/switch_register"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### post /global_fs

```py

```

### 获取邀请奖励 get /invite_reward

```py
import requests

url = "http://127.0.0.1:9090/admin/invite_reward"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 设置邀请奖励 post /invite_reward

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/invite_reward"

payload = json.dumps({
  "reward": "1"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 获取充值比例 get /recharge_ratio

```py
import requests

url = "http://127.0.0.1:9090/admin/recharge_ratio"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 设置充值比例 post /recharge_ratio

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/recharge_ratio"

payload = json.dumps({
  "ratio": 1.1
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 获取用户状态 get /user_status

```py
import requests

url = "http://127.0.0.1:9090/admin/user_status?user_id=1668598236"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 设置用户状态 post /user_status

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/user_status"

payload = json.dumps({
  "user_id": "0",
  "status": "0"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 设置助记词 post /mnemonic

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/mnemonic"

payload = json.dumps({
  "mnemonic": "bamboo obscure glove ancient furnace control isolate basket mutual gaze monster debris"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 获取助记词 get /mnemonic

```py
import requests

url = "http://127.0.0.1:9090/admin/mnemonic"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 获取所有助记词 get /all_mnemonic

```py
import requests

url = "http://127.0.0.1:9090/admin/all_mnemonic"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 设置 bsc apikey post /bsc_apikey

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/bsc_apikey"

payload = json.dumps({
  "apikey": "xxxxx"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 获取当前使用的 bsc apikey get /bsc_apikey

```py
import requests

url = "http://127.0.0.1:9090/admin/bsc_apikey"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 获取所有 bsc apikey get /all_bsc_apikey

```py
import requests

url = "http://127.0.0.1:9090/admin/all_bsc_apikey"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 删除 bsc apikey delete /bsc_apikey

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/all_bsc_apikey"

payload = json.dumps({
  "apikey": "xxxxx"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("DELETE", url, headers=headers, data=payload)

print(response.text)
```

### 设置 coin post /coin

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/coin"

payload = json.dumps({
  "coin": {
    "name": "etc"
  }
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 删除 coin delete /coin

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/coin"

payload = json.dumps({
  "coin_name": "etc"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("DELETE", url, headers=headers, data=payload)

print(response.text)
```

### put /coin

```py

```

### get /coin

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/coin"

payload = json.dumps({
  "coin_name": "etc"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 获取所有 coin get /all_coin

```py
import requests

url = "http://127.0.0.1:9090/admin/all_coin"

payload = ""
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 设置 pool post /pool

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/pool"

payload = json.dumps({
  "coin_name": "etc",
  "pool": {
    "name": "pool1",
    "urls": [
      {
        "name": "ssl-xxx",
        "host": "https://xxx.xxx"
      }
    ]
  }
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
```

### 删除 pool delete /pool

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/pool"

payload = json.dumps({
  "coin_name": "etc",
  "pool_name": "pool1"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("DELETE", url, headers=headers, data=payload)

print(response.text)
```

### put /pool

```py

```

### 获取 pool get /pool

```py
import requests
import json

url = "http://127.0.0.1:9090/admin/pool"

payload = json.dumps({
  "coin_name": "etc",
  "pool_name": "pool1"
})
headers = {
  'Content-Type': 'application/json',
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```

### 获取所有 pool get /all_pool

```py
import requests

url = "http://127.0.0.1:9090/admin/all_pool"

payload = {}
headers = {
  'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTYzMTcxNTkyIiwidXNlcl9uYW1lIjoiMUAxLmNvbSIsImlzcyI6Im1pbmVyIiwiZXhwIjoxNzQxMzk3NDY1LCJpYXQiOjE3NDEzMTEwNjV9.qRcJaAjHshsMJG_9U0aGTjWy4tRx0MjFFrCZh5oNzGU'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
```
