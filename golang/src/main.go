package main

import (
	//"pact"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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

func parseReqBodyJSON(content interface{}) []string {
	var elements []string
	json := content.(map[string]interface{})
	for k, v := range json {
		element := fmt.Sprintf("request[\"%v\"] == \"%v\"", k, v)
		elements = append(elements, element)
	}
	return elements
}

func parseReqHeaders(content interface{}) []string {
	var elements []string
	json := content.(map[string]interface{})
	for k, v := range json {
		element := fmt.Sprintf("headerContains('%v', '%v')", k, v)
		elements = append(elements, element)
	}
	return elements
}

func convertToKarateStub(pact Pact) {
	provider := pact.Provider
	consumer := pact.Consumer
	fmt.Printf("%s%s%s%s%s\n", "Feature: Provider '", provider.Name, "' responding to consumer '", consumer.Name, "'")
	fmt.Println()
	fmt.Println("Background:")
	fmt.Println("  * configure cors = true")

	for _, d := range pact.Interactions {
		// fmt.Println(d)
		fmt.Println()
		fmt.Printf("%s %s\n", "#", d.Description)
		fmt.Printf("%s%s%s%s%s", "Scenario: pathMatches('", d.Request.Path, "') && methodIs('", d.Request.Method, "')")

		// fmt.Printf("Request.Headers: %s\n", d.Request.Headers)
		reqHeaders := parseReqHeaders(d.Request.Headers)
		if len(reqHeaders) != 0 {
			fmt.Printf("%s", " && ")
			reqH := strings.Join(reqHeaders, " && ")
			fmt.Printf("%s", reqH)
		}

		reqBody := parseReqBodyJSON(d.Request.Body)
		if len(reqBody) != 0 {
			fmt.Printf("%s", " && ")
			reqB := strings.Join(reqBody, " && ")
			fmt.Printf("%s\n", reqB)
		}
		fmt.Println()

		// fmt.Printf("Response.Headers: %s\n", d.Response.Headers)

		// fmt.Printf("Response.Body: %s\n", d.Request.Body)
		if d.Response.Body != nil {
			m, err := json.Marshal(d.Response.Body)
			if err == nil {
				fmt.Printf("    * def response = %s\n", m)
			} else {
				fmt.Printf("    * def response = %s\n", d.Response.Body)
			}
		}
		if d.Response.Status != 0 {
			fmt.Printf("    * def responseStatus = %d\n", d.Response.Status)
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
	// fmt.Printf("%v\n", pact)
	convertToKarateStub(pact)
}
