package wechat

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	PrepayID string `json:"prepay_id"` // 假设 JSON 响应中的字段是小写，且使用下划线命名
}

// WeChatPay 结构体包含微信支付 API 所需的基本信息
type WeChatPay struct {
	MchID               string
	AppID               string
	APIKey              string
	CertPath            string
	PrivateKey          *rsa.PrivateKey
	CertificateSerialNo string
}

// OrderRequest 结构体表示一个订单请求的详细信息
type OrderRequest struct {
	AppID       string `json:"appid"`
	MchID       string `json:"mchid"`
	Description string `json:"description"`
	OutTradeNo  string `json:"out_trade_no"`
	NotifyURL   string `json:"notify_url"`
	Amount      Amount `json:"amount"`
	Payer       Payer  `json:"payer"`
}

// Amount 结构体表示订单金额信息
type Amount struct {
	Total    float64 `json:"total"`
	Currency string  `json:"currency"`
}

// Payer 结构体表示支付方信息
type Payer struct {
	OpenID string `json:"openid"`
}

type OrderResponse struct {
	PrepayID string `json:"prepay_id"`
	PaySign  string `json:"paySign"`
}

// generateSignature 生成待签名消息的 RSA 签名
func (wc *WeChatPay) generateSignature(message []byte) (string, error) {
	hash := sha256.New()
	hash.Write(message)
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, wc.PrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// buildMessage 构建待签名的消息字符串
func (wc *WeChatPay) buildMessage(method, canonicalUrl string, timestamp int64, nonceStr, body string) string {
	return fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, canonicalUrl, timestamp, nonceStr, body)
}

// CreateOrder 创建订单并发送请求到微信支付 API
func (wc *WeChatPay) CreateOrder(orderRequest OrderRequest) (OrderResponse, error) {
	requestURL := "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"

	orderResponse := OrderResponse{}

	// 将订单请求转换为 JSON 格式
	orderJSON, err := json.Marshal(orderRequest)
	if err != nil {
		return orderResponse, err
	}

	// 生成随机字符串和时间戳
	nonceStr := "random_string"
	timestamp := time.Now().Unix()
	canonicalUrl := "/v3/pay/transactions/jsapi"
	// 构建待签名消息
	message := wc.buildMessage("POST", canonicalUrl, timestamp, nonceStr, string(orderJSON))
	// 生成签名
	signature, err := wc.generateSignature([]byte(message))
	if err != nil {
		return orderResponse, err
	}

	// 构建 Authorization 请求头
	authorization := fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",timestamp="%d",serial_no="%s",signature="%s"`,
		wc.MchID, nonceStr, timestamp, wc.CertificateSerialNo, signature)

	// 创建 HTTP 客户端并发送请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", requestURL, bytes.NewReader(orderJSON))
	if err != nil {
		return orderResponse, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Authorization", authorization)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return orderResponse, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // 使用 io.ReadAll 替换 ioutil.ReadAll
	if err != nil {
		return orderResponse, err
	}

	// 现在 bodyBytes 包含了响应体的字节序列
	var jsonResponse Response
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON response: %v", err)
	}

	// 打印解析后的 PrepayID
	fmt.Printf("PrepayID: %s\n", jsonResponse.PrepayID)

	fmt.Println("Response:", string(body))
	orderResponse.PrepayID = jsonResponse.PrepayID
	orderResponse.PaySign = wc.BuildSignature(jsonResponse.PrepayID)
	return orderResponse, nil
}

// 构建签名
func (wc *WeChatPay) BuildSignature(PrepayID string) string {
	appid := wc.AppID
	timestamp := time.Now().Unix()
	nonceStr := "random_string"
	packages := "prepay_id=" + PrepayID
	// 构建待签名消息
	message := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", appid, nonceStr, timestamp, packages, packages)
	// 生成签名
	s, err := wc.generateSignature([]byte(message))
	if err != nil {
		log.Fatalf("Error generating signature: %v", err)
	}

	return s
}

func (wc *WeChatPay) QueryOrder() error {
	requestURL := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s?mchid=%s", "D20240810234309-admin", "1680676110")
	//requestURL := "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"

	// 生成随机字符串和时间戳
	nonceStr := "593BEC0C930BF1AFEB40B4A08C8FB142"
	timestamp := time.Now().Unix()
	canonicalUrl := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s?mchid=%s", "D20240810234309-admin", "1680676110")
	// 构建待签名消息

	message := fmt.Sprintf("%s\n%s\n%d\n%s\n\n", "GET", canonicalUrl, timestamp, nonceStr)

	// 生成签名
	signature, err := wc.generateSignature([]byte(message))
	if err != nil {
		return err
	}

	// 构建 Authorization 请求头
	authorization := fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",timestamp="%d",serial_no="%s",signature="%s"`,
		wc.MchID, nonceStr, timestamp, wc.CertificateSerialNo, signature)

	// 创建 HTTP 客户端并发送请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Authorization", authorization)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // 使用 io.ReadAll 替换 ioutil.ReadAll
	if err != nil {
		return err
	}
	fmt.Printf("Response:", string(body))
	return nil
}
