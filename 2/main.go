package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	// Criando um Buffer
	// O Post só vai aceitar um Buffer, por causa do Reader
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Erick"}`))
	// Realizando um HTTP Post
	response, err := c.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Pega os dados os Body e joga no Stdout
	io.CopyBuffer(os.Stdout, response.Body, nil)
}
