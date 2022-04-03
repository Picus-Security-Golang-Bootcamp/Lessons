package author

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/h4yfans/patika-bookstore/internal/api"
	httpErr "github.com/h4yfans/patika-bookstore/internal/httpErrors"
)

type authorHandler struct {
	repo *AuthorRepository
}

func NewAuthorHandler(r *gin.RouterGroup, repo *AuthorRepository) {
	h := &authorHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (a *authorHandler) create(c *gin.Context) {
	authorBody := &api.Author{}
	if err := c.Bind(&authorBody); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := authorBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	author, err := a.repo.create(responseToAuthor(authorBody))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, authorToResponse(author))
}

func (a *authorHandler) getByID(c *gin.Context) {
	author, err := a.repo.getByID(c.Param("id"))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, authorToResponse(author))
}

func (a *authorHandler) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
	}

	authorBody := &api.Author{
		ID: int64(id),
	}
	if err := c.Bind(&authorBody); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := authorBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	author, err := a.repo.update(responseToAuthor(authorBody))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, authorToResponse(author))
}

func (a *authorHandler) delete(c *gin.Context) {
	err := a.repo.delete(c.Param("id"))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
