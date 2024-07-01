package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

	}
}
