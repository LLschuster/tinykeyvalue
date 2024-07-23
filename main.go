package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Start of tinykeyvalue")
	var port uint = 6969
	if os.Args[1] != "" {
		var (
			err error
			v   int
		)
		v, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid port number")
			os.Exit(1)
		}
		port = uint(v)
	}
	StartServer(port)
}
