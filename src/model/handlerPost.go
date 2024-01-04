package model

import (
	httperr "example/src/httpErr"
	"fmt"
)

type HandlerPost struct {
	posts []Post
}

func NewHandlerPost(posts []Post) *HandlerPost {
	return &HandlerPost{
		posts: posts,
	}
}

func (hp *HandlerPost) GetAll() ([]Post, error) { return hp.posts, nil }

func (hp *HandlerPost) CreatePost(post *Post) (*Post, error) {
	post.Id = PostId(len(hp.posts) + 1)
	hp.posts = append(hp.posts, *post)
	return post, nil
}

func (hp *HandlerPost) GetById(postId PostId) (*Post, error) {
	for _, currentPost := range hp.posts {
		if currentPost.Id == postId {
			return &currentPost, nil
		}

	}
	return nil, httperr.ErrNotFound
}

func RemoveElement[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func (hp *HandlerPost) Update(post *Post) (*Post, error) {
	for index, currentPost := range hp.posts {
		if currentPost.Id != post.Id {
			continue
		}
		hp.posts[index] = *post
		return post, nil
	}

	return nil, httperr.ErrInternalFailure
}

func (hp *HandlerPost) DeleteById(postId PostId) (*Post, error) {
	fmt.Println(postId)
	for index, currentPost := range hp.posts {
		if currentPost.Id != postId {
			continue
		}
		hp.posts = RemoveElement[Post](hp.posts, index)
		return &currentPost, nil
	}

	return nil, httperr.ErrNotFound
}