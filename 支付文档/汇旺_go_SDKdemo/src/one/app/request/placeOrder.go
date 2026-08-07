package request

import (
	"GoHuioneSDKdemo/src/one/app/common"
	"GoHuioneSDKdemo/src/one/app/config"
	"GoHuioneSDKdemo/src/one/app/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"time"
)

// 创建订单
func PlaceOrder() {
	// 构造请求参数
	paramMap := map[string]interface{}{
		"amount":      big.NewFloat(100),
		"currency":    "USD",
		"description": "汇旺支付测试",
		"outTradeNo":  fmt.Sprintf("GO%s%d", time.Now().Format("060102150405"), rand.Intn(1000)),
		"attach":      "汇旺支付测试",
		"timeExpire":  600,
		"timestamp":   time.Now().UnixMilli(),
		"nonce":       utils.GetNonce(),
	}
	pri := config.HuiOneWeb.PrivateKey
	// 加签
	sign, err := utils.Sign(paramMap, pri)
	if err != nil {
		log.Println("签名失败:", err)
	}
	// 添加签名到参数中
	paramMap["sign"] = sign
	log.Println("参数:", paramMap)

	// TODO:
	// 验签
	//verify, err := utils.Verify(sign, paramMap, config.HuiOneWeb.PublicKey)
	//if err != nil {
	//	log.Fatal("验签失败:", err)
	//}
	//log.Println("验签结果:", verify)

	// TODO:
	// 发起 POST 请求,完整地址是host+path
	host := config.Host[config.Uat]
	path := config.PathEnum[config.CreatePrepayOrder]
	url := host + path
	appId := config.HuiOneWeb.AppId
	resp, _ := utils.HttpPost(url, paramMap, appId)
	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Println("响应内容:", string(bodyBytes))

	// 到这里已经响应了响应内容: {"code":"000000","data":{"fee":null,"preTransactionId":null,"qrCode":"huione://openReceive?code=SDK1944681080734908418","scheme":"huione://huione/payOrder?data=u0xkVwiNp0XMiOZBfITminIXAPoipmieeUMSMj7d8loTeAQKzpFJXHTMHGAzGsomoVv6wColMwvt8750i9MdI6NIg1AE1skqflgEt8VY2NTsdxVlasDcaB5Uz0XbaryLucEeKYnQw8xdwxWMk24m68i8qekNpDbSMw%2BLMRqNvbk%3D"},"msg":"success","success":true,"systemTime":1752483074758,"traceId":"35a67f0221c48f52"}
	// 获取data
	var res PlaceOrderResponse
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		log.Println("响应解析失败:", err)
	}

	log.Println("data:", res.Data)
}

type PlaceOrderResponse struct {
	common.BaseResponse
	Data struct {
		QRCode string `json:"qrCode"`
	} `json:"data"`
}
