package models

type Person struct {
	User string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
	Name interface{}
}