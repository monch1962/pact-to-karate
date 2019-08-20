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

func convertToKarateStub(pact Pact) {
	provider := pact.Provider
	consumer := pact.Consumer
	fmt.Printf("%s %s %s %s\n", "Feature: Provider", provider, "responding to consumer", consumer)
	fmt.Println()
	fmt.Println("Background:")
	fmt.Println("  * configure cors = true")
	fmt.Println()
	for i := range pact.Interactions {
		// fmt.Println(d)
		fmt.Printf("%s %s\n", "#", pact.Interactions[i].Description)
		fmt.Printf("%s%s%s%s%s\n", "Scenario: pathMatches('", pact.Interactions[i].Request.Path, "') && methodIs('", pact.Interactions[i].Request.Method, "')")
		if pact.Interactions[i].Response.Body != nil {
			fmt.Printf("    * def response = %s\n", pact.Interactions[i].Response.Body)
		}
		if pact.Interactions[i].Response.Status != 0 {
			fmt.Printf("    * def responseStatus = %d\n", pact.Interactions[i].Response.Status)
		}
	}
	fmt.Println()
	fmt.Println("# No match found - default scenario is to return a 404")
	fmt.Println("Scenario:")
	fmt.Println("    * def responseStatus = 404")
}

func main() {
	var pact Pact
	err := json.NewDecoder(os.Stdin).Decode(&pact)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", pact)
	convertToKarateStub(pact)
}