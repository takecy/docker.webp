package img

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Upload struct {
	File []byte `form:"file" json:"file" binding:"required"`
}

func New(router *gin.Engine) {
	router.POST("/:id/webp", ToWebp)
}

func ToWebp(c *gin.Context) {
	if c.Request.Body == nil {
		c.AbortWithError(400, errors.New("body.required"))
		return
	}

	data := Upload{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		c.AbortWithError(400, errors.New("body.parse.error"))
		return
	}

	// encode

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func encode(srcImg []byte) (dst []byte) {

	return
}
