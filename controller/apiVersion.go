package controller

import "net/http"
import (
	"fmt"
	"encoding/json"
)



func ApiVersionGet(w http.ResponseWriter, r *http.Request){
	fmt.Println("ApiVersionGet")
	v := ApiVersionResponse{Version : `v2`}
	j, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprint(w, string(j))
}

type ApiVersionResponse struct {
	Version string `json:"version"`
}