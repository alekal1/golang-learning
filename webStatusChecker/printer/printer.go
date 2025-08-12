package printer

import (
	reader "aleksale/webchecker/io"
	"log"
)

func PrintErrors(errors []error) {
	for _, err := range errors {
		log.Println(err)
	}
}

func PrintChanWebsite(channel chan reader.WebSite) {
	for w := range channel {
		log.Println(w)
	}
}
