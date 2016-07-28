package data

import (
	"encoding/json"
	"strings"
	"unicode"
)

type Dictionary interface {
	Parse(string) string
}

type DefaultDictionary struct {
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

func LoadDefaultDictionary() (*DefaultDictionary, error) {
	binData, err := Asset("data/dictionary.json")
	if err != nil {
		return nil, err
	}

	return LoadJSON(binData)
}

func LoadJSON(b []byte) (*DefaultDictionary, error) {
	d := &DefaultDictionary{}
	err := json.Unmarshal(b, d)
	return d, err
}

func (d *DefaultDictionary) IsVowel(r rune) bool {
	return strings.ContainsRune(d.Data.Vowel, unicode.ToLower(r))
}

func (d *DefaultDictionary) IsConsonant(r rune) bool {
	return strings.ContainsRune(d.Data.Consonant, unicode.ToLower(r))
}

func (d *DefaultDictionary) IsNumber(r rune) bool {
	return strings.ContainsRune(d.Data.Number, unicode.ToLower(r))
}

func (d *DefaultDictionary) IsPunctuation(r rune) bool {
	return !(d.IsVowel(r) || d.IsConsonant(r))
}

func (d *DefaultDictionary) IsCaseSensitive(r rune) bool {
	return strings.ContainsRune(d.Data.CaseSensitive, unicode.ToLower(r))
}

func (d *DefaultDictionary) IsExact(needle string, haystack string, start int, end int, not bool) bool {
	return (start >= 0 && end < len(haystack) && (haystack[start:end] == (needle))) != not
}

func (d *DefaultDictionary) FixCase(s string) string {
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

func (d *DefaultDictionary) Parse(text string) string {
	fixed := d.FixCase(text)
	var output string

	for cur := 0; cur < len(fixed); cur++ {
		start := cur
		end := cur + 1
		prev := start - 1
		matched := false
		for _, pattern := range d.Data.Patterns {
			end = cur + len(pattern.Find)
			if (end <= len(fixed)) && fixed[start:end] == pattern.Find {
				prev = start - 1
				for _, rule := range pattern.Rules {
					replace := true

					chk := 0
					for _, match := range rule.Matches {
						if match.Type == "suffix" {
							chk = end
						} else {
							chk = prev
						}
						match.Negative = false
						if strings.HasPrefix(match.Scope, "!") {
							match.Negative = true
							match.Scope = match.Scope[1:]
						}
						switch match.Scope {
						case "punctuation":
							if ((chk < 0 && match.Type == "prefix") ||
								(chk >= len(fixed) && match.Type == "suffix") ||
								d.IsPunctuation(rune(fixed[chk]))) == match.Negative {
								replace = false
								break
							}
						case "vowel":
							if (((chk >= 0 && match.Type == "prefix") ||
								(chk < len(fixed) && match.Type == "suffix")) &&
								d.IsVowel(rune(fixed[chk]))) == match.Negative {
								replace = false
								break
							}
						case "consonant":
							if (((chk >= 0 && match.Type == "prefix") ||
								(chk < len(fixed) && match.Type == "suffix")) &&
								d.IsConsonant(rune(fixed[chk]))) == match.Negative {
								replace = false
								break
							}
						case "exact":
							var s, e int
							if match.Type == "suffix" {
								s = end
								e = end + len(match.Value)
							} else {
								s = start - len(match.Value)
								e = start
							}

							if !d.IsExact(match.Value, fixed, s, e, match.Negative) {
								replace = false
								break
							}
						}
					}
					if replace {
						output += rule.Replace
						cur = end - 1
						matched = true
						break
					}
				}
				if matched {
					break
				}

				// Default
				output += pattern.Replace
				cur = end - 1
				matched = true
				break
			}
		}
		if !matched {
			output += string(fixed[cur])
		}
	}
	return output
}
