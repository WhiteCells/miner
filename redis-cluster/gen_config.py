import os

template = """# redis 端口
port {port}

# 后台运行
daemonize yes

# 任何 IP 可以访问
bind 0.0.0.0

# 指定集群密码
masterauth m3i2n1e0r

# 指定集群密码
requirepass m3i2n1e0r

# cluster
cluster-enabled yes

# cluster config
cluster-config-file nodes.conf

# 超时时间，单位 ms
cluster-node-timeout 5000

# 开启持久化（AOF）  
appendonly yes

# 数据存储目录
dir {dir}{sep}store

# 日志
logfile {dir}{sep}log.txt
"""

base_dir = os.getcwd()
ports = [7000, 7001, 7002, 7003, 7004, 7005]


def generate_redis_configs(base_dir, ports):
    for port in ports:
        dir_path = os.path.join(base_dir, str(port))
        os.makedirs(dir_path, exist_ok=True)
        config_path = os.path.join(dir_path, "redis.conf")
        with open(config_path, "w", encoding="utf-8") as conf_file:
            conf_file.write(template.format(port=port, dir=dir_path, sep=os.sep))
        print(f"Generated config for port {port}: {config_path}")


if __name__ == "__main__":
    generate_redis_configs(base_dir, ports)
