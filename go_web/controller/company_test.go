package controller

import (
	"encoding/json"
	"go_web/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCompanyCorrect(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/companies", nil) //模拟一个请求
	w := httptest.NewRecorder()                                 //用于捕获和记录响应
	handlerCompany(w, r)                                        //传入Handler

	result, _ := ioutil.ReadAll(w.Result().Body) //获取响应，提取body

	c := model.Company{}
	json.Unmarshal(result, &c) //解码并重新放入c这个变量里边

	if c.ID != 123 { //然后判断
		t.Error("this is a failed")
	}
}

//通过就显示PASS
//否则显示：
/*
=== RUN   TestHandleCompanyCorrect
    company_test.go:23: this is a failed
--- FAIL: TestHandleCompanyCorrect (0.00s)

FAIL
*/
