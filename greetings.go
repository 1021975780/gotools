package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("名字不能为空")
	}
	//message := fmt.Sprintf("Hello world, my name is %v !", name)
	var message string = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	var resultMap map[string]string = make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		resultMap[name] = message
	}
	return resultMap, nil
}
func randomFormat() string {
	var formats []string = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	rand.Seed(time.Now().Unix())
	return formats[rand.Intn(len(formats))]
}
