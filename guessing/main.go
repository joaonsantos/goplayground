package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func gameLoop(secret int) {
	for {
		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Fprintln(os.Stderr, "please input a number")
			os.Exit(1)
		}

		switch {
		case guess > secret:
			fmt.Println("Number is too big")
		case guess < secret:
			fmt.Println("Number is too small")
		case guess == secret:
			fmt.Println("You got it")
			return
		}
	}
}

func generateSecret(max int) int {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	return r.Intn(max + 1)
}

func main() {
	fmt.Println("Guess the number!")

	secret := generateSecret(25)
	gameLoop(secret)
}
