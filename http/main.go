package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// not the best way:
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// correct way:
	// io.Copy(os.Stdout, resp.Body)

	lw := logWritter{}
	io.Copy(lw, resp.Body)

}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this manu bytes: ", len(bs))
	return len(bs), nil
}
