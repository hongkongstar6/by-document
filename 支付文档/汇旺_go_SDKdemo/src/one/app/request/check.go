package request

import (
	"GoHuioneSDKdemo/src/one/app/config"
	"GoHuioneSDKdemo/src/one/app/utils"
	"encoding/json"
	"log"
)

// 对汇旺回调验签
func Check() {
	p := "{\"appId\":\"2611782313897625313282\",\"attach\":\"汇旺支付测试\",\"hash\":\"F2h5HnFqkMZ4Tcozmge2k7eHDvbRHx1Gpv5G15Qx1b15\",\"merchantId\":\"1780857138277335042\",\"nonce\":\"nqpyy\",\"outTradeNo\":\"GO250715175739495\",\"sign\":\"RtA32qHxhSwt6Hyq6ITdqCVlZAbeoKSP8215UTkAX5qVyXqh7AfSjmw8PedIZmebu+3j1GX3m0FEPF+GXsWw4zKoxcPbdYegmLrzG9DDAuAhO4DGI9BDO91BLCnkHeizRnoib5tOKM4P8cVajEwUbIxG1MO/o+crDmZ/xF0zB9A=\",\"status\":\"DONE_PAYMENT\",\"timestamp\":1752577090446,\"transactionId\":\"7344223286453694464\"}"

	var data map[string]interface{}
	err := json.Unmarshal([]byte(p), &data)
	if err != nil {
		log.Println("JSON 解析失败: ", err)
	}

	log.Println("data:", data)

	serverPublicKey := config.HuiOneWeb.ServerPublicKey
	// 验签
	verify, err := utils.Verify(data["sign"].(string), data, serverPublicKey)
	if err != nil {
		log.Println("验签失败: ", err)
	} else {
		log.Println("verify:", verify)
	}
}
