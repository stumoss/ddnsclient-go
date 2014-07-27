package ddnsclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

var (
	ddns_domain   string
	ddns_username string
	ddns_password string
	responseRegex = regexp.MustCompile("<\\w.*=\\W(\\d+)\\W\\s+TEXT=\\W(\\w.*)\\.\\W")
)

const (
	update_url    = "https://dynamic.zoneedit.com/auth/dynamic.html?host=*."
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)

	ddns_domain = os.Getenv("DDNS_DOMAIN")
	if ddns_domain == "" {
		log.Fatal("Please specify DDNS_DOMAIN environment variable")
	}

	ddns_username = os.Getenv("DDNS_USERNAME")
	if ddns_username == "" {
		log.Fatal("Please specify DDNS_USERNAME environment variable")
	}

	ddns_password = os.Getenv("DDNS_PASSWORD")
	if ddns_password == "" {
		log.Fatal("Please specify DDNS_PASSWORD environment variable")
	}
}

func UpdateDnsEntries(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: DNS update requested\n", ddns_domain)

	client := &http.Client{}
	req, err := http.NewRequest("GET", update_url+ddns_domain, nil)
	if err != nil {
		log.Printf("%s: %s", ddns_domain, err)
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(ddns_username, ddns_password)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s: %s", ddns_domain, err)
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: %s", ddns_domain, err)
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}

	matches := responseRegex.FindStringSubmatch(string(bodyStr))
	if len(matches) != 3 {
		log.Printf("%s: failed to parse response: %s", ddns_domain, string(bodyStr))
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}

	log.Printf("%s: %s", ddns_domain, matches[2])
	fmt.Fprintf(w, "%s: %s", ddns_domain, matches[2])
}
