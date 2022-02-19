package base

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

func FeatureTest(handler gin.HandlerFunc, method string, body io.Reader, paramters string) *httptest.ResponseRecorder {
	rPath := "/test"
	router := gin.Default()

	if method == "GET"{
		router.GET(rPath, handler)
	} else if method == "POST"{
		router.POST(rPath, handler)
	} else if method == "PUT"{
		router.PUT(rPath, handler)
	} else if method == "DELETE"{
		router.DELETE(rPath, handler)
	}

	req, _ := http.NewRequest(method, rPath , body)
	req.URL.RawQuery = paramters
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
