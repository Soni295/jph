package controller

import (
	httperr "example/src/httpErr"
	"example/src/model"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	handlerUser *model.HandlerUser
	errLogger   *log.Logger
}

func NewUserController(handlerUser *model.HandlerUser) *UserController {
	return &UserController{
		handlerUser: handlerUser,
		errLogger:   log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	newUser := model.User{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := c.handlerUser.CreateUser(&newUser)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Update(ctx *gin.Context) {
	newUser := model.User{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := c.handlerUser.Update(&newUser)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.handlerUser.GetAll()
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	users, err := c.handlerUser.GetById(model.UserId(id))
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) DeleteById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	users, err := c.handlerUser.DeleteById(model.UserId(id))
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func validateUser(user model.User) error {
	if user.Name == "" {
		return httperr.NewError(httperr.ErrBadRequest, ErrUserNameRequired)
	}
	if user.Email == "" {
		return httperr.NewError(httperr.ErrBadRequest, ErrUserEmailRequired)
	}
	return nil
}