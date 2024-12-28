### MINERS HIVE CONFIG ################################\n
#URL where hive server is\n
str = """
HIVE_HOST_URL=\"http://127.0.0.1:9090/hiveos\"\n
API_HOST_URLs=\"http://127.0.0.1:9090/hiveos\"\n
RIG_ID=10101\n
RIG_PASSWD=\"1q2w3e4r\"\n
#Rig hostname\n
WORKER_NAME=\"15\"\n
#Id of the farm\n
FARM_ID=3335302\n
#Selected miners\n
MINER=custom\n
MINER2=\n
#Rig timezone\n
TIMEZONE=\"Europe/Kiev\"\n
#Watchdog\n
WD_ENABLED=1\n
WD_MINER=3\n
WD_REBOOT=5\n
WD_CHECK_GPU=0\n
WD_MAX_LA=900\n
WD_ASR=\n
WD_POWER_ENABLED=0\n
WD_POWER_MIN=\n
WD_POWER_MAX=\n
WD_POWER_ACTION=\n
WD_CHECK_CONN=0\n
WD_SHARE_TIME=\n
WD_MINHASHES='{}'\n
WD_MINHASHES_ALGO='{}'\n
WD_TYPE='miner'\n
#Hive Shell server host\n
HSSH_SRV=\"http://127.0.0.1:9090/hiveos\"\n
#Options\n
X_DISABLED=1\n
MINER_DELAY=1\n
DOH_ENABLED=0\n
SHELLINABOX_ENABLE=1\n
SSH_ENABLE=1\n
SSH_PASSWORD_ENABLE=1\n
"""

print(str)