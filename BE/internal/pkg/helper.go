package pkg

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get id param from url and convert to uint, if error return 0 and false, else return id and true
func GetIDParam(c *gin.Context, name string) (uint, bool) {
	val := c.Param(name)
	id, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tham số " + name + " không hợp lệ"})
		return 0, false
	}
	return uint(id), true
}
