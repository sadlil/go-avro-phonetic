package main

import (
	"io/ioutil"
	"os"
	"strings"

	avro "github.com/sadlil/go-avro-phonetic"
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
		os.Stderr.WriteString("Failed to parse text, error" + err.Error())
	}
	os.Stdout.WriteString(text)
}

func parseFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		os.Stderr.WriteString("Failed to parse file, error" + err.Error())
	}
	text, err := avro.Parse(string(data))
	if err != nil {
		os.Stderr.WriteString("Failed to parse text, error" + err.Error())
	}
	destFileName := filePath[:strings.LastIndex(filePath, ".")] +
		".bn" +
		filePath[strings.LastIndex(filePath, "."):]
	err = ioutil.WriteFile(destFileName, []byte(text), os.ModePerm)
	if err != nil {
		os.Stderr.WriteString("Failed to write text to dest file, error" + err.Error())
	}
}
