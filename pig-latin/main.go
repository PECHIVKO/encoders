package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/PECHIVKO/encoders/pig-latin/encoder"
)

func main() {
	var inputPhrase string

	reader := bufio.NewReader(os.Stdin)
	println("Please enter phrase to convert:")
	inputPhrase, _ = reader.ReadString('\n')
	inputPhrase = strings.TrimSuffix(inputPhrase, "\n")
	println(encoder.Encode(inputPhrase))
}
