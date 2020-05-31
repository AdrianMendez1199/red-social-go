package routes

import (
	"encoding/json"
	"net/http"
	"github.com/AdrianMendez1199/red-social-go/db"
	"github.com/AdrianMendez1199/red-social-go/models"
)


func register(w http.ResponseWriter, r *http.Request) {
  var user models.User
  err := json.NewDecoder(r.Body).Decode(&user)

  if err != nil {
	  http.Error(w, "Error in data " + err.Error(), 400) 
	  return
  }

  if len(user.Email) == 0 {
	http.Error(w, "user mail is required", 400) 
	return
  }

  if len(user.Email) < 6 {
	http.Error(w, "password most 6 chartets or more", 400) 
	return
  }

 	_, _, err = db.Register(user)

  if err != nil {
	  http.Error(w, "error creating user", 400)
	  return
  }

  w.WriteHeader(http.StatusCreated)

}