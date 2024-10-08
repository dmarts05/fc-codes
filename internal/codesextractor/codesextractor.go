package codesextractor

import (
	"regexp"
	"strings"
)

func GetForocochesCodesFromText(text string) []string {
	// Define the regular expression to match the hidden codes
	// This regex captures sequences like E.K.e.W.Q.k.D.u.5 or Q.J.G.Q.8.D.A.j.5
	codePattern := regexp.MustCompile(`([a-zA-Z0-9]\.){8}[a-zA-Z0-9]`)

	// Find all matches in the text
	matches := codePattern.FindAllString(text, -1)

	// Remove periods
	codes := make([]string, 0, len(matches))
	for _, match := range matches {
		codes = append(codes, strings.ReplaceAll(match, ".", ""))
	}

	return codes
}
