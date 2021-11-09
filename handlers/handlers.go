package handlers

import (
	"encoding/json"
	"net/http"

	. "../data"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request){
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı listesii", URI: "/users"}

	// Data
	users := LoadUser()
	interests := LoadInterests()
	interestMapping := LoadInterestMapping()

	// işlem

	var newUser []User
	for _, user := range users {
		for _, interestMapping := range interestMapping {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUser = append(newUser, users...)
	}
	viewModel :=  UserViewModel{Page: page, Users: newUser}
	veri, _ := json.Marshal(viewModel)
	w.Write([]byte(veri))
}