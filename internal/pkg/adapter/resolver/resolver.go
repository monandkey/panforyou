package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/monandkey/panforyou/internal/pkg/adapter/controller"
	graphql1 "github.com/monandkey/panforyou/internal/pkg/infrastructure/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	controller controller.GraphQLController
}

func NewSchema(controller controller.GraphQLController) graphql.ExecutableSchema {
	return graphql1.NewExecutableSchema(graphql1.Config{
		Resolvers: &Resolver{
			controller: controller,
		},
	})
}
