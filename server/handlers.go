package server

import (
	"encoding/json"
	"net/http"
	"restapi/todo"

	"github.com/gorilla/mux"
)

type HttpHandlers struct {
	list *todo.List
}

func NewHttpHandlers(l *todo.List) *HttpHandlers {
	return &HttpHandlers{
		list: l,
	}
}

// json with name and description\
// return json representing created task
// из джсона декоднуть, провалидировать, добавить и вернуть ответ
func (h *HttpHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto TaskDto

	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		errDto := NewErrorDto(err.Error())
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDto.ValidateForCreate(); err != nil {
		WriteError(err, w)
		return
	}

	t := todo.NewTask(taskDto.Name, taskDto.Description)

	if err := h.list.AddTask(*t); err != nil {
		WriteError(err, w)
		return
	}

	b, err := json.MarshalIndent(t, "", "--")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

// name in patern
// вычитать из патерна, ошибку, удалить, вернуть пустое тело
func (h *HttpHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	if err := h.list.DeleteTask(name); err != nil {
		WriteError(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// имя в паттерне + json с completed:true
// вычитать из патерна, из джсона, прописать комплит и анкомплит, вернуть ПОЛНУЮ изменённую задачу
func (h *HttpHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	var taskStatus CTaskDto

	if err := json.NewDecoder(r.Body).Decode(&taskStatus); err != nil {
		e := NewErrorDto(err.Error())
		http.Error(w, e.ToString(), http.StatusBadRequest)
		return
	}

	switch {
	case taskStatus.Completed:
		if err := h.list.CompleteTask(name); err != nil {
			WriteError(err, w)
			return
		}
	default:
		if err := h.list.UnCompleteTask(name); err != nil {
			WriteError(err, w)
			return
		}
	}

	b, err := json.MarshalIndent(taskStatus, "", "--")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *HttpHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.list.GetAllTasks()

	b, err := json.MarshalIndent(tasks, "", "--")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// имя в паттерне
// вычитать из патерна, гетнуть, в ответе саму таску
func (h *HttpHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	task, err := h.list.GetTask(name)
	if err != nil {
		WriteError(err, w)
		return
	}

	b, err := json.MarshalIndent(task, "", "--")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// вызвать аллунктаскс и передать в ответ
func (h *HttpHandlers) HandleGetUncompletedTask(w http.ResponseWriter, r *http.Request) {
	UTasks := h.list.GetUncTasks()

	b, err := json.MarshalIndent(UTasks, "", "--")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// вычитать имя таски из патерна, вычетать переданное поле из джсона, проверять если что то пустое то вызывать это, записать ответ
func (h *HttpHandlers) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	var TaskDto TaskDto

	name := mux.Vars(r)["name"]

	TmpTask, err := h.list.GetTask(name)

	if err != nil {
		WriteError(err, w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&TaskDto); err != nil {
		e := NewErrorDto(err.Error())
		http.Error(w, e.ToString(), http.StatusBadRequest)
		return
	}

	if err := TaskDto.ValdateForUpdate(); err != nil {
		WriteError(err, w)
		return
	}

	newName := TmpTask.Name
	newDescription := TmpTask.Description

	if TaskDto.Name != "" {
		newName = TaskDto.Name
	}
	if TaskDto.Description != "" {
		newDescription = TaskDto.Description
	}

	switch {
	case newName != TmpTask.Name && newDescription != TmpTask.Description:
		if err := h.list.UpdateFullTask(name, newName, newDescription); err != nil {
			WriteError(err, w)
			return
		}
	case newName != TmpTask.Name && newDescription == TmpTask.Description:
		if err := h.list.UpdateTaskName(name, newName); err != nil {
			WriteError(err, w)
			return
		}
	case newName == TmpTask.Name && newDescription != TmpTask.Description:
		if err := h.list.UpdateTaskDescription(name, newDescription); err != nil {
			WriteError(err, w)
			return
		}
	}

	responseTask, err := h.list.GetTask(newName)
	if err != nil {
		WriteError(err, w)
		return
	}

	b, errJ := json.MarshalIndent(responseTask, "", "--")
	if errJ != nil {
		panic(errJ)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
