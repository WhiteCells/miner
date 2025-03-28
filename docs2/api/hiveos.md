### Hiveos

#### 客户端请求服务器

##### POST /hiveos/worker/api

params
```
id_rig=1231231
&method=stats
```

#### 服务器向客户端递交任务

##### POST /hiveos/task

body
```
{
    "farm_id": "4131235123423",
    "miner_id": "1329812313124",
    "rig_id": "89203292",
    "type": "cmd" # cmd/config
    "content": "miner start"
}
```

回包任务 ID

#### 获取任务结果

params
```
task_id=18392184715213
```