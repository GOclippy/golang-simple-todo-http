package server

import (
	"encoding/json"
	"net/http"
	"restapi/todo"
	"time"
)

type TaskDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CTaskDto struct {
	Completed bool `json:"completed"`
}

type ErrorDto struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewErrorDto(err string) *ErrorDto {
	return &ErrorDto{
		Message: err,
		Time:    time.Now(),
	}
}

func (e *ErrorDto) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	s := string(b)

	return s
}

func (t *TaskDto) ValidateForCreate() error {
	switch {
	case t.Name == "":
		return todo.ErrDescriptionIsEmpty
	case t.Description == "":
		return todo.ErrNameIsEmpty
	}
	return nil
}

func (t *TaskDto) ValdateForUpdate() error {
	if t.Name == "" && t.Description == "" {
		return todo.ErrNameAndDescriptionIsEmpty
	}

	return nil
}

func WriteError(err error, w http.ResponseWriter) {
	e := NewErrorDto(err.Error())

	switch {
	case e.Message == "ErrTaskAlreadyExist":
		http.Error(w, e.ToString(), http.StatusConflict)
	case e.Message == "ErrTaskNotFound":
		http.Error(w, e.ToString(), http.StatusNotFound)
	case e.Message == "ErrNameIsEmpty":
		http.Error(w, e.ToString(), http.StatusBadRequest)
	case e.Message == "ErrDescriptionIsEmpty":
		http.Error(w, e.ToString(), http.StatusBadRequest)
	case e.Message == "ErrNameAndDescriptionIsEmpty":
		http.Error(w, e.ToString(), http.StatusBadRequest)
	default:
		http.Error(w, e.ToString(), http.StatusInternalServerError)
	}
}
