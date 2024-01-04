package compose

import (
	"embed"
	"encoding/json"
	postController "example/src/controllers/postController"
	userControler "example/src/controllers/userController"
	"example/src/model"
	"log"
)

//go:embed data/*.json
var content embed.FS

type compose struct {
	HUser *model.HandlerUser
	CUser *userControler.UserController
	HPost *model.HandlerPost
	CPost *postController.PostController
}

func NewCompose() *compose {
	fileContent, err := content.ReadFile("data/users.json")
	if err != nil {
		log.Panic(err)
		return nil
	}

	users := []model.User{}
	if err := json.Unmarshal(fileContent, &users); err != nil {
		log.Panic(err)
		return nil
	}

	fileContent, err = content.ReadFile("data/posts.json")
	if err != nil {
		log.Panic(err)
		return nil
	}

	posts := []model.Post{}
	if err := json.Unmarshal(fileContent, &posts); err != nil {
		log.Panic(err)
		return nil
	}

	hUser := model.NewHandlerUser(users)
	cUser := userControler.NewUserController(hUser)

	hPost := model.NewHandlerPost(posts)
	cPost := postController.NewPostController(hPost)

	return &compose{
		HUser: hUser,
		CUser: cUser,
		CPost: cPost,
		HPost: hPost,
	}
}