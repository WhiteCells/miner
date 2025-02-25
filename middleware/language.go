package middleware

import "github.com/gin-gonic/gin"

var languages = map[string]bool{
	"zh": true,
	"en": true,
}

func Language() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.GetHeader("Accept-Language")
		if _, ok := languages[lang]; !ok {
			lang = "zh"
		}
		ctx.Set("lang", lang)
		ctx.Next()
	}
}
