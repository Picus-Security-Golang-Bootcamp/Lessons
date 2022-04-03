package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/h4yfans/patika-bookstore/internal/auth"
	"github.com/h4yfans/patika-bookstore/internal/author"
	"github.com/h4yfans/patika-bookstore/internal/book"
	"github.com/h4yfans/patika-bookstore/pkg/config"
	db "github.com/h4yfans/patika-bookstore/pkg/database"
	"github.com/h4yfans/patika-bookstore/pkg/graceful"
	logger "github.com/h4yfans/patika-bookstore/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	log.Println("Book store service starting...")

	// Set envs for local development
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// Set global logger
	logger.NewLogger(cfg)
	defer logger.Close()

	// Connect DB
	// Use golang-migrate instead of gorm auto migrate
	//https://github.com/golang-migrate/migrate
	DB := db.Connect(cfg)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	// Router group
	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
	authorRouter := rootRouter.Group("/authors")
	bookRouter := rootRouter.Group("/books")
	authRooter := rootRouter.Group("/user")

	// Book Repository
	bookRepo := book.NewBookRepository(DB)
	bookRepo.Migration()
	book.NewBookHandler(bookRouter, bookRepo)

	// Author Repository
	authorRepo := author.NewAuthorRepository(DB)
	authorRepo.Migration()
	author.NewAuthorHandler(authorRouter, authorRepo)

	auth.NewAuthHandler(authRooter, cfg)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.GET("/readyz", func(c *gin.Context) {
		db, err := DB.DB()
		if err != nil {
			zap.L().Fatal("cannot get sql database instance", zap.Error(err))
		}
		if err := db.Ping(); err != nil {
			zap.L().Fatal("cannot ping database", zap.Error(err))
		}
		c.JSON(http.StatusOK, nil)
	})

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Book store service started")
	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))
}
