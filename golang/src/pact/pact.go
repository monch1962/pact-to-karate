package pact

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
	Headers []Header    `json:"headers"`
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
	Interactions []Interaction `json:"interaction"`
	Metadata     Metadata      `json:"metadata"`
}
