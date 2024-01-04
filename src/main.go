package main

import (
	"example/src/compose"
	"example/src/utils"

	"github.com/gin-gonic/gin"
)

const (
	PORT = ":3000"
)

func main() {
	comp := compose.NewCompose()
	cu := comp.CUser
	cp := comp.CPost

	r := gin.Default()
	r.Use(utils.HandleErrors)

	r.Group("users").
		GET("", cu.GetAll).
		GET("/:id", cu.GetById).
		POST("", cu.CreateUser).
		DELETE("/:id", cu.DeleteById)

	r.Group("posts").
		GET("", cp.GetAll).
		GET("/:id", cp.GetById).
		POST("", cp.CreatePost).
		DELETE("/:id", cp.DeleteById)

	r.Run(":3000")
}