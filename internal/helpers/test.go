package helpers

import (
	"io"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type TestFuncCall struct {
	Called bool
	Input  []interface{}
	Output []interface{}
}

func CreateGinContext(method string, target string, body io.Reader) (c *gin.Context, w *httptest.ResponseRecorder) {
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(method, target, body)
	return
}
