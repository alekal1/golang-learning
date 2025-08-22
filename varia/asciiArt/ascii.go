package asciiart

import (
	"strings"
)

const (
	REFERENCE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?"
)

func AsciiArt(letterWidth, letterHeigh int, text string, reference []string) string {
	splited := strings.Split(text, "")
	var output string

	for i := 0; i < letterHeigh; i++ {
		row := reference[i]

		results := splitEveryN(row, letterWidth)

		for _, r := range splited {
			var indexOfReference int
			if !strings.Contains(REFERENCE, r) {
				indexOfReference = strings.Index(REFERENCE, "?")
			} else {
				indexOfReference = strings.Index(REFERENCE, r)
			}

			output += results[indexOfReference]
		}

		output += "\n"
	}

	return output
}

func splitEveryN(s string, n int) []string {
	var chunks []string
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}
