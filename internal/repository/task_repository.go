package repository

import (
	"context"
	"github.com/abhi00999/task-management/models"
	"github.com/abhi00999/task-management/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	Create(ctx context.Context, task models.Task) (models.Task, error)
	List(ctx context.Context, status string, limit, skip int64) ([]models.Task, error)
	Update(ctx context.Context, id primitive.ObjectID, task models.Task) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type taskRepo struct {
	coll string
}

func NewTaskRepository() TaskRepository {
	return &taskRepo{coll: "tasks"}
}

func (r *taskRepo) Create(ctx context.Context, task models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := db.GetCollection("taskdb", r.coll).InsertOne(ctx, task)
	return task, err
}

func (r *taskRepo) List(ctx context.Context, status string, limit, skip int64) ([]models.Task, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	opts := options.Find().SetLimit(limit).SetSkip(skip)
	cur, err := db.GetCollection("taskdb", r.coll).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	err = cur.All(ctx, &tasks)
	return tasks, err
}

func (r *taskRepo) Update(ctx context.Context, id primitive.ObjectID, task models.Task) error {
	_, err := db.GetCollection("taskdb", r.coll).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": task})
	return err
}

func (r *taskRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := db.GetCollection("taskdb", r.coll).DeleteOne(ctx, bson.M{"_id": id})
	return err
}
