package cmd

import (
	serve "go-gin/api"
	"go-gin/package/database"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "backend",
	Long: "Go backend service",
}

func init() {
	initCommand.AddCommand(serve.ServeCMD)
	initCommand.AddCommand(database.MigrationCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}