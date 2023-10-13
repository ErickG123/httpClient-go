package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// O Timeout indica o tempo máximo que a requisição
	// pode durar
	c := http.Client{Timeout: time.Second}
	response, err := c.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
