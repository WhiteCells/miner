package controller

import (
	"miner/common/dto"
	"miner/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserContorller() *UserController {
	return &UserController{
		userService: service.NewUserSerivce(),
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	var req dto.LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, userID, err := c.userService.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(ctx)
	session.Set("user_id", userID)

	// 需要在 service 中完成，因为登陆失败后，无法设置 user_id
	// 导致无法记录登陆失败的记录
	// ctx.Set("user_id", userID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
		"user_id": userID,
	})
}

func (c *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.userService.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"data":    req.Username,
	})
}
