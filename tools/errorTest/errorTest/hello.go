package errorTest

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error1) {
	//fmt.Println(quote.Go())
	if name == "" {
		return "", New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
