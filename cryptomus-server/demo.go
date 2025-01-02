package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Cryptomus API 配置
const (
	// BaseURL Cryptomus API 基础 URL
	BaseURL = "https://api.cryptomus.com/v1"
	// APIKey API 密钥
	APIKey = "k3aKfqebrOxTEvOaDfwPyLvM9Ebln5rBGdSmTYrIIxmsiTpirFoSMBI3wpLubROVAGisCEgb5Fz9CgBN8Erv3g7pxiADu3PkiUukNgRexrRsB0ZZK5ARd7Z1qFiiUknw\n "
	// Merchant 替换为你的商户 ID
	Merchant = "b8ccc979-636f-4879-9c07-0889e8e1c53f"
)

// CreateInvoiceRequest 订单请求结构体
type CreateInvoiceRequest struct {
	Amount   string `json:"amount"`   // 支付金额
	Currency string `json:"currency"` // 币种 (如 USD, BTC)
	OrderID  string `json:"order_id"` // 订单号
}

// CreateInvoiceResponse 订单响应结构体
type CreateInvoiceResponse struct {
	State  int `json:"state"` // 响应状态码
	Result struct {
		UUID           string  `json:"uuid"`            // 支付 UUID
		OrderID        string  `json:"order_id"`        // 订单号
		Amount         string  `json:"amount"`          // 支付金额
		PaymentAmount  *string `json:"payment_amount"`  // 支付金额（加密货币）
		PayerAmount    *string `json:"payer_amount"`    // 付款金额
		Discount       string  `json:"discount"`        // 折扣
		Currency       string  `json:"currency"`        // 支付币种
		PaymentStatus  string  `json:"payment_status"`  // 支付状态
		URL            string  `json:"url"`             // 支付链接
		ExpiredAt      int64   `json:"expired_at"`      // 到期时间戳
		Status         string  `json:"status"`          // 订单状态
		IsFinal        bool    `json:"is_final"`        // 是否最终状态
		CreatedAt      string  `json:"created_at"`      // 创建时间
		UpdatedAt      string  `json:"updated_at"`      // 更新时间
		AdditionalData *string `json:"additional_data"` // 附加数据
	} `json:"result"`
}

// PaymentInfoRequest 定义支付信息查询请求的结构体
type PaymentInfoRequest struct {
	UUID    string `json:"uuid"`
	OrderID string `json:"order_id"`
}

// PaymentInfoResponse 定义支付信息查询响应的结构体
type PaymentInfoResponse struct {
	State  int `json:"state"`
	Result struct {
		UUID            string    `json:"uuid"`
		OrderID         string    `json:"order_id"`
		Amount          string    `json:"amount"`
		PaymentAmount   string    `json:"payment_amount"`
		PayerAmount     string    `json:"payer_amount"`
		DiscountPercent int       `json:"discount_percent"`
		Discount        string    `json:"discount"`
		PayerCurrency   string    `json:"payer_currency"`
		Currency        string    `json:"currency"`
		Comments        *string   `json:"comments"`
		MerchantAmount  string    `json:"merchant_amount"`
		Network         string    `json:"network"`
		Address         string    `json:"address"`
		From            *string   `json:"from"`
		TxId            *string   `json:"txid"`
		PaymentStatus   string    `json:"payment_status"`
		URL             string    `json:"url"`
		ExpiredAt       int64     `json:"expired_at"`
		Status          string    `json:"status"` //状态
		IsFinal         bool      `json:"is_final"`
		AdditionalData  *string   `json:"additional_data"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	} `json:"result"`
}

// PaymentStatusMap 定义一个支付状态的字典
var PaymentStatusMap = map[string]string{
	"paid":                 "付款是成功的，客户的付款准确地按要求。",
	"paid_over":            "付款成功，客户的支付超出了要求。",
	"wrong_amount":         "客户支付的费用低于要求",
	"process":              "处理中的付款",
	"confirm_check":        "我们已经看到了区块链中的交易，并正在等待所需的网络确认数。",
	"wrong_amount_waiting": "客户支付的支付少于要求，并可能需要额外付款",
	"check":                "等待交易出现在区块链上",
	"fail":                 "付款错误",
	"cancel":               "付款取消，客户没有付款",
	"system_fail":          "系统错误发生了",
	"refund_process":       "退款正在处理",
	"refund_fail":          "退款期间发生错误",
	"refund_paid":          "退款成功了",
	"locked":               "由于AML计划，资金被锁定",
}

// 生成签名函数
func generateSignature(data []byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 创建支付订单函数
func createInvoice(requestData CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	// 将请求数据转换为 JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("JSON 序列化失败: %v", err)
	}

	// 生成签名
	signature := generateSignature(jsonData, APIKey)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", BaseURL+"/payment", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建 HTTP 请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("merchant", Merchant)
	req.Header.Set("signature", signature)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送 HTTP 请求失败: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析响应 JSON
	var response CreateInvoiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	// 返回响应结果
	return &response, nil
}

// QueryPaymentStatus 查询支付状态
func QueryPaymentStatus(uuid, orderID string) (*PaymentInfoResponse, error) {

	// 创建请求体
	requestBody := PaymentInfoRequest{
		UUID:    uuid,
		OrderID: orderID,
	}

	// 将请求体转换为JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// 生成签名
	signature := generateSignature(jsonData, APIKey)

	// 创建HTTP请求
	req, err := http.NewRequest("POST", BaseURL+"/payment/info", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("merchant", Merchant)
	req.Header.Set("signature", signature)

	// 发送HTTP请求
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析响应体
	var paymentInfoResponse PaymentInfoResponse
	err = json.Unmarshal(body, &paymentInfoResponse)
	if err != nil {
		return nil, err
	}

	return &paymentInfoResponse, nil
}

// QueryPaymentStatusDemo 查询支付状态示例demo
func QueryPaymentStatusDemo() {
	uuid := "70b8db5c-b952-406d-af26-4e1c34c27f15"
	orderID := "65bbe87b4098c17a31cff3e71e515243"

	paymentInfo, err := QueryPaymentStatus(uuid, orderID)
	fmt.Println("paymentInfo:", paymentInfo)
	if err != nil {
		log.Fatalf("Error querying payment status: %v", err)
	}

	statusDescription, exists := PaymentStatusMap[paymentInfo.Result.PaymentStatus]
	if !exists {
		statusDescription = "未知状态"
	}

	fmt.Printf("Payment Status: %s (%s)\n", paymentInfo.Result.PaymentStatus, statusDescription)
}

func CreateInvoiceDemo() {
	// 示例请求参数
	invoiceRequest := CreateInvoiceRequest{
		Amount:   "15",  // 支付金额
		Currency: "USD", // 币种
		OrderID:  "1",   // 商户订单号
	}

	// 创建支付订单
	invoiceResponse, err := createInvoice(invoiceRequest)
	fmt.Println("invoiceResponse:", invoiceResponse)
	if err != nil {
		log.Fatalf("创建支付订单失败: %v", err)
	}

	// 输出支付订单详情
	fmt.Printf("支付订单创建成功！\n")
	fmt.Printf("支付链接: %s\n", invoiceResponse.Result.URL)
	fmt.Printf("订单号: %s\n", invoiceResponse.Result.OrderID)
	fmt.Printf("支付金额: %s %s\n", invoiceResponse.Result.Amount, invoiceResponse.Result.Currency)
	fmt.Printf("支付状态: %s\n", invoiceResponse.Result.PaymentStatus)
	fmt.Printf("创建时间: %s\n", invoiceResponse.Result.CreatedAt)
	fmt.Printf("到期时间戳: %d\n", invoiceResponse.Result.ExpiredAt)
}

// 主函数
func main() {
	// 创建订单
	CreateInvoiceDemo()
	// 查询订单状态
	QueryPaymentStatusDemo()
}
