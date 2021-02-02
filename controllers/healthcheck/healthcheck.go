package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "Health Ok!")
}
