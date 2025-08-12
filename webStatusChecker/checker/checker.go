package checker

import (
	reader "aleksale/webchecker/io"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func WebStatusCheck() {
	urls := reader.GetWebsites()

	ticker := time.NewTicker(5 * time.Second)

	for {
		fetched := make(chan reader.WebSite, len(urls))

		fetchOnTick(urls, fetched)

		for w := range fetched {
			fmt.Println(w)
		}

		<-ticker.C
	}
}

func fetchOnTick(urls []reader.WebSite, fetched chan reader.WebSite) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(website reader.WebSite) {
			defer wg.Done()
			fetchWebsite(url, fetched)
		}(url)
	}

	go func() {
		wg.Wait()
		close(fetched)
	}()
}

func fetchWebsite(website reader.WebSite, fetched chan reader.WebSite) {
	resp, err := http.Get(website.Url)

	if err != nil || resp == nil {
		fetched <- reader.WebSite{
			Url:      website.Url,
			Status:   http.StatusNotFound,
			ErrorMsg: err.Error(),
		}
		return
	}
	defer resp.Body.Close()

	fetched <- reader.WebSite{
		Url:      website.Url,
		Status:   resp.StatusCode,
		ErrorMsg: "",
	}
}
