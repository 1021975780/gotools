package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("名字不能为空")
	}
	message := fmt.Sprintf("Hello world, my name is %v !", name)
	fmt.Println("v0.0.2")
	return message, nil
}
