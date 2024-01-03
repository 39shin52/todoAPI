package interfaces

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	Router *gin.Engine
}

// gin使ってみる

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (s *Server) Init() error {
	return nil
}

func (s *Server) Route() {
	// 依存性注入
	// dig使う？

}
