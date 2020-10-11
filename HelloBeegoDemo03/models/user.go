package models

type User struct{
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Password string `json:"password"`
}
