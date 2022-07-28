package main

import (
	"fmt"
	"strings"

	"github.com/Neo-Desktop/emojipasta"
)

func main() {
	input := "This library allows you to create emojipastas like this one"
	output := emojipasta.Generate(strings.Split(input, " "), 3, false)
	fmt.Println(output)
}
