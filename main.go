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
	Add(Task) error
	Count() int
	Remove(Task) error
	Show(int) ([]Task, error)
}

type taskService struct{}

var errNotImplemented = errors.New("Not implemented yet")

var tasks = make([]Task, 0)

func (taskService) Add(t Task) error {
	return errNotImplemented
}

func (taskService) Count() int {
	return 0
}

func (taskService) Remove(t Task) error {
	return errNotImplemented
}

func (taskService) Show(maxSize int) ([]Task, error) {
	if cap(tasks) < maxSize {
		maxSize = cap(tasks)
	}
	return tasks[0:maxSize], nil
}

func main() {
	http.Handle("/count", httpCountHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func httpCountHandler() *httptransport.Server {
	return httptransport.NewServer(
		context.Background(),
		makeCountEndpoint(taskService{}),
		func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
		encodeResponse,
	)
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
