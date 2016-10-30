package wikipedia

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Wikipedia holds the base Page name of the wiki page to which later the
// API call will get appended and the username to the wiki.
type Wikipedia struct {
	Page string
	User string
}

// GetUserEdits calls wikiUrl Wikipedia API to retrieve the number of edits
// the user username has done.
func GetUserEdits(wikiPage string, username string) (count int, err error) {
	wikiURL, err := url.Parse("https://tools.wmflabs.org/sigma/usersearch.py")
	if err != nil {
		return 0, errors.New("Not a valid URL")
	}
	parameters := url.Values{}
	parameters.Add("max", "500")
	parameters.Add("server", "enwiki")
	parameters.Add("page", wikiPage)
	parameters.Add("name", username)
	wikiURL.RawQuery = parameters.Encode()

	resp, err := http.Get(wikiURL.String())
	if err != nil {
		return 0, errors.New("Not able to HTTP GET")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split(string(body), "\n") {
		line = strings.TrimSpace(line)
		if len(line) < 6 {
			continue
		}
		if line[0:5] == "Found" {
			count, err = strconv.Atoi(strings.Split(line, " ")[1])
		}
	}
	return
}
