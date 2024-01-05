package usecase

import (
	"testing"

	"github.com/39shin52/todoAPI/app/domain/entity"
	mock_repository "github.com/39shin52/todoAPI/app/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type taskMocks struct {
	taskMockRespository *mock_repository.MockTaskRepository
}

func newMocks(ctrl *gomock.Controller) (*TaskUsecase, *taskMocks) {
	mocks := &taskMocks{
		taskMockRespository: mock_repository.NewMockTaskRepository(ctrl),
	}

	taskUsecase := &TaskUsecase{
		taskRepository: mocks.taskMockRespository,
	}

	return taskUsecase, mocks
}

func TestTaskUsecase_ReadTasks(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		taskUsecase, mocks := newMocks(ctrl)
		tasks := []entity.Task{
			{
				ID:          "1",
				UserID:      "1",
				Title:       "title 1",
				Description: "description 1",
				IsComplete:  "1",
			},
			{
				ID:          "2",
				UserID:      "1",
				Title:       "title 2",
				Description: "description 2",
				IsComplete:  "0",
			},
		}
		userID := "1"

		mocks.taskMockRespository.EXPECT().SelectTasks(userID).Return(tasks, nil)
		got, err := taskUsecase.ReadTasks(userID)
		require.NoError(t, err)
		if err != nil {
			t.Errorf("error message: %v", err)
		}
		want := []entity.Task{
			{
				ID:          "1",
				UserID:      "1",
				Title:       "title 1",
				Description: "description 1",
				IsComplete:  "1",
			},
			{
				ID:          "2",
				UserID:      "1",
				Title:       "title 2",
				Description: "description 2",
				IsComplete:  "0",
			},
		}
		assert.Equal(t, want, got)
	})
}

func TestTaskUsecase_ReadTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		taskUsecase, mocks := newMocks(ctrl)
		task := &entity.Task{
			ID:          "1",
			UserID:      "1",
			Title:       "test title",
			Description: "test description",
			IsComplete:  "0",
		}
		taskID := "1"
		mocks.taskMockRespository.EXPECT().SearchTaskByTaskID(taskID).Return(task, nil)
		got, err := taskUsecase.ReadTask(taskID)
		require.NoError(t, err)
		if err != nil {
			t.Errorf("error message: %v", err)
		}

		want := &entity.Task{
			ID:          "1",
			UserID:      "1",
			Title:       "test title",
			Description: "test description",
			IsComplete:  "0",
		}
		assert.Equal(t, want, got)
	})
}
