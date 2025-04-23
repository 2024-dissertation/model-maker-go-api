package services_test

import (
	"errors"
	"testing"

	"github.com/Soup666/diss-api/mocks"
	models "github.com/Soup666/diss-api/model"
	"github.com/Soup666/diss-api/services"
	"github.com/stretchr/testify/assert"
)

func TestTaskService(t *testing.T) {
	mockTaskRepository := new(mocks.MockTaskRepository)
	mockChatRepository := new(mocks.MockChatRepository)

	mockAppFileService := new(mocks.MockAppFileService)

	taskService := services.NewTaskService(mockTaskRepository, mockAppFileService, mockChatRepository)

	t.Run("CreateTask", func(t *testing.T) {

		var userId = uint(1)

		task := &models.Task{
			Id:          1,
			Title:       "Test Task",
			Description: "This is a test task",
			UserId:      &userId,
		}

		mockTaskRepository.On("CreateTask", task).Return(nil)
		mockTaskRepository.On("CreateTask", nil).Return(errors.New("error"))
		mockTaskRepository.On("CreateTask", &models.Task{}).Return(errors.New("error"))
		mockTaskRepository.On("CreateTask", &models.Task{Id: 1}).Return(nil)

		err := taskService.CreateTask(task)

		mockTaskRepository.AssertCalled(t, "CreateTask", task)
		assert.NoError(t, err)
		assert.NotNil(t, task)
		assert.NotNil(t, task.Id)
		assert.Equal(t, task.Title, "Test Task")
	})

	t.Run("GetTask", func(t *testing.T) {
		task := &models.Task{
			Id:          1,
			Title:       "Test Task",
			Description: "This is a test task",
		}

		mockTaskRepository.On("GetTaskByID", uint(1)).Return(task, nil)
		mockTaskRepository.On("GetTaskByID", uint(2)).Return(nil, errors.New("error"))

		fetchedTask, err := taskService.GetTask(1)

		mockTaskRepository.AssertCalled(t, "GetTaskByID", uint(1))
		assert.NoError(t, err)
		assert.NotNil(t, fetchedTask)
		assert.Equal(t, fetchedTask.Title, "Test Task")
	})

	t.Run("GetTaskByID with non-existent ID", func(t *testing.T) {
		mockTaskRepository.On("GetTaskByID", uint(2)).Return(nil, errors.New("error"))

		fetchedTask, err := taskService.GetTask(2)

		mockTaskRepository.AssertCalled(t, "GetTaskByID", uint(2))

		assert.Error(t, err)
		assert.Nil(t, fetchedTask)
		assert.Equal(t, err.Error(), "error")
	})

	t.Run("GetTasks", func(t *testing.T) {

		var userId = uint(1)

		tasks := []*models.Task{
			{
				Id:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				UserId:      &userId,
			},
			{
				Id:          2,
				Title:       "Test Task 2",
				Description: "This is a test task 2",
				UserId:      &userId,
			},
		}

		mockTaskRepository.On("GetTasksByUser", userId).Return(tasks, nil)
		mockTaskRepository.On("GetTasksByUser", uint(2)).Return(nil, errors.New("error"))

		fetchedTasks, err := taskService.GetTasks(userId)

		mockTaskRepository.AssertCalled(t, "GetTasksByUser", userId)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedTasks)
		assert.Equal(t, len(fetchedTasks), 2)
	})
}
