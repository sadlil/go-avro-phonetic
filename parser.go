package avro_phonetic

import (
	"strings"

	"github.com/sadlil/go-avro-phonetic/data"
)

func Parse(text string) (string, error) {
	data, err := data.Load()
	if err != nil {
		return "", err
	}
	return ParseWithDictionary(data, text), nil
}

func ParseOrDie(test string) string {
	res, err := Parse(test)
	if err != nil {
		panic("Failed to parse string, cause: " + err.Error())
	}
	return res
}

func ParseWithDictionary(d *data.Dictionary, text string) string {
	fixed := d.FixCase(text)
	return parse(d, fixed)
}

func parse(d *data.Dictionary, fixed string) string {
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

						if match.Scope == "punctuation" {
							if ((chk < 0 && match.Type == "prefix") ||
								(chk >= len(fixed) && match.Type == "suffix") ||
								d.IsPunctuation(rune(fixed[chk]))) == match.Negative {
								replace = false
								break

							}
						} else if match.Scope == "vowel" {
							if (((chk >= 0 && match.Type == "prefix") ||
								(chk < len(fixed) && match.Type == "suffix")) &&
								d.IsVowel(rune(fixed[chk]))) == match.Negative {
								replace = false
								break
							}
						} else if match.Scope == "consonant" {
							if (((chk >= 0 && match.Type == "prefix") ||
								(chk < len(fixed) && match.Type == "suffix")) &&
								d.IsConsonant(rune(fixed[chk]))) == match.Negative {
								replace = false
								break
							}
						} else if match.Scope == "exact" {
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
