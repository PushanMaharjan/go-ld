package autodocumentation

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitiateDocumentation(routes gin.RoutesInfo) {
	for _, routeInfo := range routes {
		fmt.Printf("Method: %s, Path: %s, Handler: %s\n", routeInfo.Method, routeInfo.Path, routeInfo.Handler)
	}
}
