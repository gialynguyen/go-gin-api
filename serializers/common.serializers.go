package serializers

import (
	"encoding/json"
	"github.com/golang-gin/interface"
)

func SuccessBaseResponse(data ...interface{}) _interface.SuccessBaseResponse {
	resData := map[string]interface{}{}
	for _, t := range data {
		tJson, _ := json.Marshal(t)
		_  = json.Unmarshal(tJson, &resData)
	}
	return _interface.SuccessBaseResponse{Success: true, Data: resData}
}


func ErrorBaseResponse(errorCode int, errorMessage interface{}) _interface.ErrorBaseResponse {

	return _interface.ErrorBaseResponse{Success: false, Error: _interface.ErrorDetail{ErrorCode: errorCode, ErrorMes: errorMessage}}
}
