package utils

import (
	httperr "example/src/httpErr"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrors(ctx *gin.Context) {
	ctx.Next()
	log.Println("Error")
	errorToPrint := ctx.Errors.ByType(gin.ErrorTypePublic).Last()

	if errorToPrint != nil {
		type hErr struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}

		body := hErr{Message: errorToPrint.Error()}
		switch body.Message {

		case httperr.ErrNotFound.Error():
			body.Status = http.StatusNotFound
		case httperr.ErrBadRequest.Error():
			body.Status = http.StatusInternalServerError
		case httperr.ErrInternalFailure.Error():
			body.Status = http.StatusInternalServerError
		case httperr.ErrNotFound.Error():
			body.Status = http.StatusNotFound
		case httperr.ErrUnprocessableEntity.Error():
			body.Status = http.StatusUnprocessableEntity
		default:
			body.Status = http.StatusInternalServerError
		}

		ctx.JSON(body.Status, body)
		return
	}
}