package main

import (
	"os"
	"strings"

	"github.com/aerokite/google-translate-go/pkg"
	avro "github.com/sadlil/go-avro-phonetic"
	"github.com/spf13/cobra"
)

func NewTranslateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "translate",
		Short: "file or text to translate",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				text := strings.Join(args, " ")
				translate(text)
			}
		},
	}
	return cmd
}

func translate(text string) {
	text, err := avro.Parse(text)
	if err != nil {
		os.Stderr.WriteString("Failed to parse text, error" + err.Error())
	}

	req := pkg.TranslateRequest{
		SourceLang: "bn",
		TargetLang: "en",
		Text:       text,
	}
	translated, err := pkg.Translate(req)
	if err != nil {
		os.Stderr.WriteString("Failed to translate text, error" + err.Error())
	}

	os.Stdout.WriteString(translated)
}
