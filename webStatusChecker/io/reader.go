package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	filePath = "urls.txt"
)

type WebSite struct {
	Status   int
	Url      string
	ErrorMsg string
}

func (website WebSite) String() string {
	if website.ErrorMsg != "" {
		return fmt.Sprintf("[%d] %s - %s", website.Status, website.Url, website.ErrorMsg)
	} else {
		return fmt.Sprintf("[%d] %s", website.Status, website.Url)
	}
}

func GetWebsites() []WebSite {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Did not find urls.txt file!", err)
	}
	defer file.Close()

	var websites []WebSite
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			websites = append(websites, WebSite{Url: line})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error with scanner!", err)
	}

	return websites
}
