package postcontroller

import (
	"example/src/model"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	handlerPost *model.HandlerPost
	errLogger   *log.Logger
}

func NewPostController(handlerPost *model.HandlerPost) *PostController {
	return &PostController{
		handlerPost: handlerPost,
		errLogger:   log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	newPost := model.Post{}

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := c.handlerPost.CreatePost(&newPost)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *PostController) GetAll(ctx *gin.Context) {
	users, err := c.handlerPost.GetAll()
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *PostController) GetById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	users, err := c.handlerPost.GetById(model.PostId(id))
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *PostController) DeleteById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	users, err := c.handlerPost.DeleteById(model.PostId(id))
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, users)
}