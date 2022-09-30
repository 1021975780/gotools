package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("名字不能为空")
	}
	//message := fmt.Sprintf("Hello world, my name is %v !", name)
	var message string = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	var formats []string = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
