package ddnsclient

import (
	"log"
	"os"
	"runtime"
	"net/url"
)

type DDNSUpdateArgs struct {
	domain   string
	username string
	password string
	url      *url.URL
}

var dns DDNSUpdateArgs

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)

	dns.domain = os.Getenv("DDNS_DOMAIN")
	if dns.domain == "" {
		log.Fatal("Please specify DDNS_DOMAIN environment variable")
	}

	dns.username = os.Getenv("DDNS_USERNAME")
	if dns.username == "" {
		log.Fatal("Please specify DDNS_USERNAME environment variable")
	}

	dns.password = os.Getenv("DDNS_PASSWORD")
	if dns.password == "" {
		log.Fatal("Please specify DDNS_PASSWORD environment variable")
	}
}
