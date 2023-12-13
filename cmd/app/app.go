package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/monandkey/panforyou/internal/pkg/adapter/controller"
	"github.com/monandkey/panforyou/internal/pkg/adapter/presenter"
	"github.com/monandkey/panforyou/internal/pkg/adapter/repository"
	"github.com/monandkey/panforyou/internal/pkg/adapter/resolver"
	"github.com/monandkey/panforyou/internal/pkg/infrastructure/database"
	"github.com/monandkey/panforyou/internal/pkg/usecase"
)

const defaultPort = "8080"

func runApp() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.NewDatabaseSQLFactory(database.InstancePostgres)
	if err != nil {
		return err
	}

	srv := handler.NewDefaultServer(resolver.NewSchema(
		controller.NewGraphQLController(
			usecase.NewFindBreadUsecase(
				repository.NewBreadSQL(db),
				presenter.NewFindBreadPresenter(),
			),
		),
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
