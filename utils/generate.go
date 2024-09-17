package util

type Response struct {
	ResponseCode    string
	ResponseMessage string
	Result          interface{}
}

func GenerateResponse(respcode string, respmsg string, result interface{}) *Response {
	return &Response{
		ResponseCode:    respcode,
		ResponseMessage: respmsg,
		Result:          result,
	}

}
