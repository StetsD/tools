package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"tools/copyer"
)

var src, dest string
var bytesLimit int
var wd string

var copyerCmd = &cobra.Command{
	Use:   "copyer",
	Short: "Recursive or by one copy your files",
	Long: `
NAME
	copyer - recursive copying or by one your files with autoincrement names

SYNOPSIS
	tools copyer -src PATH -dest PATH

DESCRIPTION
	-limit
		set bytes limit for your files during copying
`,
	Run: func(cmd *cobra.Command, args []string) {
		if src == "" {
			log.Fatal("\"src\" flag must be defined\n")
		}

		_, err := copyer.Copy(src, dest, bytesLimit)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(copyerCmd)

	_, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	wd, _ = os.Getwd()

	copyerCmd.Flags().StringVarP(&src, "src", "s", "", "path from")
	copyerCmd.Flags().StringVarP(&dest, "dest", "d", wd, "path to")
	copyerCmd.Flags().IntVarP(&bytesLimit, "limit", "l", 0, "limit of bytes for other file")
}
