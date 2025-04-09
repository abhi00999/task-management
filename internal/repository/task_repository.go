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
	collection := db.GetCollection("taskdb", r.coll)
	_, err := collection.InsertOne(ctx, task)
	return task, err
}

func (r *taskRepo) List(ctx context.Context, status string, limit, skip int64) ([]models.Task, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}

	opts := options.Find().SetLimit(limit).SetSkip(skip)
	collection := db.GetCollection("taskdb", r.coll)
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepo) Update(ctx context.Context, id primitive.ObjectID, task models.Task) error {
	collection := db.GetCollection("taskdb", r.coll)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": task})
	return err
}

func (r *taskRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	collection := db.GetCollection("taskdb", r.coll)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
