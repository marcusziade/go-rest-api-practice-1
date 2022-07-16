package comment

import (
	"time"

	"github.com/jinzhu/gorm"
)

// The comment service
type Service struct {
	Database *gorm.DB
}

// Comment
type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
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

// Returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		Database: db,
	}
}

// Retrieves comments by their ID from the database
func (service *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := service.Database.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// Retrieves all comments by slug (path - /article/name/)
func (service *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := service.Database.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// Adds a new comment to the database
func (service *Service) PostComment(comment Comment) (Comment, error) {
	if result := service.Database.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// Updates a comment by ID with new comment info.
func (service *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, error := service.GetComment(ID)
	if error != nil {
		return Comment{}, error
	}

	if result := service.Database.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// Deletes a comment from the database by ID
func (service *Service) DeleteComment(ID uint) error {
	if result := service.Database.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}
