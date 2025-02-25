package utils

var languages = map[string]map[string]string{
	"zh": {
		"LOGIN_SUCCESS":     "登录成功",
		"EMAIL_PASS_ERROR":  "邮箱或密码错误",
		"EMAIL_EXIST":       "邮箱已存在",
		"USER_NOT_EXIST":    "用户不存在",
		"OPERATE_SUCCESS":   "操作成功",
		"OPERATE_FAILED":    "操作失败",
		"CREATE_SUCCESS":    "创建成功",
		"CREATE_FAILED":     "创建失败",
		"DELETE_SUCCESS":    "删除成功",
		"DELETE_FAILED":     "删除失败",
		"UPDATE_SUCCESS":    "更新成功",
		"UPDATE_FAILED":     "更新失败",
		"GET_SUCCESS":       "获取成功",
		"GET_FAILED":        "获取失败",
		"PERMISSION_DENIED": "权限不足",
		"CONTEXT_ERROR":     "上下文错误",
		"TRANSFER_FAILED":   "转移失败",
	},
	"en": {
		"OPERATE_SUCCESS": "operate success",
		"OPERATE_FAILED":  "operate failed",
	},
}

func GetLanguageMsg(lang string, key string) string {
	if languages[lang] == nil {
		lang = "zh"
	}
	return languages[lang][key]
}
