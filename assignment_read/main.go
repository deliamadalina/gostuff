package main

import (
	"fmt"
	"io"
	"os"
)

// type logWriter struct{}

func main() {
	arg := os.Args[1]

	//func OpenFile(name string, flag int, perm FileMode) (*File, error)

	f, err := os.OpenFile(arg, os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	// lw := logWriter{}
	// io.Copy(lw, f)
	io.Copy(os.Stdout, f)

}

// func (logWriter) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	return len(bs), nil
// }
