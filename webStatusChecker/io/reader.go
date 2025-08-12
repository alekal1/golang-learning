package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	filePath    = "urls.txt"
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

type WebSite struct {
	Status   int
	Url      string
	ErrorMsg string
}

type InvalidUrl struct {
	msg string
}

func (e InvalidUrl) Error() string {
	return e.msg
}

func (website WebSite) String() string {
	if website.ErrorMsg != "" {
		return fmt.Sprintf("[%d] %s - %s", website.Status, website.Url, website.ErrorMsg)
	} else {
		return fmt.Sprintf("[%d] %s", website.Status, website.Url)
	}
}

func GetWebsites() ([]error, []WebSite) {
	file := openFile()
	defer file.Close()

	return scanFile(file)
}

func openFile() *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Did not find urls.txt file!", err)
	}

	return file
}

func scanFile(file *os.File) ([]error, []WebSite) {
	var websites []WebSite
	var errors []error

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if err := validateLineInput(line); err != nil {
			errors = append(errors, err)
		} else {
			websites = append(websites, WebSite{Url: line})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error with scanner!", err)
	}

	return errors, websites
}

func validateLineInput(line string) error {
	if line == "" || !(strings.HasPrefix(line, httpPrefix) || strings.HasPrefix(line, httpsPrefix)) {
		return InvalidUrl{fmt.Sprintf("URL '%s' will not be fetched. Should start with '%s' or '%s'", line, httpPrefix, httpsPrefix)}
	}
	return nil
}
