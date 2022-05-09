package model

import "testing"

func TestCompany_GetCompanyType(t *testing.T) {
	c := Company{
		ID:      123,
		Name:    "Google.LTD",
		Country: "USA",
	}

	companyType := c.GetCompanyType() //获取类型

	if companyType != "Limited Liability Company" { //判断
		t.Error("this is others!")
	}
}

//如果测试通过，则会直接显示PASS
//如果测试不通过。结果如下：
/*
UN   TestCompany_GetCompanyType
    company_test.go:15: this is others!
--- FAIL: TestCompany_GetCompanyType (0.00s)

FAIL
*/
