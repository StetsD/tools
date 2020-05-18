package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"tools/envparser"
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
		if file == "" {
			log.Fatal("\"file\" flag must be defined\n")
		}

		envs, err := envparser.Parse(file)
		if err != nil {
			log.Fatal(err)
		}
		envparser.Apply(envs)

	},
}

func init() {
	rootCmd.AddCommand(envparserCmd)

	envparserCmd.Flags().StringVarP(&file, "file", "f", "", "file with env vars")
}
