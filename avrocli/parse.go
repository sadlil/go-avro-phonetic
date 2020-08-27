package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sadlil/go-avro-phonetic"
	"github.com/spf13/cobra"
)

func NewParseCmd() *cobra.Command {
	var filePath string
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "file or text to parse",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("file").Changed {
				if filePath != "" {
					parseFile(filePath)
					return
				}
			}
			if len(args) > 0 {
				text := strings.Join(args, " ")
				parse(text)
			}
		},
	}
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "file location to parse")
	return cmd
}

func parse(text string) {
	text, err := avro.Parse(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse text, error: " + err.Error())
		return
	}
	fmt.Fprintln(os.Stdout, text)
}

func parseFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse file, error: " + err.Error())
		return
	}
	text, err := avro.Parse(string(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse text, error: " + err.Error())
		return
	}
	destFileName := filePath[:strings.LastIndex(filePath, ".")] +
		".bn" +
		filePath[strings.LastIndex(filePath, "."):]
	err = ioutil.WriteFile(destFileName, []byte(text), os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write text to the destination file, error: " + err.Error())
	}
}
