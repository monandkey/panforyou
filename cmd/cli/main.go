package main

import (
	"os"
	"time"

	"github.com/monandkey/panforyou/internal/pkg/adapter/controller"
	"github.com/monandkey/panforyou/internal/pkg/adapter/presenter"
	"github.com/monandkey/panforyou/internal/pkg/adapter/repository"
	"github.com/monandkey/panforyou/internal/pkg/infrastructure/api/contentful"
	"github.com/monandkey/panforyou/internal/pkg/infrastructure/database"
	"github.com/monandkey/panforyou/internal/pkg/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Use = "cli"
	rootCmd.Short = "This command accesses the Content Delivery API to retrieve information about the specified bread and store it in the DB"
	rootCmd.Version = "1.0"
	rootCmd.SilenceUsage = true

	var entryID string
	rootCmd.Flags().StringVarP(&entryID, "entry-id", "e", entryID, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if entryID == "" {
			return rootCmd.Help()
		}

		db, err := database.NewDatabaseSQLFactory(database.InstancePostgres)
		if err != nil {
			return err
		}

		cmdCtr := controller.NewCommandController(
			usecase.NewFindContentfulUsecase(
				repository.NewContentfulAPI(
					contentful.NewContentfulAPIFactory(),
				),
				presenter.NewFindContentfulPresenter(),
			),
			usecase.NewCreateBreadUsecase(
				repository.NewBreadSQL(db),
				time.Second,
			),
			entryID,
		)
		if err := cmdCtr.Create(); err != nil {
			return err
		}
		return nil
	}
}

func main() {
	execute()
}
