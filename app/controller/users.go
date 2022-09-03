package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"catchMeStore/app/models"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get a list of all users
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	// return list pf users in json format
	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
