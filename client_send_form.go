package zdpgo_httprouter

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
)

func SendForm(targetUrl string, formData map[string]string) (*http.Response, error) {
	if targetUrl == "" {
		return nil, errors.New("[SendForm] targetUrl is empty")
	}
	if formData == nil {
		return nil, errors.New("[SendForm] formData is empty")
	}

	// 构造表单参数
	formValues := url.Values{}
	for k, v := range formData {
		formValues.Set(k, v)
	}
	formValues.Set("age", "23")
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)

	// 生成post请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", targetUrl, formBytesReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Do方法发送请求
	return client.Do(req)
}
