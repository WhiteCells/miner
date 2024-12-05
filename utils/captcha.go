package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaService struct{}

var store = base64Captcha.DefaultMemStore

type GenerateCaptchaResponse struct {
	CaptchaID   string `json:"captcha_id"`
	ImageBase64 string `json:"image_base64"`
}

type VerifyCaptchaRequest struct {
	CaptchaID string `json:"captcha_id"`
	Value     string `json:"value"`
}

func (c *CaptchaService) GenerateCaptcha(ctx *gin.Context) {
	// 配置验证码参数
	driver := base64Captcha.NewDriverDigit(
		80,  // 高度
		240, // 宽度
		6,   // 验证码长度
		0.7, // 干扰强度
		50,  // 干扰数量
	)

	// 创建验证码
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "生成验证码失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, GenerateCaptchaResponse{
		CaptchaID:   id,
		ImageBase64: b64s,
	})
}

func (c *CaptchaService) VerifyCaptcha(ctx *gin.Context) {
	var req VerifyCaptchaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数",
		})
		return
	}

	if !store.Verify(req.CaptchaID, req.Value, true) {
		ctx.JSON(http.StatusOK, gin.H{
			"valid":   false,
			"message": "验证码错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"message": "验证码正确",
	})
}
