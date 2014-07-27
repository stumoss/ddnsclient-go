// +build zoneedit

package ddnsclient

import (
    "regexp"
    "net/http"
    "net/url"
    "log"
    "io/ioutil"
    "fmt"
)

var responseRegex = regexp.MustCompile("<\\w.*=\\W(\\d+)\\W\\s+TEXT=\\W(\\w.*)\\.\\W")

func UpdateDnsEntries(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s: DNS update requested\n", dns.domain)

    var err error
    dns.url, err = url.Parse("https://dynamic.zoneedit.com/auth/dynamic.html") //?host=*."
    if err != nil {
        log.Printf("%s: %s", dns.domain, err)
        http.Error(w, "oops", http.StatusInternalServerError)
    }

    parameters := url.Values{}
    parameters.Add("host", "*."+dns.domain)
    dns.url.RawQuery = parameters.Encode()

    client := &http.Client{}
    req, err := http.NewRequest("GET", dns.url.String(), nil)
    if err != nil {
        log.Printf("%s: %s", dns.domain, err)
        http.Error(w, "oops", http.StatusInternalServerError)
        return
    }

    req.SetBasicAuth(dns.username, dns.password)
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("%s: %s", dns.domain, err)
        http.Error(w, "oops", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    bodyStr, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("%s: %s", dns.domain, err)
        http.Error(w, "oops", http.StatusInternalServerError)
        return
    }

    matches := responseRegex.FindStringSubmatch(string(bodyStr))
    if len(matches) != 3 {
        log.Printf("%s: failed to parse response: %s", dns.domain, string(bodyStr))
        http.Error(w, "oops", http.StatusInternalServerError)
        return
    }

    log.Printf("%s: %s", dns.domain, matches[2])
    fmt.Fprintf(w, "%s: %s", dns.domain, matches[2])
}
