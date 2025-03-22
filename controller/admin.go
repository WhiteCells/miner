package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController() *AdminController {
	return &AdminController{
		adminService: services.NewAdminService(),
	}
}

// 获取所有用户
func (c *AdminController) GetUsers(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}
	users, total, err := c.adminService.GetUsers(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get all user success", users, total)
}

// 获取所有用户操作日志
func (c *AdminController) GetUserOperlogs(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	users, total, err := c.adminService.GetUserOperlogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get all user success", users, total)
}

// 获取用户登陆日志
func (c *AdminController) GetUserLoginlogs(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	users, total, err := c.adminService.GetUserLoginlogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get user login logs success", users, total)
}

// 获取用户的积分记录
func (c *AdminController) GetUserPointslogs(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	users, total, err := c.adminService.GetUserPointslogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get user points records success", users, total)
}

// 获取用户的所有矿场
func (c *AdminController) GetUserFarms(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	userID := ctx.GetInt("user_id")
	farms, total, err := c.adminService.GetUserFarms(ctx, userID, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get user farms success", farms, total)
}

// 获取用户的所有矿机
func (c *AdminController) GetUserMiners(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid parmas", nil)
		return
	}
	query := map[string]any{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	miners, total, err := c.adminService.GetUserMiners(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get user miners success", miners, total)
}

// 用户注册开关
func (c *AdminController) SetSwitchRegister(ctx *gin.Context) {
	var req dto.AdminSwitchRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}

	if err := c.adminService.SetSwitchRegister(ctx, req.Status); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin switch register success", nil)
}

func (c *AdminController) GetSwitchRegister(ctx *gin.Context) {
	switchRegister, err := c.adminService.GetSwitchRegister(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get switch register success", switchRegister)
}

// SetGlobalFs 设置全局飞行表
// func (c *AdminController) SetGlobalFs(ctx *gin.Context) {
// 	var req dto.CreateFsReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
// 		return
// 	}

// 	if err := c.adminService.CreateGlobalFs(ctx, &req); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "admin set global flightsheet faild", nil)
// 		return
// 	}

// 	rsp.Success(ctx, http.StatusOK, "admin set global flightsheet success", nil)
// }

// 获取邀请积分奖励
func (c *AdminController) GetInviteReward(ctx *gin.Context) {
	reward, err := c.adminService.GetInviteReward(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get invite reward success", reward)
}

// 设置邀请积分奖励
func (c *AdminController) SetInviteReward(ctx *gin.Context) {
	var req dto.AdminSetInviteRewardReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetInviteReward(ctx, req.Reward); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set invite reward success", nil)
}

// 获取邀请积分奖励
func (c *AdminController) GetRechargeRatio(ctx *gin.Context) {
	ratio, err := c.adminService.GetRechargeRatio(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get recharge ratiosuccess", ratio)
}

// 设置充值兑换积分比率
func (c *AdminController) SetRechargeRatio(ctx *gin.Context) {
	var req dto.AdminSetRechargeRewardReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetRechargeRatio(ctx, req.Ratio); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set recharge reward success", nil)
}

// 获取用户状态
func (c *AdminController) GetUserStatus(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	status, err := c.adminService.GetUserStatus(ctx, userID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user status success", status)
}

// 设置用户状态
func (c *AdminController) SetUserStatus(ctx *gin.Context) {
	var req dto.AdminSetUserStatusReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}
	if err := c.adminService.SetUserStatus(ctx, req.UserID, req.Status); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "admin set user status success", nil)
}

// GetMinePoolCost 获取矿池费用
// func (c *AdminController) GetMinePoolCost(ctx *gin.Context) {
// 	cost, err := c.adminService.GetMinePoolCost(ctx)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "admin get miner poolCost failed", nil)
// 		return
// 	}

// 	rsp.QuerySuccess(ctx, http.StatusOK, "admin get miner poolCost success", cost)
// }

// SetMinerPoolCost 设置矿池费用
// func (c *AdminController) SetMinePoolCost(ctx *gin.Context) {
// 	var req dto.AdminSetMinePoolCostReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
// 		return
// 	}

// 	if err := c.adminService.SetMinepoolCost(ctx, &req); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "admin set miner poolCost faild", nil)
// 		return
// 	}

// 	rsp.Success(ctx, http.StatusOK, "admin set miner poolCost success", nil)
// }

// 设置助记词
func (c *AdminController) SetMnemonic(ctx *gin.Context) {
	var req dto.AdminSetMnemonicReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.SetMnemonic(ctx, req.Mnemonic); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "set mnemonic", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "set mnemonics success", nil)
}

// 获取活跃助记词
func (c *AdminController) GetMnemonic(ctx *gin.Context) {
	mnemonic, err := c.adminService.GetMnemonic(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "no active mnemonic", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get mnemonics success", mnemonic)
}

// 获取所有助记词
func (c *AdminController) GetAllMnemonic(ctx *gin.Context) {
	mnemonics, err := c.adminService.GetAllMnemonic(ctx)
	if err != nil {
		if err.Error() == "redis: nil" {
			rsp.Success(ctx, http.StatusOK, "get mnemonics stats success", nil)
			return
		}
		rsp.Error(ctx, http.StatusInternalServerError, "no mnemonic", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get mnemonics success", mnemonics)
}

// 添加 apikey
func (c *AdminController) AddBscApiKey(ctx *gin.Context) {
	var req dto.AdminAddBscApiKeyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.AddBscApiKey(ctx, req.Apikey); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "add bsc apikey success", "")
}

// 获取 apikey（获取使用最少的）
func (c *AdminController) GetBscApiKey(ctx *gin.Context) {
	apikey, err := c.adminService.GetBscApiKey(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get bsc apikey success", apikey)
}

// 获取所有 apikey
func (c *AdminController) GetAllBscApiKey(ctx *gin.Context) {
	apiKeys, err := c.adminService.GetAllBscApiKey(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get bsc apikey success", apiKeys)
}

// 删除 apikey
func (c *AdminController) DelBscApiKey(ctx *gin.Context) {
	var req dto.AdminDelBscApiKeyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}
	if err := c.adminService.DelBscApiKey(ctx, req.Apikey); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "del bsc apikey success", nil)
}

// 添加 coin

// 删除 coin

// 获取 coin

// 获取所有 coin

// 添加 pool
// func (c *AdminController) AddPool(ctx *gin.Context) {
// 	var req dto.AdminAddPoolReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
// 		return
// 	}
// 	if err := c.adminService.AddPool(ctx, &req.Pool); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "failed to add pool", err.Error())
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "add pool success", "")
// }

// 删除 pool
// func (c *AdminController) DelPool(ctx *gin.Context) {
// 	var req dto.AdminDelPoolReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
// 		return
// 	}
// 	if err := c.adminService.DelPool(ctx, req.CoinName, req.PoolName); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "failed to del pool", err.Error())
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "del pool success", "")
// }

// 获取 pool
// func (c *AdminController) GetPool(ctx *gin.Context) {
// 	coinName := ctx.Query("coin_name")
// 	poolName := ctx.Query("pool_name")
// 	pool, err := c.adminService.GetPool(ctx, coinName, poolName)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "failed to get pool", err.Error())
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "get pool success", pool)
// }

// 获取所有 pool
// func (c *AdminController) GetAllPool(ctx *gin.Context) {
// 	coinName := ctx.Query("coinName")
// 	if coinName == "" {
// 		pools, err := c.adminService.GetAllPool(ctx)
// 		if err != nil {
// 			rsp.Error(ctx, http.StatusInternalServerError, "failed to get all pool", err.Error())
// 			return
// 		}
// 		rsp.Success(ctx, http.StatusOK, "get all pool success", pools)
// 	} else {
// 		pools, err := c.adminService.GetAllPoolByCoin(ctx, coinName)
// 		if err != nil {
// 			rsp.Error(ctx, http.StatusInternalServerError, "failed to get all pool", err.Error())
// 			return
// 		}
// 		rsp.Success(ctx, http.StatusOK, "get all pool success", pools)
// 	}
// }

// 设置免费 gpu 数量
func (m *AdminController) SetFreeGpuNum(ctx *gin.Context) {
	var req dto.AdminSetFreeGpuNumReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}
	if err := m.adminService.SetFreeGpuNum(ctx, req.GpuNum); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "set free gpu num success", nil)
}

// 设置免费 gpu 数量
func (m *AdminController) GetFreeGpuNum(ctx *gin.Context) {
	num, err := m.adminService.GetFreeGpuNum(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.QuerySuccess(ctx, http.StatusOK, "get free gpu num success", num)
}

// 添加 soft
// func (c *AdminController) AddSoft(ctx *gin.Context) {
// 	var req dto.AdminAddSoftReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
// 		return
// 	}
// 	if err := c.adminService.AddSoft(ctx, req.Soft.Coin, &req.Soft); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "failed to add soft", err.Error())
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "add pool success", "")
// }

// 删除 soft
// func (c *AdminController) DelSoft(ctx *gin.Context) {
// 	var req dto.AdminDelSoftReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
// 		return
// 	}
// 	if err := c.adminService.DelSoft(ctx, req.CoinName, req.Name); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "failed to del soft", err.Error())
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "del pool success", "")
// }

// 获取 soft
// func (c *AdminController) GetSoft(ctx *gin.Context) {
// 	coinName := ctx.Query("coinName")
// 	poolName := ctx.Query("poolName")
// 	pool, err := c.adminService.GetPool(ctx, coinName, poolName)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, "failed to get soft", err.Error())
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "get soft success", pool)
// }

// 获取所有 soft
// func (c *AdminController) GetAllSoft(ctx *gin.Context) {
// 	coinName := ctx.Query("coinName")
// 	if coinName == "" {
// 		pools, err := c.adminService.GetAllSoft(ctx)
// 		if err != nil {
// 			rsp.Error(ctx, http.StatusInternalServerError, "failed to get all soft", err.Error())
// 			return
// 		}
// 		rsp.Success(ctx, http.StatusOK, "get all soft success", pools)
// 	} else {
// 		pools, err := c.adminService.GetAllSoftByCoin(ctx, coinName)
// 		if err != nil {
// 			rsp.Error(ctx, http.StatusInternalServerError, "failed to get all soft", err.Error())
// 			return
// 		}
// 		rsp.Success(ctx, http.StatusOK, "get all soft success", pools)
// 	}
// }
