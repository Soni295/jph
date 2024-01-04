package model

type CommentId uint64

type Comment struct {
	Id     CommentId `json:"id"`
	PostId PostId    `json:"postId"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Body   string    `json:"body"`
}