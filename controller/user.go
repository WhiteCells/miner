package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/service"
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

// GetPointsBalance 获取积分余额
func (c *UserController) GetPointsBalance(ctx *gin.Context) {
	balance, err := c.userService.GetPointsBalance(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get points balance success", balance)
}

// AuditAmount 查账
func (c *UserController) AuditAmount(ctx *gin.Context) {
	resultChan, errorChan := c.userService.AuditAmount(ctx)

	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case result := <-resultChan:
			rsp.Success(ctx, http.StatusOK, "success aduit", result)
		case err := <-errorChan:
			rsp.Error(ctx, http.StatusInternalServerError, "failed aduit", err.Error())
		case <-ctx.Done(): // 超时或取消
			rsp.Error(ctx, http.StatusGatewayTimeout, "request timeout", nil)
		}
	}()
	<-done
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

func (c *UserController) GetCoins(ctx *gin.Context) {
	coins, err := c.userService.GetCoins(ctx)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get coins success", coins)
}
