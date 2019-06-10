package controllers

import (
	"encoding/json"
	"github.com/melardev/GoGormApiCrudPagination/dtos"
	"github.com/melardev/GoGormApiCrudPagination/models"
	"github.com/melardev/GoGormApiCrudPagination/services"
	"net/http"
	"strings"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPagingParams(r)
	todos, totalTodoCount := services.FetchTodos(page, pageSize)
	SendAsJson(w, http.StatusOK, dtos.CreateTodoPagedResponse(r, todos, page, pageSize, totalTodoCount))
}

func GetAllPendingTodos(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPagingParams(r)
	todos, totalTodoCount := services.FetchPendingTodos(page, pageSize, false)
	SendAsJson(w, http.StatusOK, dtos.CreateTodoPagedResponse(r, todos, page, pageSize, totalTodoCount))
}

func GetAllCompletedTodos(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPagingParams(r)
	todos, totalTodoCount := services.FetchPendingTodos(page, pageSize, true)
	SendAsJson(w, http.StatusOK, dtos.CreateTodoPagedResponse(r, todos, page, pageSize, totalTodoCount))
}

func GetTodoById(id int, w http.ResponseWriter, r *http.Request) {
	todo, err := services.FetchById(uint(id))
	if err != nil {
		SendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not find Todo"))
		return
	}

	// Just to prove that sendAsJson2 also works
	sendAsJson2(w, http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil || strings.TrimSpace(todo.Title) == "" ||
		strings.TrimSpace(todo.Description) == "" {
		if err != nil {
			SendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage(err.Error()))
		} else {
			SendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must fill the title and description"))
		}
		return
	}
	defer r.Body.Close()

	todo, err := services.CreateTodo(todo.Title, todo.Description, todo.Completed)
	if err != nil {
		SendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	SendAsJson(w, http.StatusCreated, dtos.CreateTodoCreatedDto(&todo))
}

func UpdateTodo(id int, w http.ResponseWriter, r *http.Request) {
	var todoInput models.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todoInput); err != nil || strings.TrimSpace(todoInput.Title) == "" ||
		strings.TrimSpace(todoInput.Description) == "" {
		if err != nil {
			SendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage(err.Error()))
		} else {
			SendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must fill the title and description"))
		}
		return
	}

	defer r.Body.Close()

	todo, err := services.UpdateTodo(uint(id), todoInput.Title, todoInput.Description, todoInput.Completed)
	if err != nil {
		SendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	SendAsJson(w, http.StatusOK, dtos.CreateTodoUpdatedDto(&todo))
}

func DeleteTodo(id int, w http.ResponseWriter, r *http.Request) {
	todo, err := services.FetchById(uint(id))
	if err != nil {
		SendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("todo not found"))
		return
	}

	err = services.DeleteTodo(&todo)

	if err != nil {
		SendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not delete Todo"))
		return
	}

	SendAsJson(w, http.StatusNoContent, dtos.CreateSuccessWithMessageDto("Todo deleted successfully"))
}

func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	services.DeleteAllTodos()
	SendAsJson(w, http.StatusNoContent, dtos.CreateSuccessWithMessageDto("All Todos deleted successfully"))
}
