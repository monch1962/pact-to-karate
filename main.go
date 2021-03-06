package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
    "strings"
	// "bytes"
	"time"
    "io/ioutil"
    "github.com/aymerick/raymond"
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
	// fmt.Printf("%+v\n", pact)

	templateFile := "./templates/" + os.Getenv("TEMPLATE") + ".hbs"
	f, err := os.Open(templateFile)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

    raymond.RegisterHelper("toJSON", func(body interface{}, options *raymond.Options) string {
        jsonBytes, _ := json.Marshal(&body)
        return string(jsonBytes)
    })

    raymond.RegisterHelper("lowerCase", func(s string) string {
        return string(strings.ToLower(s))
	})
	
	raymond.RegisterHelper("isoTimeRFC3339", func() string {
		return string(time.Now().UTC().Format(time.RFC3339))
	})

	raymond.RegisterHelper("envVar", func(envvar string) string {
		return string(os.Getenv(envvar))
	})

    template, err := ioutil.ReadAll(f)
    result, err := raymond.Render(string(template), pact)
    fmt.Println(result)
}
