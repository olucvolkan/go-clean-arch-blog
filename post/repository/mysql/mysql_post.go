package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/olucvolkan/go-clean-arch-blog/domain"
	_ "github.com/sirupsen/logrus"
)

type mysqlPostRepository struct {
	Conn *gorm.DB
}

// NewMysqlPostRepository will create an object that represent the article.Repository interface
func NewMysqlPostRepository(conn *gorm.DB) *mysqlPostRepository {
	return &mysqlPostRepository{conn}
}

func (m mysqlPostRepository) GetPosts(ctx context.Context, limit int) (res []domain.Post) {
	var posts []domain.Post

	if err := m.Conn.Limit(limit).Find(&posts).Error; err != nil {
		return nil
	}

	return posts
}

func (m mysqlPostRepository) CreatePost(post *domain.Post) (err error) {

	if err := m.Conn.Create(post).Error; err != nil {
		return err
	}
	return nil

}
