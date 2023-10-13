package main

import (
	"io"
	"net/http"
)

func main() {
	// Criando o Cliente HTTP
	c := http.Client{}

	// Fazendo a Requisição HTTP GET
	request, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	// Falando que aceitamos JSON
	request.Header.Set("Accept", "application/json")

	// Juntando o HTTP Client com a Request
	response, err := c.Do(request)
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
