package webserver

import (
	"net/http"
	"byte"

	rice "github.com/GeertJohan/go.rice"
	"github.com/asdine/storm"
)

// WebServer create the main web server that handle the gui requests
type WebServer struct {
	db *storm.DB
}

// New create the webserver
func New(db *storm.DB, dev bool) *WebServer {

	s := &WebServer{
		db: db,
	}

	// Server the web app and the files in the docker compose tree
	if dev {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./src/project/gui/out/"))))

	} else {
		box := rice.MustFindBox("../gui/out/")
		http.Handle("/", http.FileServer(box.HTTPBox()))
	}

	http.HandleFunc("/test", s.Test)

	http.HandleFunc("/user/add", s.UserAdd)
	http.HandleFunc("/user/delete", s.UserDelete)
	http.HandleFunc("/user/setAvatar", s.UserSetAvatar)
	http.HandleFunc("/user/getAvatar", s.UserGetAvatar)
	http.HandleFunc("/user/setScore", s.UserSetScore)
	http.HandleFunc("/user/getScore", s.UserGetScore)
	http.HandleFunc("/user/getUsers", s.UserGetUsers)

	http.HandleFunc("/tasks/getTaskList", s.TasksGetTaskList)
	http.HandleFunc("/tasks/addTaskForUser", s.TasksAddTaskForUser)
	http.HandleFunc("/tasks/validateUserTask", s.TasksValidateUserTask)
	http.HandleFunc("/tasks/cancelValidateUserTask", s.TasksCancelValidateUserTask)
	http.HandleFunc("/tasks/getDayUserTasks", s.TasksGetDayUserTasks)
	http.HandleFunc("/tasks/getWeekUserTasks", s.TasksGetWeekUserTasks)
	http.HandleFunc("/tasks/getMonthUserTasks", s.TasksGetMonthUserTasks)

	http.HandleFunc("/tasks/timeGetPeriods", s.TimeGetPeriods)
	http.HandleFunc("/tasks/timeInsertPeriod", s.TimeInsertPeriod)
	http.HandleFunc("/tasks/timeDeletePeriod", s.TimeDeletePeriod)

	return s
}

// Run start the web server
func (s *WebServer) Run() error {

	return http.ListenAndServe(":8123", nil)
}


func (s *WebServer) Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
