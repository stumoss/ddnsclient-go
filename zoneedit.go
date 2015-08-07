package ddnsclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var responseRegex = regexp.MustCompile("<\\w.*=\"(\\d+)\"\\s+TEXT=\"(\\w[^\"]*)\"\\s+ZONE=\"([^\"]*)\">")

func init() {
	updateFunc = ZoneEditUpdate
}

func ZoneEditUpdate() (string, error) {
	log.Printf("%s: ZoneEdit DNS update requested\n", domain)

	var err error
	updateURL, err = url.Parse("https://dynamic.zoneedit.com/auth/dynamic.html")
	if err != nil {
		log.Printf("%s: %s\n", domain, err)
		return "", err
	}

	parameters := url.Values{}
	parameters.Add("host", "*."+domain)
	updateURL.RawQuery = parameters.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", updateURL.String(), nil)
	if err != nil {
		log.Printf("%s: %s\n", domain, err)
		return "", err
	}

	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s: %s\n", domain, err)
		return "", err
	}
	defer resp.Body.Close()

	bodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: %s\n", domain, err)
		return "", err
	}

	matches := responseRegex.FindStringSubmatch(string(bodyStr))
	if len(matches) != 4 {
		log.Printf("%s: failed to parse response: %s\n", domain, string(bodyStr))
		return "", err
	}

	return fmt.Sprintf("%s: %s", domain, matches[2]), nil
}
