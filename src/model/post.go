package model

type PostId uint64

type Post struct {
	Id     PostId `json:"id"`
	UserId UserId `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}