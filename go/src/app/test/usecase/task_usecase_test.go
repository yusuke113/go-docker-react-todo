package test

import (
	"app/model"
	"app/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllTasks(t *testing.T) {
	mockRepo := &mocks.TaskRepository{}
	tasks := []model.Task{
		{
			ID:        1,
			Title:     "Task 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserId:    1,
		},
		{
			ID:        2,
			Title:     "Task 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserId:    1,
		},
	}
	mockRepo.On("GetAllTasks", mock.AnythingOfType("*[]model.Task"), uint(1)).Return(nil).Run(func(args mock.Arguments) {
		tasksPtr := args.Get(0).(*[]model.Task)
		*tasksPtr = tasks
	})
	taskUseCase := usecase.NewTaskUseCase(mockRepo)
	resTasks, err := taskUseCase.GetAllTasks(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(resTasks))
	assert.Equal(t, tasks[0].ID, resTasks[0].ID)
	assert.Equal(t, tasks[0].Title, resTasks[0].Title)
	assert.Equal(t, tasks[0].CreatedAt, resTasks[0].CreatedAt)
	assert.Equal(t, tasks[0].UpdatedAt, resTasks[0].UpdatedAt)
	assert.Equal(t, tasks[1].ID, resTasks[1].ID)
	assert.Equal(t, tasks[1].Title, resTasks[1].Title)
	assert.Equal(t, tasks[1].CreatedAt, resTasks[1].CreatedAt)
	assert.Equal(t, tasks[1].UpdatedAt, resTasks[1].UpdatedAt)
	mockRepo.AssertExpectations(t)
}
