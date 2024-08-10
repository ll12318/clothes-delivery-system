package initialize

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/wechat"
	"math/rand"
	"os"
	"time"
)

// 声明微信支付参数

func InitWeChatPay() {
	privateKey, err := loadPrivateKey("cert/apiclient_key.pem")
	if err != nil {
		panic(err)
	}
	global.WeChatPay = wechat.WeChatPay{
		MchID:               "1680676110",
		AppID:               "wx64982e69637a1fef",
		APIKey:              "23A4533B0D986C18FAC58FC69889B257",
		CertPath:            "cert/apiclient_cert.pem",
		PrivateKey:          privateKey,
		CertificateSerialNo: "2D23E41523A4533B0D986C18FAC58FC69889B257", // 替换为实际的证书序列号
	}

	fmt.Println("InitWeChatPay success")

	//orderRequest := wechat.OrderRequest{
	//	AppID:       global.WeChatPay.AppID,
	//	MchID:       global.WeChatPay.MchID,
	//	Description: "Test order",
	//	OutTradeNo:  GenerateOutTradeNo(), // 使用唯一订单号生成函数
	//	NotifyURL:   "https://your_notify_url",
	//	Amount: wechat.Amount{
	//		Total:    100,
	//		Currency: "CNY",
	//	},
	//	Payer: wechat.Payer{
	//		OpenID: "o1u8W7abV4UTsrgfuk6ewZ8gZa2c",
	//	},
	//}
	//// 发送创建订单请求
	//resp, err := global.WeChatPay.CreateOrder(orderRequest)
	//if err != nil {
	//	log.Fatalf("Error creating order: %v", err)
	//}
	//
	//// 打印响应
	//fmt.Printf("resp: %v\n", resp)
	//fmt.Println("InitWeChatPay", privateKey)
	//fmt.Println("WeChatPay", global.WeChatPay)
	//fmt.Println("OrderRequest", orderRequest)

}
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey.(*rsa.PrivateKey), nil
}

func GenerateOutTradeNo() string {
	timestamp := time.Now().Unix()
	randomInt := rand.Intn(100000)
	return fmt.Sprintf("order_%d_%d", timestamp, randomInt)
}
