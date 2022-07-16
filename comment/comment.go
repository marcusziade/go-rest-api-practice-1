package comment

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service - our comment service
type Service struct {
	DB *gorm.DB
}

// Comment
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
	Created time.Time
}

// CommentService -
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
