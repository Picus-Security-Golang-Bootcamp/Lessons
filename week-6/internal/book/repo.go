package book

import (
	"github.com/h4yfans/patika-bookstore/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) create(b *models.Book) (*models.Book, error) {
	zap.L().Debug("book.repo.create", zap.Reflect("book", b))

	if err := r.db.Create(b).Error; err != nil {
		zap.L().Error("book.repo.Create failed toΩ create book", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (r *BookRepository) getAll(pageIndex, pageSize int) (*[]models.Book, int) {
	zap.L().Debug("book.repo.getAll")

	var bs = &[]models.Book{}
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Preload("Author").Find(&bs).Count(&count)

	return bs, int(count)
}

func (r *BookRepository) getByID(id string) (*models.Book, error) {
	zap.L().Debug("book.repo.getByID", zap.Reflect("id", id))

	var book = &models.Book{}
	if result := r.db.Preload("Author").First(&book, id); result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (r *BookRepository) update(a *models.Book) (*models.Book, error) {
	zap.L().Debug("book.repo.update", zap.Reflect("book", a))

	if result := r.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (r *BookRepository) delete(id string) error {
	zap.L().Debug("book.repo.delete", zap.Reflect("id", id))

	book, err := r.getByID(id)
	if err != nil {
		return err
	}

	if result := r.db.Delete(&book); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&models.Book{})
}
