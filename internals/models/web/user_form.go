package web

import "time"

type UsersRegister struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersEdit struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Update_At time.Time `json:"updated_at"`
}
