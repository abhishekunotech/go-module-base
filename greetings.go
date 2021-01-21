package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("%v , Hello and Welcome to my base module.",name)
	return message
}