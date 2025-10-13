package dto

//This file keep the details related to User

type UserDto struct {
	UserID    string `json:"userID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponseDto struct {
	Message string `json:"message"`
	UserID  string `json:"userID"`
}
