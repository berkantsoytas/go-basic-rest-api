package data

import (
	"encoding/json"

	. "../models"
	. "../utils"
)

func LoadUser() []User {
	bytes, _ := ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterests() []Interest {
	bytes, _ := ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterestMapping() []InterestMapping {
	bytes, _ := ReadFile("../json/interest_mapping.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}