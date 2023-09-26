package helper

type BaseHttpResponse struct {
	Result     any        `json:"result"`
	Success    bool       `json:"success"`
	ResultCode ResultCode `json:"resultCode"`
}

func GenerateBaseResponse(result any, success bool, resultCode ResultCode) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}
