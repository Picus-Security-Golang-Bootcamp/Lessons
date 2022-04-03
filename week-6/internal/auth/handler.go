package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/h4yfans/patika-bookstore/internal/api"
	httpErr "github.com/h4yfans/patika-bookstore/internal/httpErrors"
	"github.com/h4yfans/patika-bookstore/pkg/config"
	jwtHelper "github.com/h4yfans/patika-bookstore/pkg/jwt"
	mw "github.com/h4yfans/patika-bookstore/pkg/middleware"
)

type authHandler struct {
	cfg *config.Config
}

func NewAuthHandler(r *gin.RouterGroup, cfg *config.Config) {
	a := authHandler{
		cfg: cfg,
	}
	r.POST("/login", a.login)

	r.Use(mw.AuthMiddleware(cfg.ServerConfig.Mode))
	//r.POST("/decode", a.VerifyToken)

}

func (a *authHandler) login(c *gin.Context) {
	var req api.Login
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
	}
	user := GetUser(*req.Email, *req.Password)
	if user == nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "user not found", nil)))

	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"iss":    os.Getenv("ENV"),
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"roles":  user.Roles,
	})
	token := jwtHelper.GenerateToken(jwtClaims, a.cfg.JWTConfig.SecretKey)
	c.JSON(http.StatusOK, token)
}

func (a *authHandler) VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, decodedClaims)
}
