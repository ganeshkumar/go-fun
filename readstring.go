package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}
