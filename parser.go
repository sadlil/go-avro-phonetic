package avro

import (
	"github.com/sadlil/go-avro-phonetic/data"
)

func Parse(text string) (string, error) {
	data, err := data.LoadDefaultDictionary()
	if err != nil {
		return "", err
	}
	return ParseWith(data, text), nil
}

func MustParse(test string) string {
	res, err := Parse(test)
	if err != nil {
		panic("Failed to parse string, cause: " + err.Error())
	}
	return res
}

func ParseWith(d data.Dictionary, text string) string {
	return d.Parse(text)
}
