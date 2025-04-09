package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/abhi00999/task-management/models"
	"github.com/abhi00999/task-management/internal/repository"
)

type TaskService interface {
	Create(ctx context.Context, task models.Task) (models.Task, error)
	List(ctx context.Context, status string, limit, skip int64) ([]models.Task, error)
	Update(ctx context.Context, id primitive.ObjectID, task models.Task) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService() TaskService {
	return &taskService{repo: repository.NewTaskRepository()}
}

func (s *taskService) Create(ctx context.Context, task models.Task) (models.Task, error) {
	return s.repo.Create(ctx, task)
}

func (s *taskService) List(ctx context.Context, status string, limit, skip int64) ([]models.Task, error) {
	return s.repo.List(ctx, status, limit, skip)
}

func (s *taskService) Update(ctx context.Context, id primitive.ObjectID, task models.Task) error {
	return s.repo.Update(ctx, id, task)
}

func (s *taskService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
