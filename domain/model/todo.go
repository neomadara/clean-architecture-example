package model

type Todo struct {
	Title     string `json:"title"`
	Desc      string `json:"description"`
	Completed bool   `json:"completed"`
}
