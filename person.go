package swagger_test

import "github.com/ivanovaleksey/swagger_test/color"

type Person struct {
	FullName string `json:"full_name"`

	// Color      color.Color            `json:"color"`

	// AliasColor color.AliasColor `json:"alias_color"`

	MyAliasColor PersonColor `json:"my_alias_color"`

	Info PersonInfo `json:"info"`
}

type PersonColor color.Color

type PersonInfo Info

type Info struct {
	Foo string `json:"foo"`
}
