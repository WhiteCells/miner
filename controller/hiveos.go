package controller

import (
	"miner/service"

	"github.com/gin-gonic/gin"
)

type HiveOsController struct {
	hiveOsService *service.HiveOsService
}

func NewHiveOsController() *HiveOsController {
	return &HiveOsController{
		hiveOsService: service.NewHiveOsService(),
	}
}

// 交互
// func (c *HiveOsController) Poll(ctx *gin.Context) {
// 	id_rig := ctx.Query("id_rig")
// 	method := ctx.Query("method")

// 	fmt.Println(id_rig, method)

// 	id_rig_int, err := strconv.Atoi(id_rig)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
// 		return
// 	}

// 	// 区分请求
// 	switch method {
// 	case "hello":
// 		var req dto.HiveosReq
// 		if err := ctx.ShouldBindJSON(&req); err != nil {
// 			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
// 			return
// 		}
// 		//
// 		jsonInd, err := json.MarshalIndent(req, "", "  ")
// 		if err != nil {
// 			fmt.Println("")
// 			return
// 		}
// 		// fmt.Printf("<< message >>: %+v\n", req)
// 		fmt.Printf("%s\n", jsonInd)
// 		//
// 		fmt.Printf("<< stats >> %+v\n", req)
// 		ctx.JSON(http.StatusOK, &dto.ServerRsp{
// 			ID:      id_rig_int,
// 			Jsonrpc: "2.0",
// 			Result: struct {
// 				ID        int    `json:"id"`
// 				Config    string `json:"config"`
// 				Wallet    string `json:"wallet"`
// 				Autofan   string `json:"autofan"`
// 				Justwrite int    `json:"justwrite"`
// 				Command   string `json:"command"`
// 				Exec      string `json:"exec"`
// 				Confseq   int    `json:"confseq"`
// 			}{
// 				ID:        99999111,
// 				Config:    "HIVE_HOST_URL=\"http://172.16.0.176:9090/hiveos\"\nAPI_HOST_URLs=\"http://172.16.0.176:9090/hiveos\"\nRIG_ID=10101\nRIG_PASSWD=\"1q2w3e4r\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
// 				Wallet:    "### Wallet \n# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
// 				Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
// 				Justwrite: 1,
// 				// Command:   "",
// 				// Exec:      "",
// 				Confseq: 1,
// 			},
// 		})
// 	case "stats":
// 		var req dto.HiveosReq
// 		if err := ctx.ShouldBindJSON(&req); err != nil {
// 			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
// 			return
// 		}
// 		//
// 		jsonInd, err := json.MarshalIndent(req, "", "  ")
// 		if err != nil {
// 			fmt.Println("")
// 			return
// 		}
// 		// fmt.Printf("<< message >>: %+v\n", req)
// 		fmt.Printf("%s\n", jsonInd)
// 		//
// 		fmt.Printf("<< stats >> %+v\n", req)
// 		ctx.JSON(http.StatusOK, &dto.ServerRsp{
// 			ID:      id_rig_int,
// 			Jsonrpc: "2.0",
// 			Result: struct {
// 				ID        int    `json:"id"`
// 				Config    string `json:"config"`
// 				Wallet    string `json:"wallet"`
// 				Autofan   string `json:"autofan"`
// 				Justwrite int    `json:"justwrite"`
// 				Command   string `json:"command"`
// 				Exec      string `json:"exec"`
// 				Confseq   int    `json:"confseq"`
// 			}{
// 				ID:        99999,
// 				Config:    "HIVE_HOST_URL=\"http://172.16.0.176:9090/hiveos\"\nAPI_HOST_URLs=\"http://172.16.0.176:9090/hiveos\"\nRIG_ID=10101\nRIG_PASSWD=\"1q2w3e4r\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
// 				Wallet:    "### Wallet \n# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
// 				Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
// 				Justwrite: 1,
// 				Command:   "exec",
// 				Exec:      "ls -la",
// 				Confseq:   1,
// 			},
// 		})
// 	case "message":
// 		var req dto.HiveosResReq
// 		if err := ctx.ShouldBindJSON(&req); err != nil {
// 			rsp.Error(ctx, http.StatusBadRequest, err.Error(), "")
// 			return
// 		}
// 		//
// 		jsonInd, err := json.MarshalIndent(req, "", "  ")
// 		if err != nil {
// 			fmt.Println("")
// 			return
// 		}
// 		// fmt.Printf("<< message >>: %+v\n", req)
// 		fmt.Printf("%s\n", jsonInd)
// 		//
// 		ctx.JSON(http.StatusOK, &dto.ServerRsp{
// 			ID:      id_rig_int,
// 			Jsonrpc: "2.0",
// 			Result: struct {
// 				ID        int    `json:"id"`
// 				Config    string `json:"config"`
// 				Wallet    string `json:"wallet"`
// 				Autofan   string `json:"autofan"`
// 				Justwrite int    `json:"justwrite"`
// 				Command   string `json:"command"`
// 				Exec      string `json:"exec"`
// 				Confseq   int    `json:"confseq"`
// 			}{
// 				ID:        99999,
// 				Config:    "HIVE_HOST_URL=\"http://172.16.0.176:9090/hiveos\"\nAPI_HOST_URLs=\"http://172.16.0.176.4:9090/hiveos\"\nRIG_ID=10101\nRIG_PASSWD=\"1q2w3e4r\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
// 				Wallet:    "# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
// 				Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
// 				Justwrite: 1,
// 				// Command:   "exec",
// 				// Exec:      "ls -la",
// 				Confseq: 1,
// 			},
// 		})
// 	}
// }

// 轮询
func (c *HiveOsController) Poll(ctx *gin.Context) {
	id_rig := ctx.Query("id_rig")
	method := ctx.Query("method")
	c.hiveOsService.Poll(ctx, id_rig, method)
}

// 发送命令
func (c *HiveOsController) SendCmd(ctx *gin.Context) {
	// 构建
	// 构建 Task
}

// 获取命令结果
func (c *HiveOsController) GetCmdRes(ctx *gin.Context) {

}

// 设置配置
func (c *HiveOsController) SetConfig(ctx *gin.Context) {

}

// 获取设置配置结果
func (c *HiveOsController) GetConfigRes(ctx *gin.Context) {

}

// 获取状态
func (c *HiveOsController) GetStats(ctx *gin.Context) {

}
