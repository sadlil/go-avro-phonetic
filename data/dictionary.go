package data

import (
	"encoding/json"
	"strings"
	"unicode"
)

type Dictionary struct {
	Meta struct {
		FileName         string `json:"file_name"`
		FileDescription  string `json:"file_description"`
		Package          string `json:"package"`
		License          string `json:"license"`
		Source           string `json:"source"`
		OriginalCode     string `json:"original_code"`
		InitialDeveloper string `json:"initial_developer"`
		Copyright        string `json:"copyright"`
		AdaptedBy        string `json:"adapted_by"`
		Encoding         string `json:"encoding"`
	} `json:"meta"`
	Data struct {
		Vowel         string `json:"vowel"`
		Consonant     string `json:"consonant"`
		CaseSensitive string `json:"case_sensitive"`
		Number        string `json:"number"`
		Patterns      []struct {
			Find    string `json:"find"`
			Replace string `json:"replace"`
			Rules   []struct {
				Matches []struct {
					Negative bool   `json:"negative"`
					Type     string `json:"type"`
					Scope    string `json:"scope"`
					Value    string `json:"value"`
				} `json:"matches"`
				Replace string `json:"replace"`
			} `json:"rules,omitempty"`
		} `json:"patterns"`
	} `json:"data"`
}

func Load() (*Dictionary, error) {
	binData, err := Asset("data/dictionary.json")
	if err != nil {
		return nil, err
	}

	return LoadJSON(binData)
}

func LoadJSON(b []byte) (*Dictionary, error) {
	d := &Dictionary{}
	err := json.Unmarshal(b, d)
	return d, err
}

func (d *Dictionary) IsVowel(r rune) bool {
	return strings.ContainsRune(d.Data.Vowel, unicode.ToLower(r))
}

func (d *Dictionary) IsConsonant(r rune) bool {
	return strings.ContainsRune(d.Data.Consonant, unicode.ToLower(r))
}

func (d *Dictionary) IsNumber(r rune) bool {
	return strings.ContainsRune(d.Data.Number, unicode.ToLower(r))
}

func (d *Dictionary) IsPunctuation(r rune) bool {
	return !(d.IsVowel(r) || d.IsConsonant(r))
}

func (d *Dictionary) IsCaseSensitive(r rune) bool {
	return strings.ContainsRune(d.Data.CaseSensitive, unicode.ToLower(r))
}

func (d *Dictionary) IsExact(needle string, haystack string, start int, end int, not bool) bool {
	return (start >= 0 && end < len(haystack) && (haystack[start:end] == (needle))) != not
}

func (d *Dictionary) FixCase(s string) string {
	fixed := make([]rune, 0)
	for _, c := range s {
		if d.IsCaseSensitive(c) {
			fixed = append(fixed, c)
		} else {
			fixed = append(fixed, unicode.ToLower(c))
		}
	}
	return string(fixed)
}
