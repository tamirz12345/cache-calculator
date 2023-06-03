package calculator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Request is a struct that represents a request to the calculator
type Request struct {
	Parameters []float64 `json:"parameters"`
	UserId     string    `json:"userId"`
	Operation  string    `json:"operation"`
}

// Response is a struct that represents a response from the calculator
type Response struct {
	Result float64 `json:"result"`
}

// ParseHttpRequest parses http request to a calculator request
func ParseHttpRequest(r *http.Request) (*Request, error) {
	request := &Request{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

// GenerateRequestId generates a unique request id based on the request parameters
func (r *Request) GenerateRequestId() string {
	parametersStrings := make([]string, 0, len(r.Parameters))
	for _, param := range r.Parameters {
		parametersStrings = append(parametersStrings, fmt.Sprintf("%f", param))
	}
	// joins request parameters to string seperated by dash
	return fmt.Sprintf("%s-%s-%s", r.UserId, r.Operation, strings.Join(parametersStrings, "-"))
}
