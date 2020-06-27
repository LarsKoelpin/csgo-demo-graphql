package usecase

import (
	"github.com/graphql-go/graphql/language/parser"
	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
)

func ParseQuery(query string) domain.DemoTemplate {

	AST, _ := parser.Parse(parser.ParseParams{Source: query})

	println(AST)
	return domain.DemoTemplate{}

}
