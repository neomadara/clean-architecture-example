package model

type Todo struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}
