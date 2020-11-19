package domain

import (
	"context"
	"time"
)

// Post ...
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}



// PostService represent the post's usecases
type PostService interface {
	GetPosts(ctx context.Context, limit int) (res []Post)
	CreatePost(post *Post) (err error)
}

// PostRepository represent the post's repository contract
type PostRepository interface {
	GetPosts(ctx context.Context, limit int) (res []Post)
	CreatePost(post *Post) (err error)
}
