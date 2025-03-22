package controller

import (
	"miner/common/dto"
	"miner/common/rsp"
	"miner/services"
	"miner/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
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

	clientIP := ctx.ClientIP()

	permissions, token, user, err := c.userService.Login(ctx, clientIP, &req)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ctx.Set("user_id", user.ID)

	rsp.LoginSuccess(ctx, http.StatusOK, "login success", user, token, permissions)
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
// func (c *UserController) UpdatePasswd(ctx *gin.Context) {
// 	var req dto.UpdatePasswdReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}
// 	if err := c.userService.UpdatePasswd(ctx, &req); err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "update passwd success", nil)
// }

// GetPointsBalance 获取积分余额
func (c *UserController) GetPointsBalance(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	balance, err := c.userService.GetUserPointsBalance(ctx, userID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get points balance success", balance)
}

func (c *UserController) GetUserAddress(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	address, err := c.userService.GetUserAddress(ctx, userID)
	if err != nil {
		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rsp.Success(ctx, http.StatusOK, "get user address success", address)
}

// GetCoins 获取币种信息
// func (c *UserController) GetCoins(ctx *gin.Context) {
// 	coins, err := c.userService.GetCoins(ctx)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "get coins success", coins)
// }

// GetPools 获取矿池信息
// func (c *UserController) GetPools(ctx *gin.Context) {
// 	coinName := ctx.Query("coin_name")
// 	pools, err := c.userService.GetPools(ctx, coinName)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "get pools success", pools)
// }

//// SetSoft 设置 Custon miner soft 信息
//func (c *UserController) SetSoft(ctx *gin.Context) {
//	var req dto.AddSoftReq
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
//		return
//	}
//	if err := c.userService.AddSoft(ctx, req.SoftName, &req.Soft); err != nil {
//		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), err.Error())
//		return
//	}
//	rsp.Success(ctx, http.StatusOK, "add soft success", nil)
//}
//
//// DelSoft 删除 Custon miner soft 信息
//func (c *UserController) DelSoft(ctx *gin.Context) {
//	var req dto.DelSoftReq
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
//		return
//	}
//	if err := c.userService.DelSoft(ctx, req.SoftName); err != nil {
//		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), err.Error())
//		return
//	}
//	rsp.Success(ctx, http.StatusOK, "add soft success", nil)
//}
//
//// UpdateSoft 修改 Custon miner soft 信息
//func (c *UserController) UpdateSoft(ctx *gin.Context) {
//	var req dto.UpdateSoftReq
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		rsp.Error(ctx, http.StatusBadRequest, err.Error(), nil)
//		return
//	}
//	if err := c.userService.AddSoft(ctx, req.SoftName, &req.Soft); err != nil {
//		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), err.Error())
//		return
//	}
//	rsp.Success(ctx, http.StatusOK, "update soft success", nil)
//}

//// GetSoft 获取 Custom miner soft 信息
//func (c *UserController) GetSoft(ctx *gin.Context) {
//	soft_name := ctx.Query("soft_name")
//	softs, err := c.userService.GetSoft(ctx, soft_name)
//	if err != nil {
//		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
//		return
//	}
//	rsp.Success(ctx, http.StatusOK, "get soft success", softs)
//}

func (c *UserController) GetRouters(ctx *gin.Context) {
	data, err := utils.UtilsGetRouters()
	if err != nil {
		rsp.Error(ctx, http.StatusForbidden, "get routers error", nil)
	}
	rsp.Success(ctx, http.StatusOK, "get routers success", data)
}

// 获取全局挖矿软件
// func (c *UserController) GetSoftAll(ctx *gin.Context) {
// 	coinName := ctx.Query("coin_name")
// 	softList, err := c.userService.GetSoftAll(ctx, coinName)
// 	if err != nil {
// 		rsp.Error(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	rsp.Success(ctx, http.StatusOK, "get soft list success", softList)
// }
