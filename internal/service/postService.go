package service

import (
	"Shiro/internal/model"
	"Shiro/internal/repository"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (service *PostService) GetAllPosts() ([]model.Post, error) {
	return service.postRepo.GetAllPosts()
}
