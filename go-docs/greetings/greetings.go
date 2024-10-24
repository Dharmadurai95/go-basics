package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(greetingMessage(), name)
	return message, nil
}

func greetingMessage() string {
	formate := []string{
		"Hi %v . Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formate[rand.Intn(len(formate))]
}

func Hellos(names []string) (map[string]string, error) {

	message := make(map[string]string)
	for _, eachName := range names {
		msg, err := Hello(eachName)
		if err != nil {
			return nil, err
		}
		message[eachName] = msg
	}
	return message, nil

}
