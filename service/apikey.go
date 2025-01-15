package service

import (
	"encoding/json"
	"fmt"
	"miner/common/dto"
	"miner/dao/redis"
	"miner/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BscApiKeyService struct {
	bscApiKeyRDB *redis.BscApiKeyRDB
}

func NewBscApiKeyService() *BscApiKeyService {
	return &BscApiKeyService{
		bscApiKeyRDB: redis.NewBscApiKeyRDB(),
	}
}

// 获取地址余额
// 后端发送余额查询请求 ID
// 前端始终记录最新 ID，在用户刷新积分后使用 ID 查询结果
func (s *BscApiKeyService) GetTokenBalance(ctx *gin.Context, address string) (string, error) {
	apikey, err := s.bscApiKeyRDB.ZRangeWithScore(ctx)
	if err != nil {
		return "", err
	}
	// 增加 apikey 使用次数
	if err := s.bscApiKeyRDB.ZIncrBy(ctx, apikey, 1); err != nil {
		return "", err
	}

	// 测试使用阻塞
	// 使用数据库记录
	return s.requestBscApi(ctx, address, apikey)
}

// 调用接口
func (s *BscApiKeyService) requestBscApi(ctx *gin.Context, address string, apikey string) (string, error) {
	url := fmt.Sprintf(utils.Config.Bsc.Api, address, apikey)
	go func() {
		utils.Logger.Info("generate " + url + "failed")
	}()

	rsp, err := http.Get(url)

	time.Sleep(300000000)

	// 减少 apikey 的使用次数
	s.bscApiKeyRDB.ZIncrBy(ctx, apikey, -1)

	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("BSC API returned non-200 status code: %d", rsp.StatusCode)
	}

	var body dto.BscApiRspBody
	if err := json.NewDecoder(rsp.Body).Decode(&body); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	return body.Result, nil
}
