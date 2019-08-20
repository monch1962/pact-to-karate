package main

import (
	//"pact"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Consumer struct {
	Name string `json:"name"`
}

type Provider struct {
	Name string `json:"name"`
}

type PactVersion struct {
	PactVersion string `json:"pactVersion"`
}

type PactSpecification struct {
	PactVersion string `json:"pactVersion"`
}

type Metadata struct {
	PactSpecification PactSpecification `json:"pactSpecification"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Request struct {
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Headers interface{} `json:"headers"`
	Body    interface{} `json:"body"`
}

type Response struct {
	Status  int         `json:"status"`
	Headers []Header    `json:"headers"`
	Body    interface{} `json:"body"`
}

type Interaction struct {
	Description   string      `json:"description"`
	ProviderState interface{} `json:"providerState"`
	Request       Request     `json:"request"`
	Response      Response    `json:"response"`
}

type Pact struct {
	Consumer     Consumer      `json:"consumer"`
	Provider     Provider      `json:"provider"`
	Interactions []Interaction `json:"interactions"`
	Metadata     Metadata      `json:"metadata"`
}

func main() {
	var pact Pact
	err := json.NewDecoder(os.Stdin).Decode(&pact)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", pact)
}
