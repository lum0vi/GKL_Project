package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:1323/get")
	if err != nil {
		fmt.Errorf("error to connect: %w", err)
	}
	defer response.Body.Close()
	fmt.Println("yes")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("error to read body %w", err)
	}
	fmt.Println(string(body))
}
