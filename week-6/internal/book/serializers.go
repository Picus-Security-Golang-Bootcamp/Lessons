package book

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/h4yfans/patika-bookstore/internal/api"
	"github.com/h4yfans/patika-bookstore/internal/models"
	"gorm.io/gorm"
)

func BookToResponse(b *models.Book) *api.Book {
	date, _ := time.Parse(strfmt.RFC3339FullDate, b.ReleaseDate)
	releaseDate := strfmt.Date(date)

	return &api.Book{
		Author: &api.Author{
			ID:   int64(b.Author.ID),
			Name: b.Author.Name,
		},
		Genre:       &b.Genre,
		ID:          int64(b.ID),
		Isbn:        &b.Isbn,
		Name:        b.Name,
		PageNumber:  &b.PageNumber,
		ReleaseDate: &releaseDate,
	}
}

func booksToResponse(bs *[]models.Book) []*api.Book {
	books := make([]*api.Book, 0)
	for _, b := range *bs {
		books = append(books, BookToResponse(&b))
	}
	return books
}

func responseToBook(b *api.Book) *models.Book {
	return &models.Book{
		Model:       gorm.Model{ID: uint(b.ID)},
		Name:        b.Name,
		Genre:       *b.Genre,
		Isbn:        *b.Isbn,
		PageNumber:  *b.PageNumber,
		ReleaseDate: b.ReleaseDate.String(),
		AuthorID:    uint(b.Author.ID),
	}
}
