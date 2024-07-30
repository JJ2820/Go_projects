package main

import (
	"fmt"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `jaon: "What is your Age`
}

type MyResponc struct {
	Message string `json: "Answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponc, error) {
	return MyResponc{Message: fmt.Sprintf("%s is %d years ols", event.Name, event.Age)}, nil
}

func main() {
	lamba.Start(HandleLambdaEvent)
}
