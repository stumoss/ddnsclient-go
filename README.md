Dynamic DNS Updater Client
==========================

A project to enable dynamic DNS updating exposed as a web service.

To make use of this you must set the following environment variables:
* DDNS_DOMAIN
* DDNS_USERNAME
* DDNS_PASSWORD

Example Usage:
```go
package main

import "bitbucket.org/stumoss/ddnsclient"

func main() {
    http.HandleFunc("/dns_update", ddnsclient.UpdateDnsEntries)
    http.ListenAndServe(":" + strconv.Itoa(*port), nil)
}```

Once you have the main app up and running you can use curl or a web browser
to trigger the DNS update:

    curl -X POST http://localhost/dns_update


Supported DNS Provider
----------------------
[ZoneEdit](http://www.zoneedit.com "ZoneEdit")
