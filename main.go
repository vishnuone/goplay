package main

import (
	"net/http"
	"github.com/vishnuone/goplay/controller"
)

func main() {
	prefix := `/v1`
	host := ``
	port := `8080`


	http.Handle(`/`, basic(http.HandlerFunc(controller.ApiVersionGet)))
	http.Handle(prefix + `/`, basic(auth(http.HandlerFunc(controller.ApiVersionGet))))

	http.ListenAndServe(host + `:` + port , nil)
}
