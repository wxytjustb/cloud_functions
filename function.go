// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"
)

type RequestData struct {
	Action string 					`json:"action"`
	Data map[string]string 			`json:"data"`
}

type ResponseData struct {
	Status int						`json:"status"`
	Data map[string]interface{}		`json:"data"`
	Msg string 						`json:"msg"`
}



// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func Api(w http.ResponseWriter, r *http.Request) {
	var d RequestData
	if (r.Method == "GET") {
		r.ParseForm()
		action := r.Form["action"][0]
		if action == "now_time"{
			json.NewEncoder(w).Encode(
				&ResponseData{
					0,
					map[string]interface{}{
						"now_time": getNowTime(),
					},
					"",
				})
		}

		return
	}else if (r.Method == "POST") {

	}else{
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "error")
		return
	}
}


func getNowTime() string{
	return strconv.FormatInt(time.Now().Unix(),10)
}
