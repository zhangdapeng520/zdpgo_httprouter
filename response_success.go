package zdpgo_httprouter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SuccessResult struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var successResult = SuccessResult{
	Code:    10000,
	Status:  true,
	Message: "success",
	Data:    nil,
}

// ResponseSuccess 统一返回类型，成功的返回
// @param w 网络输出流对象
// @param data 要返回的内容
func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	// 声明返回的是JSON数据
	w.Header().Set("Content-Type", "application/json")

	successResult.Data = data

	// 将字段转换为JSON
	jsonBytes, err := json.Marshal(successResult)
	successResult.Data = nil
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// 返回JSON信息
	fmt.Fprint(w, string(jsonBytes))
}
