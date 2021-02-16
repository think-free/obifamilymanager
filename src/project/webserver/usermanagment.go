package webserver

import (
	"net/http"
	"encoding/json"
)

type User struct {

	ID   int `storm:"id,increment"`
	Name string `json:"name"`
}

type UserScore struct {

	UserID int `storm:"id,increment"`
	Score  int `json:"score"`
}

func (s *WebServer) UserAdd(w http.ResponseWriter, r *http.Request) {

	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Println(err)
		w.Write([]byte("{\"type\" : \"error\", \"msg\":" + err.Error() + "}"))
	} else {

		err := s.db.Save(&user)
		if err != nil {
			log.Println(err)
			w.Write([]byte("{\"type\" : \"error\", \"msg\":" + err.Error() + "}"))
		} else {

			js, _ := json.Marshal(user)
			log.Println(js)

			js = append([]byte("{\"type\" : \"json\", \"msg\":"), js...)
			js = append(js, []byte("}")...)
			w.Write(js)
		}
	}
}

func (s *WebServer) UserDelete(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) UserSetAvatar(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) UserGetAvatar(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) UserSetScore(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) UserGetScore(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}

func (s *WebServer) UserGetUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}