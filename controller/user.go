package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
	"miner/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserSerivce(),
	}
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	secret, err := c.userService.Register(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.Success(ctx, http.StatusOK, "register success", secret)
}

// Login 用户登陆
func (c *UserController) Login(ctx *gin.Context) {
	var req dto.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	token, user, err := c.userService.Login(ctx, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp.LoginSuccess(ctx, http.StatusOK, "login success", user, token)
}

// Logout 用户登出
func (c *UserController) Logout(ctx *gin.Context) {
	if err := c.userService.Logout(ctx); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "logout success", nil)
}

// UpdatePasswd 修改密码
func (c *UserController) UpdatePasswd(ctx *gin.Context) {
	var req dto.UpdatePasswdReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err := c.userService.UpdatePasswd(ctx, &req); err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "update passwd success", nil)
}

// GetPointsBalance 获取积分余额
func (c *UserController) GetPointsBalance(ctx *gin.Context) {
	balance, err := c.userService.GetPointsBalance(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get points balance success", balance)
}

func (c *UserController) GetUserAddress(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	address, err := c.userService.GetUserAddress(ctx, userID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get user address success", address)
}

// GetCoins 获取币种信息
func (c *UserController) GetCoins(ctx *gin.Context) {
	coins, err := c.userService.GetCoins(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get coins success", coins)
}

// GetPools 获取矿池信息
func (c *UserController) GetPools(ctx *gin.Context) {
	pools, err := c.userService.GetPools(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get pools success", pools)
}

// GetSofts 获取软件信息
func (c *UserController) GetSofts(ctx *gin.Context) {
	softs, err := c.userService.GetSofts(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get softs success", softs)
}

// GenerateCaptcha
func (c *UserController) GenerateCaptcha(ctx *gin.Context) {
	id, b64s, err := utils.GenerateCaptcha(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	captcha := &dto.GenerateCaptchaRsp{
		CaptchaID: id,
		Base64:    b64s,
	}
	rsp.Success(ctx, http.StatusOK, "generate captcha success", captcha)
}

// VerifyCaptcha
func (c *UserController) VerifyCaptcha(ctx *gin.Context) {
	var req dto.VerifyCaptchaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if !utils.VerifyCaptcha(ctx, req.CaptchaID, req.Value) {
		rsp.Error(ctx, http.StatusForbidden, "captcha error", nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "verify captcha success", nil)
}
