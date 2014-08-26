package ddnsclient

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

var (
	domain     string
	username   string
	password   string
	updateURL  *url.URL
	updateFunc DnsUpdateFunc
)

type DnsUpdateFunc func() (string, error)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)

	domain = os.Getenv("DDNS_DOMAIN")
	if domain == "" {
		log.Fatal("Please specify DDNS_DOMAIN environment variable")
	}

	username = os.Getenv("DDNS_USERNAME")
	if username == "" {
		log.Fatal("Please specify DDNS_USERNAME environment variable")
	}

	password = os.Getenv("DDNS_PASSWORD")
	if password == "" {
		log.Fatal("Please specify DDNS_PASSWORD environment variable")
	}
}

func DNSUpdate() (string, error) {
	if updateFunc == nil {
		return "", errors.New("No valid DNS update function found")
	}

	result, err := updateFunc()
	if err != nil {
		return "", err
	}

	return result, nil
}

func DNSUpdateHandler(w http.ResponseWriter, r *http.Request) {
	result, err := DNSUpdate()
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}

	log.Println(result)
	fmt.Fprintf(w, "%s\n", result)
	return
}
