package main

import "fmt"

const (
	coin = iota
	push = iota
)

type State interface {
	Handle() State
}

type TurnstileContext struct {
	action int
}

type TurnstileState struct {
	context TurnstileContext
}

type TurnstileStatemachine struct {
	state TurnstileState
}

func (ts *TurnstileState) Handle() {
	switch ts.context.action {
	case coin:
		fmt.Println("coin")
	case push:
		fmt.Println("push")
	default:
		fmt.Println("this is unexpected")
	}
}

func main() {
	fmt.Printf("Turnstile\n")
}
