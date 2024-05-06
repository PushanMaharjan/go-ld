package autodocumentation

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitiateDocumentation() {
	r := gin.Default()

	fmt.Println(r.Routes())
}
