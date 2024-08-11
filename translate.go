package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var pattern = regexp.MustCompile("(?i)(§[0-9a-fklmnor]|\\b)([b-df-hj-np-tv-xz]*[aeiouy])(\\w*)\\b")

func translate(v any) any {
	switch t := v.(type) {
	case string:
		//fmt.Printf("string: %s\n", t)
		return translateText(t)
	default:
		fmt.Printf("WARN: ignoring unsupported type: %v\n", t)
	}

	return v
}

func translateText(s string) string {
	if len(strings.Trim(s, " ")) < 2 {
		return s
	}

	s = regexReplaceAllStringFunc(pattern, s, func(groups []string) string {
		word := modifyWord(groups[2], IsUpper(groups[2]+groups[3]))
		if len(word+groups[3]) <= 5 {
			return groups[0]
		}
		return groups[1] + word + groups[3]
	})

	return s
}

func modifyWord(word string, isUppercase bool) string {
	c := rune(word[0])
	if unicode.IsLower(c) || c == 'Y' {
		return word
	}

	vowel := word[len(word)-1]
	if isUppercase {
		return strings.ToUpper(*prefix) + string(vowel)
	}
	return *prefix + string(unicode.ToLower(rune(vowel)))
}

func regexReplaceAllStringFunc(re *regexp.Regexp, s string, repl func(groups []string) string) string {
	var (
		result    string
		lastIndex int
	)
	for _, match := range re.FindAllStringSubmatchIndex(s, -1) {
		var matchStart, matchEnd = match[0], match[1]
		result += s[lastIndex:matchStart]

		groups := make([]string, 0, len(match)/2)
		for i := 0; i < len(match)-1; i += 2 {
			var groupStart, groupEnd = match[i], match[i+1]
			if groupStart == -1 || groupEnd == -1 {
				groups = append(groups, "")
			} else {
				groups = append(groups, s[groupStart:groupEnd])
			}
		}

		result += repl(groups)
		lastIndex = matchEnd
	}

	result += s[lastIndex:]
	return result
}

// IsUpper is from https://stackoverflow.com/a/59293875
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsLower is from https://stackoverflow.com/a/59293875
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
