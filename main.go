package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type DadosCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {

		var URLcep = "https://viacep.com.br/ws/" + cep + "/json/"

		//request, err := http.Get(url)
		request, err := http.Get(URLcep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro no request: %v\n", err)
		}

		defer request.Body.Close()

		response, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o request: %v\n", err)
		}

		var dados DadosCEP
		err = json.Unmarshal(response, &dados)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro no unmarshal: %v\n", err)
		}

		//fmt.Println(dados.Logradouro)
		fmt.Println("CEP: ", dados.Cep)
		fmt.Println("Logradouro: ", dados.Logradouro)
		fmt.Println("Complemento: ", dados.Complemento)
		fmt.Println("Bairro: ", dados.Bairro)
		fmt.Println("Localidade: ", dados.Localidade)
		fmt.Println("UF: ", dados.Uf)
		fmt.Println("IBGE: ", dados.Ibge)
		fmt.Println("GIA: ", dados.Gia)
		fmt.Println("DDD: ", dados.Ddd)
		fmt.Println("SIAFI: ", dados.Siafi)

		//Grava os dados retornados em um arquivo de texto:
		arquivo, err := os.Create("DadosDoCep_" + dados.Cep + ".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo "+arquivo.Name()+" dos dados do cep: %v", err)
		}

		defer arquivo.Close()
		_, err = arquivo.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
			dados.Cep, dados.Logradouro, dados.Complemento, dados.Bairro, dados.Localidade, dados.Uf, dados.Ibge, dados.Gia, dados.Ddd, dados.Siafi))

		fmt.Println("Arquivo gerado com sucesso!")
	}
}
