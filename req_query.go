package zdpgo_httprouter

import (
	"log"
	"net/http"
	"strconv"
)

func GetQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetQueryInt(r *http.Request, key string) int {
	value := r.URL.Query().Get(key)
	valueInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt(value, 10, 64) error: ", value)
		return 0
	}
	return int(valueInt)
}
