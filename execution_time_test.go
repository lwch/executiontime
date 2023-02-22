package executiontime

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	testXRequestID  = "test-request-id"
	customHeaderKey = "customKey"
)

func emptySuccessResponse(g *gin.Context) {
	g.String(http.StatusOK, "")
}

func handler(middleware gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(middleware)
	r.GET("/", emptySuccessResponse)
	return r
}

func TestCreateNewRequestID(t *testing.T) {
	r := handler(New())

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Header().Get(defaultHeaderKey))
}

func TestRequestIDWithCustomHeaderKey(t *testing.T) {
	r := handler(New(
		WithCustomHeaderStrKey(customHeaderKey),
	))

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var found bool
	for k := range w.Header() {
		if k == customHeaderKey {
			found = true
		}
	}
	assert.False(t, found)
}
