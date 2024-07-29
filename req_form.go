package zdpgo_httprouter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetForm(r *http.Request, key string) string {
	r.ParseForm()
	fmt.Println(r.PostForm)
	return r.PostForm.Get(key)
}

func GetFormInt(r *http.Request, key string) int {
	age := GetForm(r, key)
	ageInt, err := strconv.ParseInt(age, 10, 64)
	if err != nil {
		log.Panicln("GetFormInt: strconv.ParseInt(age, 10, 64) error: ", err.Error())
		return 0
	}
	return int(ageInt)
}
