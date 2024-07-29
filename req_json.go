package zdpgo_httprouter

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetJson(r *http.Request, jsonObj interface{}) error {
	r.ParseForm()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, jsonObj)
}
