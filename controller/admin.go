package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController() *AdminController {
	return &AdminController{
		adminService: service.NewAdminService(),
	}
}

// AdminGetUser 获取所有用户
func (c *AdminController) GetAllUser(ctx *gin.Context) {
	users, err := c.adminService.GetAllUser(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get all user failed", nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get all user success", users)
}

// AdminGetOperLog 获取所有用户操作日志
func (c *AdminController) GetUserOperLogs(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_numt", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_size", nil)
		return
	}
	query := map[string]interface{}{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	users, total, err := c.adminService.GetUserOperLogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get all user failed", nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get all user success", users, total)
}

// GetUserLoginLogs 获取用户登陆日志
func (c *AdminController) GetUserLoginLogs(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_numt", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_size", nil)
		return
	}
	query := map[string]interface{}{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	users, total, err := c.adminService.GetUserLoginLogs(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get user login logs failed", nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get user login logs success", users, total)
}

// GetUserPointsRecords 获取用户的积分记录
func (c *AdminController) GetUserPointsRecords(ctx *gin.Context) {
	pageNum, err := strconv.Atoi(ctx.Query("page_num"))
	if err != nil || pageNum <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_numt", nil)
		return
	}
	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil || pageSize <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid page_size", nil)
		return
	}
	query := map[string]interface{}{
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	users, total, err := c.adminService.GetUserPointsRecords(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get user points records failed", nil)
		return
	}

	rsp.DBQuerySuccess(ctx, http.StatusOK, "admin get user points records success", users, total)
}

// GetUserFarms 获取用户的所有矿场
func (c *AdminController) GetUserFarms(ctx *gin.Context) {
	farms, err := c.adminService.GetUserFarms(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get user farms failed", nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user farms success", farms)
}

// GetUserMiners 获取用户的所有矿机
func (c *AdminController) GetUserMiners(ctx *gin.Context) {
	farmID := ""
	miners, err := c.adminService.GetUserMiners(ctx, farmID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get user miners failed", nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user miners success", miners)
}

// SwitchRegister 用户注册开关
func (c *AdminController) SwitchRegister(ctx *gin.Context) {
	var req dto.AdminSwitchRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetSwitchRegister(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin switch register failed", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin switch register success", nil)
}

// SetGlobalFs 设置全局飞行表
func (c *AdminController) SetGlobalFs(ctx *gin.Context) {
	var req dto.AdminSetGlobalFsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetGlobalFs(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin set global flightsheet faild", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set global flightsheet success", nil)
}

// SetInviteReward 设置邀请积分奖励
func (c *AdminController) SetInviteReward(ctx *gin.Context) {
	var req dto.AdminSetInviteRewardReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetInviteReward(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin set invite reward faild", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set invite reward success", nil)
}

// SetRechargeReward 设置充值兑换积分比率
func (c *AdminController) SetRechargeRatio(ctx *gin.Context) {
	var req dto.AdminSetRechargeRewardReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetRechargeRatio(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin set recharge reward faild", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set recharge reward success", nil)
}

// SetUserStatus 设置用户状态
func (c *AdminController) SetUserStatus(ctx *gin.Context) {
	var req dto.AdminSetUserStatusReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetUserStatus(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin set user status faild", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set user status success", nil)
}

// SetMinerPoolCost 设置矿池费用
func (c *AdminController) SetMinePoolCost(ctx *gin.Context) {
	var req dto.AdminSetMinePoolCostReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetMinepoolCost(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin set miner poolCost faild", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set miner poolCost success", nil)
}

// SetMnemonic 设置助记词
func (c *AdminController) SetMnemonic(ctx *gin.Context) {
	var req dto.AdminSetMnemonicReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.SetMnemonic(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "set mnemonic", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "set mnemonict success", nil)
}

// GetMnemonic 获取活跃助记词
func (c *AdminController) GetMnemonic(ctx *gin.Context) {
	mnemonic, err := c.adminService.GetMnemonic(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "no active mnemonic", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get mnemonict success", mnemonic)
}

// GetAllMnemonic 获取所有助记词
func (c *AdminController) GetAllMnemonic(ctx *gin.Context) {
	mnemonics, err := c.adminService.GetAllMnemonic(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "no mnemonic", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get mnemonict success", mnemonics)
}

// AddBscApiKey 添加 apikey
func (c *AdminController) AddBscApiKey(ctx *gin.Context) {
	var req dto.AdminAddBscApiKeyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.AddBscApiKey(ctx, req.Apikey); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to add bsc apikey", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "add bsc apikey success", "")
}

// 获取 apikey（获取使用最少的）
func (c *AdminController) GetBscApiKey(ctx *gin.Context) {
	apikey, err := c.adminService.GetBscApiKey(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get bsc apikey", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get bsc apikey success", apikey)
}

// DelBscApiKey 删除 apikey
func (c *AdminController) DelBscApiKey(ctx *gin.Context) {
	var req dto.AdminDelBscApiKeyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.DelBscApiKey(ctx, req.Apikey); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to del bsc apikey", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "del bsc apikey success", "")
}

// coin
func (c *AdminController) AddCoin(ctx *gin.Context) {
	var req dto.AdminAddCoinReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.AddCoin(ctx, &req.Coin); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to add coin", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "add coin success", "")
}

func (c *AdminController) DelCoin(ctx *gin.Context) {
	var req dto.AdminDelCoinReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.DelCoin(ctx, req.Name); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to del coin", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "del coin success", "")
}

func (c *AdminController) GetCoinByName(ctx *gin.Context) {
	coinName := ctx.Query("name")
	pool, err := c.adminService.GetCoinByName(ctx, coinName)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get coin", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get coin success", pool)
}

func (c *AdminController) GetAllCoin(ctx *gin.Context) {
	coins, err := c.adminService.GetAllCoin(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get all coin", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get all coin success", coins)
}

// pool
func (c *AdminController) AddPool(ctx *gin.Context) {
	var req dto.AdminAddPoolReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.AddPool(ctx, &req.Pool); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to add pool", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "add pool success", "")
}

func (c *AdminController) DelPool(ctx *gin.Context) {
	var req dto.AdminDelPoolReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.DelPool(ctx, req.Name); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to del pool", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "del pool success", "")
}

func (c *AdminController) GetPoolByName(ctx *gin.Context) {
	poolName := ctx.Query("name")
	pool, err := c.adminService.GetPoolByName(ctx, poolName)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get pool", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get pool success", pool)
}

func (c *AdminController) GetAllPool(ctx *gin.Context) {
	pools, err := c.adminService.GetAllPool(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get all pool", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get all pool success", pools)
}

// soft
func (c *AdminController) AddSoft(ctx *gin.Context) {
	var req dto.AdminAddSoftReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.AddSoft(ctx, &req.Soft); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to add soft", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "add soft success", "")
}

func (c *AdminController) DelSoft(ctx *gin.Context) {
	var req dto.AdminDelSoftReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", err.Error())
		return
	}
	if err := c.adminService.DelSoft(ctx, req.Name); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to del soft", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "del soft success", "")
}

func (c *AdminController) GetSoftByName(ctx *gin.Context) {
	softName := ctx.Query("name")
	soft, err := c.adminService.GetSoftByName(ctx, softName)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get soft", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get soft success", soft)
}

func (c *AdminController) GetAllSoft(ctx *gin.Context) {
	softs, err := c.adminService.GetAllSoft(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "failed to get all soft", err.Error())
		return
	}
	rsp.Success(ctx, http.StatusOK, "get all soft success", softs)
}
