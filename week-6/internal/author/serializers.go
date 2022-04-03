package author

import (
	"github.com/h4yfans/patika-bookstore/internal/api"
	"github.com/h4yfans/patika-bookstore/internal/book"
	"github.com/h4yfans/patika-bookstore/internal/models"
	"gorm.io/gorm"
)

func authorToResponse(a *models.Author) api.Author {

	books := make([]*api.Book, 0)
	for _, b := range a.Books {
		books = append(books, book.BookToResponse(&b))
	}

	return api.Author{
		ID:    int64(a.ID),
		Name:  a.Name,
		Books: books,
	}
}

func responseToAuthor(a *api.Author) *models.Author {
	return &models.Author{
		Name:  a.Name,
		Model: gorm.Model{ID: uint(a.ID)},
	}
}
