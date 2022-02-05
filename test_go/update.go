// from https://miek.nl/2014/august/16/go-dns-package/
// usage:
//	go run update.go <domain>
// where <domain> is the FQDN for the update
package main

import (
    "github.com/miekg/dns"
    "github.com/davecgh/go-spew/spew"
    "net"
    "os"
    "log"
    "fmt"
    // "strings"
)

func main() {

    zone, ok := os.LookupEnv("GD_ZONE")
    if !ok {
            fmt.Println("No GD_ZONE ENV var defined")
	    zone := os.Args[1]
	    fmt.Printf("Using cmdline zone value: %s\n", zone)
    } else {
            fmt.Printf("GD_ZONE = %s\n", zone)
    }
    host, ok := os.LookupEnv("GD_HOST")
    if !ok {
            fmt.Println("No GD_HOST ENV var defined")
	    host := "go-dns-test"
	    fmt.Printf("Using cmdline host value: %s\n", host)
    } else {
            fmt.Printf("GD_HOST = %s\n", host)
    }
    server, ok := os.LookupEnv("GD_SERVER")
    if !ok {
            fmt.Println("No GD_SERVER ENV var defined")
	    server := os.Args[2]
	    fmt.Printf("Using cmdline host value: %s\n", server)
    } else {
            fmt.Printf("GD_SERVER = %s\n", server)
    }
    fqdn := fmt.Sprintf("%s.%s", host, zone)
    fmt.Printf("Using FQDN RR entry of %s\n", fqdn)

    myRR := fmt.Sprintf("%s.%s 600 IN A 127.0.0.1", host, zone)
    fmt.Printf("myRR = %s", myRR)

    config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
    c := new(dns.Client)

    m := new(dns.Msg)
    m.SetUpdate(dns.Fqdn(zone))

    rrInsert, err := dns.NewRR(myRR)
    if err != nil {
        panic(err)
    }
    m.Insert([]dns.RR{rrInsert})

    spew.Dump(m.Answer)

    spew.Dump(m)

    r, _, err := c.Exchange(m, net.JoinHostPort(server, config.Port))
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
	    fmt.Printf(" ***  %s after Exchange(%s) for %s\n", server, os.Args[1])
    }
    // Stuff must be in the answer section
    for _, a := range r.Answer {
            fmt.Printf("%v\n", a)
    }
}

