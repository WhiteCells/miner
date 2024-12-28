### Hiveos 对接 Demo

虚拟机网络桥接

hiveos 执行

```sh
# ip 地址换为本地网络回环地址
firstrun -f http://192.168.1.213:9090/hiveos
```

```sh
hello
```

常用命令
```sh
miner # 查看矿机运行情况

phoenix 挖矿软件 -gpus 12345678 Teamredminer 挖矿软件-d 0,1,2,3,4,5,6,7 调试里面填写 禁卡参数，不填哪个数字禁哪张卡

firstrun -f 后台没有矿机信息，后台添加一台矿机生成矿机 ID 和密码，矿机端输入后台添加 矿机的 ID 和密码与后台同步，矿机误删除也可以用过这个命令重新添加到后台

net-test网络测试查看网关 DNS 服务器地址通不通

amd-infoA 卡查看信息

nvidia-infoN 卡查看信息

date --s "2020xxxx xx:xx:xx"(修改成当前时间打入例如:date --s "20201204 11:06:59")同步 时间

hwclock -w把时间写入主板 BIOS

netconf-set --dhcp=no --test=1 --address=你的 IP/24 --gateway=你的网关 --dns=你的 dns 地址 修改固定 IP

sed -i 's/ pci=noaer / pci=noaer,nocrs /g' /etc/default/grub && update-grub昂达 D1800 主板前 两个卡槽不能识别运行完重启

Wname改矿机名

agent-screen— 显示配置单元客户端代理(您可以使用 退出 Ctrl + A ， D )

firstrun -f— 再次要求提供装备 ID 和密码，机器端输入 firstrun -f 等待加载完，然后直接输 入 OS 的 api 地址回车 ,然后按提示输入 id 和 password，id 和 password 都在你的机器的设定 里面。

mc- 像 Norton Commander 一样的文件管理器，但适用于 Linux selfupgrade-从控制台升级，就像点击网络上的按钮一样 sreboot-进行硬重启

sreboot shutdown— 硬关机

矿工

miner— 显示正在运行的矿工屏幕(您可以使用 退出 Ctrl + A ， D ) miner start, miner stop-很明显地启动或停止当前配置的矿工

miner log, miner config— 不解自明

系统日志

dmesg— 查看系统消息，主要是查看启动日志

tail -n 100 /var/log/syslog— 显示系统日志中的最后 100 行

网络

ifconfig— 显示网络接口

iwconfig— 显示无线适配器

Ctrl + C —停止任何正在运行的命令

切换矿工屏幕，与终端分离:

Ctrl + A ， D —从屏幕(矿工或代理)分离以使其正常工作

CTRL+A ，空间或 CTRL+A ， 1 ， 2 ， 3-屏幕之间进行切换，如果你有第二矿工运 行等

状态/诊断

agent-screen log— 显示 Hive 代理的各个部分的日志(可以尝试 log1 和 log2) hello— 问好 向服务器 :刷新 IP 地址，配置等。通常在启动时运行。 net-test—检查并诊断您的网络连接

timedatectl— 显示时间和日期同步设置 top -b -n 1—显示所有过程的清单

wd status— 显示哈希率看门狗状态和日志

AMD

amd-info— 显示 AMD 卡的当前频率

amdcovc— 显示 AMD 卡的当前频率

amdmeminfo— 显示扩展的 AMD 卡信息

wolfamdctrl -i 0 --show-voltage— 显示 AMD GPU 电压表 #0 的

英伟达

journalctl -p err | grep NVRM— 显示最近的 Nvidia GPU 错误(如果有)

nvidia-info— 显示扩展的 Nvidia 卡信息

nvidia-driver-update 430— 下载并安装 430 系列的最新驱动程序。*

nvidia-driver-update --nvs— 仅重新安装 nvidia-settings

nvidia-smi— 显示 Nvidia 卡信息

nvtool --clocks— 显示所有 Nvidia GPU 的核心/内存时钟

硬件

gpu-fans-find- 将 GPU 风扇从第一个 GPU 旋转到最后一个 GPU，从而更轻松地找到所需的 GPU sensors—显示主板和 CPU 的电压/温度读数

sreboot wakealarm 120— 关闭 PSU 并在 120 秒内启动

/hive/opt/opendev/watchdog-opendev power— 将电源命令发送给 OpenDev 看门狗

/hive/opt/opendev/watchdog-opendev reset— 将重置命令发送给 OpenDev 看门狗 /hive/opt/opendev/watchdog-opendev settings—显示 OpenDev 看门狗设置

升级/安装

disk-expand -s— 扩展 Linux 分区以填充剩余的驱动器空间 hpkg list miners— 列出所有已安装的矿工 hpkg remove miners- 卸载所有矿工 nvidia-driver-update --remove— 删除所有已下载的 Nvidia 驱动程序软件包(当前安装的软件 包除外) selfupgrade --force— 强制升级; 在自我升级显示 OS 是最新的但实际上不是最新的情况下， 它可以提供帮助

```