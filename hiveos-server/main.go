package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MinerRequest struct {
	Method string `json:"method"`
	Params struct {
		V      int    `json:"v"`
		RigID  string `json:"rig_id"`
		Passwd string `json:"passwd"`
		meta   struct {
			FsID   int `json:"fs_id"`
			Custom struct {
				Coin string `json:"coin"`
			}
		}
		Temp       []int     `json:"temp"`
		Fan        []int     `json:"fan"`
		Power      []int     `json:"power"`
		Df         string    `json:"df"`
		Mem        []int     `json:"mem"`
		Cputemp    []int     `json:"cputemp"`
		Cpuavg     []float32 `json:"cpuavg"`
		Miner      string    `json:"miner"`
		TotalKhs   int       `json:"total_khs"`
		MinerStats struct {
			Status     string    `json:"status"`
			Khs        string    `json:"khs"`
			Hs         []float32 `json:"hs"`
			HsUnits    string    `json:"hs_units"`
			Temp       []int     `json:"temp"`
			Fan        []int     `json:"fan"`
			Uptime     int       `json:"uptime"`
			Ver        string    `json:"ver"`
			Ar         []int     `json:"ar"`
			Algo       string    `json:"algo"`
			BusNumbers []int     `json:"bus_numbers"`
		} `json:"miner_stats"`
	} `json:"params"`
}

type MessageRequest struct {
	Method  string `json:"method"`
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Params  struct {
		RigID   string `json:"rig_id"`
		Passwd  string `json:"passwd"`
		Type    string `json:"type"`
		Data    string `json:"data"`
		ID      string `json:"id"`
		Payload string `json:"payload"`
	}
}

type MinerResponse struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		ID        int    `json:"id"`
		Config    string `json:"config"`
		Wallet    string `json:"wallet"`
		Autofan   string `json:"autofan"`
		Justwrite int    `json:"justwrite"`
		Command   string `json:"command"`
		Exec      string `json:"exec"`
		Confseq   int    `json:"confseq"`
	} `json:"result"`
}

func main() {
	r := gin.Default()

	g := r.Group("/hiveos")
	{
		g.POST("/worker/api", func(ctx *gin.Context) {
			rig_id := ctx.Query("id_rig")
			method := ctx.Query("method")
			fmt.Println(rig_id, method)
			rig_id_int, err := strconv.Atoi(rig_id)
			if err != nil {
				return
			}

			switch method {
			case "hello":
				var req MinerRequest
				if err := ctx.ShouldBindJSON(&req); err != nil {
					return
				}

				fmt.Printf("[hello] Received request: %+v\n", req)

				if rig_id != "10101" || req.Params.Passwd != "1q2w3e4r" {
					return
				}

				ctx.JSON(http.StatusOK, &MinerResponse{
					ID:      rig_id_int,
					Jsonrpc: "2.0",
					Result: struct {
						ID        int    `json:"id"`
						Config    string `json:"config"`
						Wallet    string `json:"wallet"`
						Autofan   string `json:"autofan"`
						Justwrite int    `json:"justwrite"`
						Command   string `json:"command"`
						Exec      string `json:"exec"`
						Confseq   int    `json:"confseq"`
					}{
						ID:        1010111,
						Config:    "HIVE_HOST_URL=\"http://192.168.35.4:9090/hiveos\"\nAPI_HOST_URLs=\"http://192.168.35.4:9090/hiveos\"\nRIG_ID=10101\nRIG_PASSWD=\"1q2w3e4r\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
						Wallet:    "### FLIGHT SHEET \"k3ok.com-h9-post12\" ###\n\n# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
						Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
						Justwrite: 1,
						Command:   "",
						Exec:      "",
						Confseq:   1,
					},
				})
			case "stats":
				var req MinerRequest
				if err := ctx.ShouldBindJSON(&req); err != nil {
					return
				}

				fmt.Printf("[stats] Received request: %+v\n", req)

				if rig_id != "10101" || req.Params.Passwd != "1q2w3e4r" {
					return
				}

				// ctx.JSON(http.StatusOK, gin.H{
				// 	"id":      nil,
				// 	"jsonrpc": "2.0",
				// 	"result": map[string]interface{}{
				// 		"command": "exec",
				// 		"exec":    "miner start",
				// 		"id":      123456,
				// 	},
				// })

				ctx.JSON(http.StatusOK, &MinerResponse{
					ID:      rig_id_int,
					Jsonrpc: "2.0",
					Result: struct {
						ID        int    `json:"id"`
						Config    string `json:"config"`
						Wallet    string `json:"wallet"`
						Autofan   string `json:"autofan"`
						Justwrite int    `json:"justwrite"`
						Command   string `json:"command"`
						Exec      string `json:"exec"`
						Confseq   int    `json:"confseq"`
					}{
						ID:        101,
						Config:    "HIVE_HOST_URL=\"http://192.168.35.4:9090/hiveos\"\nAPI_HOST_URLs=\"http://192.168.35.4:9090/hiveos\"\nRIG_ID=10101\nRIG_PASSWD=\"1q2w3e4r\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
						Wallet:    "### Wallet \n# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
						Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
						Justwrite: 1,
						Command:   "exec",
						Exec:      "ls",
						Confseq:   1,
					},
				})
			case "message":
				var req MessageRequest
				if err := ctx.ShouldBindJSON(&req); err != nil {
					return
				}

				fmt.Printf("[message] Received request: %+v\n", req)

				if rig_id != "10101" || req.Params.Passwd != "1q2w3e4r" {
					return
				}

				ctx.JSON(http.StatusOK, &MinerResponse{
					ID:      rig_id_int,
					Jsonrpc: "2.0",
					Result: struct {
						ID        int    `json:"id"`
						Config    string `json:"config"`
						Wallet    string `json:"wallet"`
						Autofan   string `json:"autofan"`
						Justwrite int    `json:"justwrite"`
						Command   string `json:"command"`
						Exec      string `json:"exec"`
						Confseq   int    `json:"confseq"`
					}{
						ID:        101,
						Config:    "HIVE_HOST_URL=\"http://192.168.35.4:9090/hiveos\"\nAPI_HOST_URLs=\"http://192.168.35.4:9090/hiveos\"\nRIG_ID=10101\nRIG_PASSWD=\"1q2w3e4r\"\nWORKER_NAME=\"15\"\nFARM_ID=3335302\nMINER=custom\nMINER2=\nTIMEZONE=\"Europe/Kiev\"\nWD_ENABLED=1\nWD_MINER=3\nWD_REBOOT=5\nWD_CHECK_GPU=0\nWD_MAX_LA=900\nWD_ASR=\nWD_POWER_ENABLED=0\nWD_POWER_MIN=\nWD_POWER_MAX=\nWD_POWER_ACTION=\nWD_CHECK_CONN=0\nWD_SHARE_TIME=\nWD_MINHASHES='{}'\nWD_MINHASHES_ALGO='{}'\nWD_TYPE='miner'\nHSSH_SRV=\"http://192.168.35.15:9090/hiveos\"\nX_DISABLED=1\nMINER_DELAY=1\nDOH_ENABLED=0\nSHELLINABOX_ENABLE=1\nSSH_ENABLE=1\nSSH_PASSWORD_ENABLE=1\n",
						Wallet:    "# Miner custom\nCUSTOM_MINER=\"k3ok.com-spacemesh-s\"\nCUSTOM_INSTALL_URL=\"https://gitee.com/k3os/spacemesh/releases/download/v4.0.5/k3ok.com-spacemesh-s-v4.0.5.tar.gz\"\nCUSTOM_ALGO=\"\"\nCUSTOM_TEMPLATE=\"15\"\nCUSTOM_URL=\"http://hiveos.vip/\"\nCUSTOM_PASS=\"\"\nCUSTOM_USER_CONFIG='path:\n- /mnt/\nminerName: 15\napiKey: smh00000-0c79-5659-7b8f-565a95961ecf\nextraParams:\n  deleteLoadFail: false\n  device: \"\"\n  disableInitPost: false\n  disablePlot: true\n  disablePoST: false\n  flags: fullmem\n  maxFileSize: 32\n  nonces: 128\n  numUnits: 15\n  plotInstance: 1\n  postAffinity: 0\n  postAffinityStep: 1\n  postCpuIds: \"\"\n  postInstance: 0\n  postThread: 0\n  randomxAffinity: -1\n  randomxAffinityStep: 1\n  randomxThread: 0\n  removeInitFailed: false\n  reservedSize: 1\n  skipUninitialized: false\n  remoteK2Pow: true\nlog:\n  lv: info\n  path: ./log/\n  name: miner.log\nurl:\n  info: \"\"\n  submit: \"\"\n  line: \"\"\n  ws: \"\"\n  proxy: \"http://172.16.10.77:9090\"\nproxy:\n  url: \"\"\n  username: \"\"\n  password: \"\"\nhttp:\n  enable: false\n  host: \"\"\n  port: 0\nscanPath: false\nscanMinute: 60\ndebug: \"\"'\nCUSTOM_TLS=\"\"\n\nMETA='{\"fs_id\":20216083,\"custom\":{\"coin\":\"smh\"}}'\n",
						Autofan:   "ENABLED=\nTARGET_TEMP=\nTARGET_MEM_TEMP=\nMIN_FAN=\nMAX_FAN=\nCRITICAL_TEMP=\nCRITICAL_TEMP_ACTION=\"\"\nNO_AMD=\nREBOOT_ON_ERROR=\nSMART_MODE=\nCUSTOM_MODE=\"\"\nCUSTOM_TARGET_TEMP=\"\"\nCUSTOM_TARGET_MEM_TEMP=\"\"\nCUSTOM_MIN_FAN=\"\"\nCUSTOM_MAX_FAN=\"\"\nCUSTOM_CRITICAL_TEMP=\"\"\n",
						Justwrite: 1,
						// Command:   "exec",
						// Exec:      "miner start",
						Confseq: 1,
					},
				})
				// ctx.JSON(http.StatusOK, gin.H{
				// 	"jsonrpc": "2.0",
				// 	"id":      12345,
				// 	"result":  nil,
				// })
			}
		})
	}

	r.Run(":9090")
}
