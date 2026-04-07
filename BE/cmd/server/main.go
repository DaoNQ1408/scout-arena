package main

import (
	"scout-arena/internal/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	validator.Init()
	r.Run(":8080")
}
