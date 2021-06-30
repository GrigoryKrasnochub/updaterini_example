package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Example file dev branch")
	fmt.Println("Press any key for update")
	bufio.NewReader(os.Stdin).ReadByte()
	fmt.Println("Success!")
}
