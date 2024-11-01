package entity

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
