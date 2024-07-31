package zdpgo_httprouter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResult struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var errorResult = ErrorResult{
	Code:    10500,
	Status:  false,
	Message: "server error",
	Data:    nil,
}

// ResponseError404 统一返回类型，错误信息资源不存在的返回
// @param w 网络输出流对象
// @param data 要返回的内容
func ResponseError404(w http.ResponseWriter, data interface{}) {
	// 声明返回的是JSON数据
	w.Header().Set("Content-Type", "application/json")

	// 构造错误信息
	result := ErrorResult{
		Code:    10404,
		Status:  false,
		Message: "not found",
		Data:    data,
	}

	// 将字段转换为JSON
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// 返回JSON信息
	fmt.Fprint(w, string(jsonBytes))
}
