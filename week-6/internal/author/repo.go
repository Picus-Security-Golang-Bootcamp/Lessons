package author

import (
	"github.com/h4yfans/patika-bookstore/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) create(a *models.Author) (*models.Author, error) {
	zap.L().Debug("author.repo.create", zap.Reflect("authorBody", a))
	if err := r.db.Create(a).Error; err != nil {
		zap.L().Error("author.repo.Create failed toΩ create author", zap.Error(err))
		return nil, err
	}

	return a, nil
}

func (r *AuthorRepository) getByID(id string) (*models.Author, error) {
	zap.L().Debug("author.repo.getByID", zap.Reflect("id", id))

	var author = &models.Author{}
	if result := r.db.Preload("Books").First(&author, id); result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

func (r *AuthorRepository) update(a *models.Author) (*models.Author, error) {
	zap.L().Debug("author.repo.update", zap.Reflect("author", a))

	if result := r.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (r *AuthorRepository) delete(id string) error {
	zap.L().Debug("author.repo.delete", zap.Reflect("id", id))

	author, err := r.getByID(id)
	if err != nil {
		return err
	}

	if result := r.db.Delete(&author); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AuthorRepository) Migration() {
	r.db.AutoMigrate(&models.Author{})
}
