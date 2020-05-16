package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"tools/copyer"
)

// copyerCmd represents the copyer command
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
		fmt.Println("copyer called")
	},
}

func init() {
	rootCmd.AddCommand(copyerCmd)

	//if src == "" {
	//	log.Fatal(&Error{"\"src\" flag must be defined\n"})
	//}
	//
	//_, err := Copy(src, dest)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
}
