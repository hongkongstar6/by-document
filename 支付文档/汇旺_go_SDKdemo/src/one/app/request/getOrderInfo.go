package request

import (
	"GoHuioneSDKdemo/src/one/app/common"
	"GoHuioneSDKdemo/src/one/app/config"
	"GoHuioneSDKdemo/src/one/app/utils"
	"encoding/json"
	"io"
	"log"
	"time"
)

// 获取订单信息
func GetOrderInfo() {
	// 构造请求参数
	paramMap := map[string]interface{}{
		"nonce":     utils.GetNonce(),
		"timestamp": time.Now().UnixMilli(),
		"outTradeNoList": []string{
			"One25051917510034796",
			"One25061111425919274",
			"One250612035203000894",
		},
	}

	// 加签
	sign, err := utils.Sign(paramMap, config.HuiOneWeb.PrivateKey)
	if err != nil {
		log.Println("签名失败:", err)
	}
	paramMap["sign"] = sign
	// 发起 POST 请求,完整地址是host+path
	host := config.Host[config.Uat]
	path := config.PathEnum[config.QueryOrder]
	url := host + path
	appId := config.HuiOneWeb.AppId
	resp, _ := utils.HttpPost(url, paramMap, appId)
	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Println("响应内容:", string(bodyBytes))

	// 解析响应
	var res GetOrderInfoResponse
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		log.Println("响应解析失败: ",err,"响应内容:", string(bodyBytes))
	}

	// 打印订单信息
	for _, order := range res.Data {
		log.Println("订单信息:", order)
	}
}

type GetOrderInfoResponse struct {
	common.BaseResponse
	Data []interface{} `json:"data"`
}
