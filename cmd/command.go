package cmd

import (
	serve "github.com/pius706975/golang-boilerplate-with-gin/api"
	"github.com/pius706975/golang-boilerplate-with-gin/cmd/custom"
	"github.com/pius706975/golang-boilerplate-with-gin/package/database"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "backend",
	Long: "Go backend service",
}

func init() {
	initCommand.AddCommand(serve.ServeCMD)
	initCommand.AddCommand(database.MigrationCMD)
	initCommand.AddCommand(custom.CreateMigrationCMD)
	initCommand.AddCommand(custom.CreateSuperUserCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}