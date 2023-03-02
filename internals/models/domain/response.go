package domain

type ResponseFromDomain struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_At string `json:"created_at"`
}
