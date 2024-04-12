package cmd

import (
	"fmt"
	difyconnector "github.com/leslieleung/dify-connector/cmd/dify-connector"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "dify-connector",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(difyconnector.ServeCmd)
}
