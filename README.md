Dynamic DNS Updater Client
==========================
[![Build Status](https://travis-ci.org/stumoss/ddnsclient-go.svg?branch=master)](https://travis-ci.org/stumoss/ddnsclient-go)

A project to enable dynamic DNS updating exposed as a web service. Currently
the only supported dynamic DNS provider supported is [ZoneEdit](http://www.zoneedit.com "ZoneEdit")

To make use of this you must set the following environment variables:
* DDNS_DOMAIN
* DDNS_USERNAME
* DDNS_PASSWORD

Example Usage:
```go
package main

import "bitbucket.org/stumoss/ddnsclient"

var (
    port *int
)

func init() {
    port = flag.Int("port", 8080, "port to listen on")
    flag.Parse()
}

func main() {
    http.HandleFunc("/dns_update", ddnsclient.DNSUpdateHandler)
    ListenAndServe(":" + strconv.Itoa(*port), nil)
}
```

Once you have the main app up and running you can use curl or a web browser
to trigger the DNS update:

    curl -X POST http://localhost/dns_update


Supported DNS Provider
----------------------
[ZoneEdit](http://www.zoneedit.com "ZoneEdit")
