package greeting

import (
	"math/rand"
)

func RandomGreet() string {
	messages := []string{
		"Hello",
		"Hi",
		"Howdy",
		"Yo you back",
	}
	max := len(messages) - 1
	min := 0
	randomNumber := rand.Intn(max-min) + min
	return messages[randomNumber]
}
