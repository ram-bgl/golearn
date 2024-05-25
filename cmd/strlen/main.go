package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Length of string is ", len(text))

	str := `
	What is the quote? These aren't the droids you're looking for.
	Who said it? Obi-Wan Kenobi
	Obi-Wan Kenobi says, "These aren't the droids
	you're looking for."`
	fmt.Println(str)
}
