package controller

import (
	"encoding/json"
	"go_web/model"
	"net/http"
)

func RegisterRoutesController() {
	http.HandleFunc("/companies", handlerCompany)
}

func handlerCompany(writer http.ResponseWriter, request *http.Request) {
	c := model.Company{
		ID:      123,
		Name:    "Google",
		Country: "USA",
	}
	//收到请求后使用JSON编码，并写到响应里面，并返回回去
	enc := json.NewEncoder(writer)
	enc.Encode(c)
}
