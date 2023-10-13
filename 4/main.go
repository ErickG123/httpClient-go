package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	// Usando Contextos no HTTP
	// Criando um contexto vazio
	ctx := context.Background()

	// Criando um contexto de 1 segundos
	// Se algo que utilize o contexto passar de 1 segundo,
	// ele será cancelado
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	// Só é cancelado através do método cancel()
	// ctx, cancel := context.WithCancel(ctx)

	// Rodando a função de cancelamento no final
	// O Contexto será cancelado automaticamento no final ou
	// quando passar do tempo limite
	defer cancel()

	// Criando um Requisição que utiliza de um contexto
	request, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	// http.DefaultClient é o Client padrão do HTTP
	response, err := http.DefaultClient.Do(request)
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
