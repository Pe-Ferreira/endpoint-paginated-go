package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Broker struct {
	Bairro                 string `json:"bairro"`
	Cep                    string `json:"cep"`
	Cnpj                   string `json:"cnpj"`
	CodigoCvm              string `json:"codigo_cvm"`
	Complemento            string `json:"complemento"`
	DataInicioSituacao     string `json:"data_inicio_situacao"`
	DataPatrimonioLiquido  string `json:"data_patrimonio_liquido"`
	DataRegistro           string `json:"data_registro"`
	Email                  string `json:"email"`
	Logradouro             string `json:"logradouro"`
	Municipio              string `json:"municipio"`
	NomeSocial             string `json:"nome_social"`
	NomeComercial          string `json:"nome_comercial"`
	Pais                   string `json:"pais"`
	Status                 string `json:"status"`
	Telefone               string `json:"telefone"`
	Type                   string `json:"type"`
	Uf                     string `json:"uf"`
	ValorPatrimonioLiquido string `json:"valor_patrimonio_liquido"`
}

type BrokerDTO struct {
	NomeComercial          string
	ValorPatrimonioLiquido string
	Uf                     string
}

type BrokersDTO []BrokerDTO

func main() {
	http.HandleFunc("/", MainHandler)
	http.ListenAndServe(":8080", nil)
}

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Wrong route, young boy"))
		return
	}

	var brokers, error = getBrokers()
	if error != nil {
		panic(error)
	}

	var brokersDTO = prepareBrokersDTO(brokers)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	tmplt := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmplt.Execute(writer, brokersDTO)
	if err != nil {
		panic(err)
	}
}

func getBrokers() ([]Broker, error) {
	response, error := http.Get("https://brasilapi.com.br/api/cvm/corretoras/v1")
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()

	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var brokers []Broker
	error = json.Unmarshal(responseBody, &brokers)
	if error != nil {
		return nil, error
	}

	return brokers, nil
}

func prepareBrokersDTO(brokers []Broker) []BrokerDTO {
	var brokersDTO []BrokerDTO
	for _, value := range brokers {
		brokersDTO = append(brokersDTO, BrokerDTO{
			NomeComercial:          value.NomeComercial,
			ValorPatrimonioLiquido: value.ValorPatrimonioLiquido,
			Uf:                     value.Uf})
	}
	return brokersDTO
}
