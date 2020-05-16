package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var file string

// envparserCmd represents the envparser command
var envparserCmd = &cobra.Command{
	Use:   "envparser",
	Short: "Parse env variables from file and apply them",
	Long: `
NAME
	envparser - recursive copying or by one your files with autoincrement names

SYNOPSIS
	tools envparser --file PATH

DESCRIPTION
	--file
		path to file
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(os.Environ())
	},
}

func init() {
	rootCmd.AddCommand(envparserCmd)

	envparserCmd.Flags().StringVarP(&file, "file", "f", "", "file with env vars")
}
