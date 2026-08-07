package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// HttpPost 向指定 URL 发送 POST 请求，body 为 JSON 字符串，返回响应内容或错误
func HttpPost(url string, body interface{}, appid string) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("app-id", appid)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
