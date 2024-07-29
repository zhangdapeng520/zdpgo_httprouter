package zdpgo_httprouter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseJsonArr 返回JSON字典类型的信息
// @param w 网络输出流对象
// @param m 要返回的字典信息
func ResponseJsonArr(w http.ResponseWriter, arr []interface{}) {
	// 声明返回的是JSON数据
	w.Header().Set("Content-Type", "application/json")

	// 将字段转换为JSON
	jsonBytes, err := json.Marshal(arr)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// 返回JSON信息
	fmt.Fprint(w, string(jsonBytes))
}
