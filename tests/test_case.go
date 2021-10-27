package tests

import (
	"github.com/gin-gonic/gin"
	"go-api/route"
	"net/http"
	"net/http/httptest"
)

type testCase struct {
	 r *gin.Engine
}

var (
	TestCase testCase
)

func init()  {
	TestCase.r = SetupRouter()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	route.InitApiRoute(r)
	return r
}

// ParseToStr 将map中的键值对输出成querystring形式
func ParseToStr(mp map[string]string) string {
	if mp == nil {
		return ""
	}

	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

/**
get 请求
 */
func (t *testCase)GetRequest(uri string, mp map[string]string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	uri = uri + ParseToStr(mp)
	req, _ := http.NewRequest("GET", uri, nil)
	TestCase.r.ServeHTTP(w, req)
	return w
}

func (t *testCase)postRequest(uri string, param map[string]string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST",  uri+ParseToStr(param), nil)
	TestCase.r.ServeHTTP(w, req)
	return w
}

