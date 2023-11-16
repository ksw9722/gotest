package gotest

import (
	"bufio"
	"fmt"
	"os"
)

func EchoInput() {
	fmt.Print("Input Your String : ")
	io := bufio.NewReader(os.Stdin)
	input, err := io.ReadString('\n')
	if err != nil {
		fmt.Println("[-] Please retry input")
		return
	}

	fmt.Println(input)
}
