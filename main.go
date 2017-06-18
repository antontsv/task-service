package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Task holds a definition for a generic task
type Task struct {
	Name        string
	Description string
}

// TaskService provides operations to collect and show tasks
type TaskService interface {
	// Add creates a task given name and description
	Add(Task) error
	Count() int
	// Removes a task give a name
	Remove(string) error
	Show(int) ([]Task, error)
}

type taskService struct{}

var errNotImplemented = errors.New("Not implemented yet")

func (taskService) Add(t Task) error {
	return errNotImplemented
}

func (taskService) Count() int {
	return 0
}

func (taskService) Remove(name string) error {
	return errNotImplemented
}

func (taskService) Show(maxSize int) ([]Task, error) {
	return make([]Task, 0), nil
}

func main() {
	ctx := context.Background()
	svc := taskService{}

	countHandler := httptransport.NewServer(
		ctx,
		makeCountEndpoint(svc),
		func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
		encodeResponse,
	)

	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func makeCountEndpoint(svc TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v := svc.Count()
		return countResponse{v}, nil
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type countResponse struct {
	Size int `json:"size"`
}
