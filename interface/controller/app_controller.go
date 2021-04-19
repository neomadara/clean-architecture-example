package controller

type AppController struct {
	Todo interface{ TodoController }
}
