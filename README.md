# ExecutionTime
 
Execution Time middleware for Hertz framework.

- Adds an handler cost time using the `X-Execution-Time` header. 

## Install

```shell
go get github.com/lwch/executiontime
```

## Usage

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/lwch/executiontime"
)

func main() {
	h := server.Default()

	h.Use(
		executiontime.New(),
	)

	// Example ping request.
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.Spin()
}
```