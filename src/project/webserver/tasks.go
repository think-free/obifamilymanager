package webserver

import (
	"net/http"
)

func (s *WebServer) TasksGetTaskList(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TasksAddTaskForUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TasksValidateUserTask(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TasksCancelValidateUserTask(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TasksGetDayUserTasks(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TasksGetWeekUserTasks(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) TasksGetMonthUserTasks(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}