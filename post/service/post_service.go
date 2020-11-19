package service

import (
	"context"
	"github.com/olucvolkan/go-clean-arch-blog/domain"
)

type servicePostService struct {
	postRepository domain.PostRepository
}

func NewPostService(p domain.PostRepository) *servicePostService {
	return &servicePostService{postRepository: p}
}

func (s servicePostService) GetPosts(ctx context.Context, limit int) (res []domain.Post) {

	if limit == 0 {
		limit = 10
	}
	res = s.postRepository.GetPosts(ctx, limit)

	return res

}

func (s servicePostService) CreatePost(post *domain.Post) (err error) {

	res := s.postRepository.CreatePost(post)

	if res != nil {
		return res
	}

	return  nil
}


