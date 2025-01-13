修改配置文件中的数据存储及日志目录

使用 redis-cli 创建集群

```sh
redis-cli --cluster create \
127.0.0.1:7000 \
127.0.0.1:7001 \
127.0.0.1:7002 \
127.0.0.1:7003 \
127.0.0.1:7004 \
127.0.0.1:7005 \
--cluster-replicas 1 \
-a m3i2n1e0r # 密码
```

## win上启动集群
启动redis
```shell
start cmd /k D:\tool\Redis-x64-5.0.14\redis-server.exe 7000/redis_win.conf
start cmd /k D:\tool\Redis-x64-5.0.14\redis-server.exe 7001/redis_win.conf
start cmd /k D:\tool\Redis-x64-5.0.14\redis-server.exe 7002/redis_win.conf
start cmd /k D:\tool\Redis-x64-5.0.14\redis-server.exe 7003/redis_win.conf
start cmd /k D:\tool\Redis-x64-5.0.14\redis-server.exe 7004/redis_win.conf
start cmd /k D:\tool\Redis-x64-5.0.14\redis-server.exe 7005/redis_win.conf
```

启动集群
```shell
D:\tool\Redis-x64-5.0.14\redis-cli.exe --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1 -a m3i2n1e0r
```
