package executiontime

import (
	"context"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	hzconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
)

const (
	testXRequestID  = "test-request-id"
	customHeaderKey = "customKey"
)

func emptySuccessResponse(ctx context.Context, c *app.RequestContext) {
	c.String(http.StatusOK, "")
}

func hertzHandler(middleware app.HandlerFunc) *route.Engine {
	r := route.NewEngine(hzconfig.NewOptions([]hzconfig.Option{}))
	r.Use(middleware)
	r.GET("/", emptySuccessResponse)

	return r
}

func TestCreateNewRequestID(t *testing.T) {
	r := hertzHandler(New())
	w := ut.PerformRequest(r, http.MethodGet, "/", nil)

	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.NotEqual(t, "", string(w.Header().Peek(headerXExecutionTime)))
}

func TestRequestIDWithCustomHeaderKey(t *testing.T) {
	r := hertzHandler(New(
		WithCustomHeaderStrKey(customHeaderKey),
	))

	w := ut.PerformRequest(r, http.MethodGet, "/", nil)

	assert.DeepEqual(t, http.StatusOK, w.Code)
	var found bool
	w.Header().VisitAll(func(key, value []byte) {
		if string(key) == customHeaderKey {
			found = true
		}
	})
	assert.Assert(t, !found)
}
