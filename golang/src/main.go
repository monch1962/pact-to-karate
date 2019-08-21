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
	Headers interface{}    `json:"headers"`
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

func convertToKarateTests(pact Pact) string {
	var text strings.Builder
	provider := pact.Provider
	consumer := pact.Consumer
	fmt.Fprintf(&text, "%s%s%s%s%s\n", "Feature: Consumer '", consumer.Name, "' sending requests to provider '", provider.Name, "'")
	for _, d := range pact.Interactions {
		fmt.Fprintln(&text)
		fmt.Fprintf(&text, "  Scenario: %s\n", d.Description)
		fmt.Fprintf(&text, "    Given URL '%s'\n", d.Request.Path)
		if d.Request.Body != nil {
			jsonString, err := json.Marshal(d.Request.Body)
			if err == nil {
				fmt.Fprintf(&text, "    And request %s\n", jsonString)
			}
		}
		fmt.Fprintf(&text, "    When method %s\n", d.Request.Method)
		fmt.Fprintf(&text, "    Then status %d\n", d.Response.Status)
		if d.Response.Body != nil {
			jsonString, err := json.Marshal(d.Response.Body)
			if err == nil {
				fmt.Fprintf(&text, "    And match response == %s\n", jsonString)
			}
		}
	}
	return text.String()
}

func convertToKarateStub(pact Pact) string {
	var text strings.Builder
	provider := pact.Provider
	consumer := pact.Consumer
	fmt.Fprintf(&text, "%s%s%s%s%s\n", "Feature: Provider '", provider.Name, "' responding to consumer '", consumer.Name, "'")
	fmt.Fprintln(&text)
	fmt.Fprintln(&text, "Background:")
	fmt.Fprintln(&text, "  * configure cors = true")

	for _, d := range pact.Interactions {
		// fmt.Println(d)
		fmt.Fprintln(&text)
		fmt.Fprintf(&text, "%s %s\n", "#", d.Description)
		fmt.Fprintf(&text, "%s%s%s%s%s", "Scenario: pathMatches('", d.Request.Path, "') && methodIs('", d.Request.Method, "')")

		// fmt.Printf("Request.Headers: %s\n", d.Request.Headers)
		reqHeaders := parseReqHeaders(d.Request.Headers)
		if len(reqHeaders) != 0 {
			fmt.Fprintf(&text, "%s", " && ")
			reqH := strings.Join(reqHeaders, " && ")
			fmt.Fprintf(&text, "%s", reqH)
		}

		reqBody := parseReqBodyJSON(d.Request.Body)
		if len(reqBody) != 0 {
			fmt.Fprintf(&text, "%s", " && ")
			reqB := strings.Join(reqBody, " && ")
			fmt.Fprintf(&text, "%s\n", reqB)
		}
		fmt.Fprintln(&text)

		// fmt.Printf("Response.Headers: %s\n", d.Response.Headers)

		// fmt.Printf("Response.Body: %s\n", d.Request.Body)
		if d.Response.Body != nil {
			m, err := json.Marshal(d.Response.Body)
			if err == nil {
				fmt.Fprintf(&text, "    * def response = %s\n", m)
			} else {
				fmt.Fprintf(&text, "    * def response = %s\n", d.Response.Body)
			}
		}
        if d.Response.Headers != nil {
            jsonH,_ := json.Marshal(d.Response.Headers)
            fmt.Fprintf(&text, "    * def responseHeaders = %v\n", string(jsonH))
        }

		if d.Response.Status != 0 {
			fmt.Fprintf(&text, "    * def responseStatus = %d\n", d.Response.Status)
		}
	}
	fmt.Fprintln(&text)
	fmt.Fprintln(&text, "# No match found - default scenario is to return a 404")
	fmt.Fprintln(&text, "Scenario:")
	fmt.Fprintln(&text, "    * def responseStatus = 404")
	return text.String()
}

func main() {
	var pact Pact
	err := json.NewDecoder(os.Stdin).Decode(&pact)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", pact)
	fmt.Printf(convertToKarateStub(pact))
	fmt.Println("==================================")
	fmt.Printf(convertToKarateTests(pact))
}
