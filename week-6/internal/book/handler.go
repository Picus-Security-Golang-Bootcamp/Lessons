package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/h4yfans/patika-bookstore/internal/api"
	httpErr "github.com/h4yfans/patika-bookstore/internal/httpErrors"
)

type bookHandler struct {
	repo *BookRepository
}

func NewBookHandler(r *gin.RouterGroup, repo *BookRepository) {
	h := &bookHandler{repo: repo}

	r.GET("/", h.getAll)
	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (b *bookHandler) create(c *gin.Context) {
	bookBody := &api.Book{}
	if err := c.Bind(&bookBody); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := bookBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	book, err := b.repo.create(responseToBook(bookBody))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, BookToResponse(book))
}

func (b *bookHandler) getAll(c *gin.Context) {
	books, err := b.repo.getAll()
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, booksToResponse(books))
}

func (b *bookHandler) getByID(c *gin.Context) {
	book, err := b.repo.getByID(c.Param("id"))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, BookToResponse(book))
}

func (b *bookHandler) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
	}

	bookBody := &api.Book{ID: int64(id)}
	if err := c.Bind(&bookBody); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := bookBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	book, err := b.repo.update(responseToBook(bookBody))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, BookToResponse(book))
}

func (b *bookHandler) delete(c *gin.Context) {
	err := b.repo.delete(c.Param("id"))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
