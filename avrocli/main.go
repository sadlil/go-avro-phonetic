package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "avrocli",
		Short:   "command line tool for Avro Phonetic",
		Example: "avrocli parse amar sOnar bangla",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(NewParseCmd())
	return rootCmd
}
