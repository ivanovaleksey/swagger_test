//go:generate swagger generate spec -o ../swagger.json

// Package main Title
//
// Description
//
// BasePath: /api
// Version: 0.0.1
// Produces:
//  - application/json
// Schemes: http, https
// swagger:meta
package main

import (
	"fmt"
	"github.com/ivanovaleksey/swagger_test"
	"github.com/ivanovaleksey/swagger_test/color"
)

// PersonResponse
//swagger:response personResponse
type personResp struct {
	// in:body
	Body swagger_test.Person
}

// swagger:route GET /hello
//     Responses:
//       200: personResponse
func Test() {}

func main() {
	p := swagger_test.Person{
		FullName: "John Doe",
		Color: color.Color{
			"Red",
		},
	}
	fmt.Println(p)
}
