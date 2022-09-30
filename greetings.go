package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello world, my name is %v !", name)
	return message
}
