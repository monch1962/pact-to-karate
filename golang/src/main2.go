package main

import (
	// "pact"
	"encoding/json"
	"fmt"
	"log"
	"os"
	// "strings"
	"github.com/alexkappa/mustache"
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

type HeaderString struct {
	Header string `json:"headerString"`
}

type Request struct {
	Method        string      `json:"method"`
	Path          string      `json:"path"`
	Headers       interface{} `json:"headers"`
	// Headers	[]HeaderString `json:"headers"`
	Query         interface{} `json:"query"`
	Body          interface{} `json:"body"`
	MatchingRules interface{} `json:"matchingRules"`
}

type Response struct {
	Status        int32       `json:"status"`
	Headers       interface{} `json:"headers"`
	Body          interface{} `json:"body"`
	Generators    interface{} `json:"generators"`
	MatchingRules interface{} `json:"matchingRules"`
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
	fmt.Printf("%+v\n", pact)
	// pactStr, err := json.MarshalIndent(pact, "", "  ")
	// pactStr2 := map[string]interface{}(pactStr.(interface))
	// fmt.Printf("%s\n", pactStr)

	templateFile := "../templates/" + os.Getenv("TEMPLATE") + ".mustache"
	f, err := os.Open(templateFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t, err := mustache.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	t.Render(os.Stdout, pact)
	// t.Render(os.Stdout, map[string]interface{}{"consumer": {"name": "billy"}})
}
