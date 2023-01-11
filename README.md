# ExecutionTime
 
Execution Time middleware for gin framework.

- Adds an handler cost time using the `X-Execution-Time` header. 

## Install

```shell
go get github.com/lwch/executiontime
```

## Usage

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwch/executiontime"
)

func main() {
	h := gin.Default()

	h.Use(
		executiontime.New(),
	)

	// Example ping request.
	h.GET("/ping", func(g *gin.Context) {
		g.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	h.Run()
}
```