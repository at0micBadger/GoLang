package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("STOP!")
	fmt.Println("Who would cross the Bridge of Death must answer me these questions three,")
	fmt.Println("ere the other side he see.\n")

	reader := bufio.NewReader(os.Stdin)

	// First question - get the knight's name
	fmt.Println("Question 1: What is your name?")
	fmt.Print("> ")
	knightName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	knightName = strings.TrimSpace(knightName)
	fmt.Printf("Correct, Knight %s!\n\n", knightName)

	// Question 2 - The quest (must be "the holy grail")
	fmt.Println("Question 2: What is your quest?")
	fmt.Print("> ")
	quest, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	quest = strings.TrimSpace(strings.ToLower(quest))

	if quest != "the holy grail" {
		fmt.Printf("\nAAAAAAAAAGH! Knight %s!\n", knightName)
		fmt.Println("*You are cast into the Gorge of Eternal Peril*")
		fmt.Println("\nYou have failed the quest!")
		os.Exit(1)
	}
	fmt.Printf("Correct, Knight %s!\n\n", knightName)

	// Question 3 - Favorite color (any answer accepted if they're sure)
	fmt.Println("Question 3: What is your favorite color?")
	fmt.Print("> ")
	color, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	color = strings.TrimSpace(color)

	if color == "" {
		fmt.Printf("\nAAAAAAAAAGH! Knight %s!\n", knightName)
		fmt.Println("*You are cast into the Gorge of Eternal Peril*")
		fmt.Println("\nYou have failed the quest!")
		os.Exit(1)
	}

	fmt.Printf("%s! Correct, Knight %s!\n\n", strings.Title(color), knightName)

	fmt.Printf("Right. Off you go, Knight %s!\n", knightName)
	fmt.Println("\n🏆 You have successfully crossed the Bridge of Death! 🏆")
	fmt.Println("The Holy Grail awaits!")
}
