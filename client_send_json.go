package zdpgo_httprouter

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func SendJson(method, targetUrl string, jsonData map[string]interface{}) (*http.Response, error) {
	if method == "" {
		return nil, errors.New("[SendJson] method is empty")
	}
	if targetUrl == "" {
		return nil, errors.New("[SendJson] targetUrl is empty")
	}
	if jsonData == nil {
		return nil, errors.New("[SendJson] formData is empty")
	}

	// 构造参数
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	jsonBytesReader := bytes.NewReader(jsonBytes)

	// 生成post请求
	client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(method), targetUrl, jsonBytesReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Do方法发送请求
	return client.Do(req)
}
