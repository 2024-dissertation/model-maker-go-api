package mocks

import (
	models "github.com/Soup666/diss-api/model"
	"github.com/stretchr/testify/mock"
)

// MockTaskService is a mock implementation of TaskService.
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(task *models.Task) error {
	args := m.Called(task)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (m *MockTaskService) GetTask(taskID uint) (*models.Task, error) {
	args := m.Called(taskID)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskService) GetTasks(userID uint) ([]*models.Task, error) {
	args := m.Called(userID)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskService) UpdateTask(task *models.Task) error {
	args := m.Called(task)
	if args.Get(0) != nil {
		return args.Error(1)
	}
	return nil
}

func (m *MockTaskService) DeleteTask(taskID *models.Task) error {
	args := m.Called(taskID)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (m *MockTaskService) SaveTask(task *models.Task) error {
	args := m.Called(task)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (m *MockTaskService) FailTask(task *models.Task) error {
	args := m.Called(task)
	if args.Get(0) != nil {

		return args.Get(0).(error)
	}
	return nil
}

func (m *MockTaskService) RunPhotogrammetryProcess(task *models.Task) error {
	args := m.Called(task)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (m *MockTaskService) UpdateMeta(task *models.Task, key string, value interface{}) error {
	args := m.Called(task, key, value)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (m *MockTaskService) FullyLoadTask(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Task), args.Error(1)
	}
	return nil, nil
}

func (m *MockTaskService) SendMessage(taskID uint, message string, sender string) (*models.ChatMessage, error) {
	args := m.Called(taskID, message, sender)
	if args.Get(0) != nil {
		return args.Get(0).(*models.ChatMessage), args.Error(1)
	}
	return nil, args.Error(1)
}
