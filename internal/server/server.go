package server

import (
	"net/http"

	"github.com/cryring/blog_backend/internal/config"
	"github.com/cryring/blog_backend/internal/dao/mysql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg Config
	db  *mysql.Dao
}

func New(cfg Config) (*Server, error) {
	db, err := mysql.New(config.GetConfig().DBConfig.DSN())
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(); err != nil {
		return nil, err
	}

	return &Server{
		cfg: cfg,
		db:  db,
	}, nil
}

func (srv *Server) Run() error {
	return srv.SetupRouter().Run(srv.cfg.ListenAddr)
}

func (srv *Server) SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/api/blogs/:category", srv.handleBlogs)
	r.GET("/api/blog/:category/:filename", srv.handlBlogContent)

	return r
}
