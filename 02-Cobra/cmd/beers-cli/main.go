package main

import (
	"github.com/Leasilvaz/test_go.git/02-Cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.Execute()
}
