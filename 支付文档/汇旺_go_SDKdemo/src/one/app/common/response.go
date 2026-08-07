package common

type BaseResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	Success    bool   `json:"success"`
	SystemTime int64  `json:"systemTime"`
	TraceId    string `json:"traceId"`
}
