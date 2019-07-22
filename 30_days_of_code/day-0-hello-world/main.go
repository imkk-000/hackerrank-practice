package main

import (
	"bufio"
	"fmt"
	"os"
)

func ShowHello(message string) string {
	return fmt.Sprintf("Hello, World.\n%s", message)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()
	fmt.Println(ShowHello(message))
}
