package codesextractor

import (
	"regexp"
	"strings"
)

// GenerateFCCodesMessageFromText generates a message containing FC codes found in the provided text.
func GenerateFCCodesMessageFromText(text string) string {
	codes := getFCCodesFromText(text)

	var sb strings.Builder

	// Append each code
	for _, code := range codes {
		sb.WriteString("- `")
		sb.WriteString(code)
		sb.WriteString("`\n")
	}

	// Add footer with a link to ForoCoches invitation codes redemption page
	sb.WriteString("\n[ForoCoches - Código de Invitación](https://forocoches.com/codigo)\n")

	return sb.String()
}

// Extracts FC codes from the given text.
func getFCCodesFromText(text string) []string {
	// Define the regular expression to match the hidden codes
	// This regex captures sequences like E_K_e_W_Q_k_D_u_5 or Q*J*G*Q*8*D*A*j*5 for example
	// It could capture E-K*e-W*Q*k-D*u-5 even though it is incorrect, but it is highly unlikely given the context
	codePattern := regexp.MustCompile(`([a-zA-Z0-9][^a-zA-Z0-9\s]){8}[a-zA-Z0-9]`)
	// Find all matches in the text
	matches := codePattern.FindAllString(text, -1)

	// Remove non alphanumeric characters from the matches
	codes := make([]string, 0, len(matches))
	for _, match := range matches {
		// Remove all non-alphanumeric characters from the match
		cleaned := strings.Map(func(r rune) rune {
			if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
				return r
			}
			// Remove the character
			return -1
		}, match)
		codes = append(codes, cleaned)
	}

	return codes
}
