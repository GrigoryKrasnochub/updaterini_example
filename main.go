package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Example file")
	fmt.Println("Press any key for update")
	bufio.NewReader(os.Stdin).ReadByte()
	fmt.Println("Success!")
}
