package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/39shin52/todoAPI/app/config"
	"github.com/39shin52/todoAPI/app/domain/repository/transaction"
	"github.com/39shin52/todoAPI/app/infrastructure"
	"github.com/39shin52/todoAPI/app/interfaces/handler"
	"github.com/39shin52/todoAPI/app/usecase"
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
	conn, err := config.CreateDBConnection()
	if err != nil {
		return err
	}

	s.db = conn
	s.Route()

	return nil
}

func (s *Server) Route() {
	// 依存性注入
	// dig使う？
	txAdmin := transaction.NewTxRepository(s.db)

	taskRepository := infrastructure.NewTaskRepository(s.db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, txAdmin)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	s.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	taskGroup := s.Router.Group("/task")
	taskGroup.GET("", taskHandler.GetTasks)
	taskGroup.GET("/task_id/:taskID", taskHandler.GetTask)
	taskGroup.POST("", taskHandler.CreateTask)
	taskGroup.PUT("/task_id/:taskID", taskHandler.UpdateTask)
	taskGroup.DELETE("/task_id/:taskID")
	taskGroup.POST("/task_id/:taskID")

}
