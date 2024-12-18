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

	users, total, err := c.adminService.GetAllUser(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get all user failed", nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get all user success", users, total)
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

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get all user success", users, total)
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

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user login logs success", users, total)
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

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user points records success", users, total)
}

// GetUserFarms 获取用户的所有矿场
func (c *AdminController) GetUserFarms(ctx *gin.Context) {
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
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil || userID <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid user_id", nil)
		return
	}
	query := map[string]interface{}{
		"user_id":   userID,
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	farms, total, err := c.adminService.GetUserFarms(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get user farms failed", nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user farms success", farms, total)
}

// GetUserMiners 获取用户的所有矿机
func (c *AdminController) GetUserMiners(ctx *gin.Context) {
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
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil || userID <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid user_id", nil)
		return
	}
	farmID, err := strconv.Atoi(ctx.Query("farm_id"))
	if err != nil || farmID <= 0 {
		rsp.Error(ctx, http.StatusBadRequest, "invalid farm_id", nil)
		return
	}
	query := map[string]interface{}{
		"user_id":   userID,
		"farm_id":   farmID,
		"page_num":  pageNum,
		"page_size": pageSize,
	}

	miners, total, err := c.adminService.GetUserMiners(ctx, query)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin get user miners failed", nil)
		return
	}

	rsp.QuerySuccess(ctx, http.StatusOK, "admin get user miners success", miners, total)
}

// SwitchRegister 用户注册开关
func (c *AdminController) SwitchRegister(ctx *gin.Context) {
	var req dto.AdminSwitchRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SwitchRegister(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin switch register failed", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin switch register success", nil)
}

// SetGlobalFlightsheet 设置全局飞行表
func (c *AdminController) SetGlobalFlightsheet(ctx *gin.Context) {
	var req dto.AdminSetGlobalFlightsheetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetGlobalFlightsheet(ctx, &req); err != nil {
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

// SetRechargeReward 设置充值积分奖励
func (c *AdminController) SetRechargeReward(ctx *gin.Context) {
	var req dto.AdminSetRechargeRewardReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetRechargeReward(ctx, &req); err != nil {
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
func (c *AdminController) SetMinerPoolCost(ctx *gin.Context) {
	var req dto.AdminSetMinerPoolCostReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	if err := c.adminService.SetMinePoolCost(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, "admin set miner poolCost faild", nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "admin set miner poolCost success", nil)
}
