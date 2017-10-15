package main

import (
    "os"
    "github.com/cquon/block-letters/letters"
)

func main() {
	letters.InitializeBlockLetterMap()
	word := os.Args[1]

	printBlockWord(word)
}

func printBlockWord(word string) {
	blockword := letters.NewBlockWordFromString(word)
	blockword.Print()
}