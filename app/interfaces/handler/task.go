package handler

import (
	"net/http"

	"github.com/39shin52/todoAPI/app/domain/entity"
	"github.com/39shin52/todoAPI/app/interfaces/request"
	"github.com/39shin52/todoAPI/app/interfaces/response"
	"github.com/39shin52/todoAPI/app/usecase"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskUsecase *usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}

func (th *TaskHandler) CreateTask(c *gin.Context) {
	var task request.CreateTaskRequest
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	response, err := th.taskUsecase.CreateTask(c.Request.Context(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task_id": response,
	})

}

func (th *TaskHandler) GetTasks(c *gin.Context) {
	var user request.GetTasksRequest
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	tasks, err := th.taskUsecase.ReadTasks(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	tasksResponse := make([]response.GetTaskResponse, 0)
	for _, t := range tasks {
		newTaskResponse := response.GetTaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsComplete:  t.IsComplete,
			Created_at:  t.Created_at,
			Updated_at:  t.Updated_at,
		}

		tasksResponse = append(tasksResponse, newTaskResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": tasksResponse,
	})
}

func (th *TaskHandler) GetTask(c *gin.Context) {
	taskID := c.Param("taskID")

	task, err := th.taskUsecase.ReadTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	taskResponse := response.GetTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsComplete:  task.IsComplete,
		Created_at:  task.Created_at,
		Updated_at:  task.Updated_at,
	}

	c.JSON(http.StatusOK, gin.H{
		"response": taskResponse,
	})
}

func (th *TaskHandler) UpdateTask(c *gin.Context) {
	taskID := c.Param("taskID")
	var reqTask request.UpdateTaskRequest
	if err := c.Bind(&reqTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	newTask := &entity.Task{
		Title:       reqTask.Title,
		Description: reqTask.Description,
		IsComplete:  reqTask.IsComplete,
	}

	if err := th.taskUsecase.UpdateTask(c, taskID, newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, "Success")
}

func (th *TaskHandler) DeleteTask(c *gin.Context) {
	taskID := c.Param("taskID")

	if err := th.taskUsecase.DeleteTask(c, taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, "Success")
}

func (th *TaskHandler) DuplicateTask(c *gin.Context) {
	taskID := c.Param("taskID")

	if err := th.taskUsecase.DuplicateTask(taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, "Success")
}
