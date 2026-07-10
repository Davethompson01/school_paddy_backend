package handler

import (
	"encoding/json"
	"net/http"
)

func Studenthandler(res http.ResponseWriter, req *http.Request) {

	type users struct {
		Name         string `name:"json"`
		Email        string `email:"josn"`
		Phone_Number string `phone_number:"json"`
		Password     string `password:"json"`
	}

	decode := json.NewDecoder(req.Body)
	user := users{}
	decode.Decode(&user)

	
}
