package autodocumentation

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitiateDocumentation(routes gin.RoutesInfo) {
	fmt.Println(routes)
}
