package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	port := 3000
	path := "sudoku.db"
	if len(os.Args) > 1 {
		newPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Port must be an integer")
			os.Exit(1)
		}
		port = newPort
	}
	if len(os.Args) > 2 {
		path = os.Args[2]
		for i := 3; i < len(os.Args); i++ {
			path += " " + os.Args[i]
		}
	}
	Sudoku(port, path)
}