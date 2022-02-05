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
)

func main() {
    config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
    c := new(dns.Client)

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(os.Args[1]), dns.TypeA)
    m.SetUpdate(dns.Fqdn(os.Args[1]))
    rrInsert, err := dns.NewRR("rrsetnotused.free2air.net 600 IN A 127.0.0.1")
    if err != nil {
        panic(err)
    }
    m.Insert([]dns.RR{rrInsert})
    spew.Dump(m.Answer)


    spew.Dump(m)
    // m.RecursionDesired = true
    config.Servers[0] = "157.90.22.121"
    r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    // r, _, err := c.Exchange(m, net.JoinHostPort("abulafia.free2air.net", config.Port))
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
            // log.Fatalf(" *** invalid answer name %s after KEY query for %s\n", os.Args[1], os.Args[1])
            //spew.Dump(r.Answer)
	    fmt.Printf(" ***  %s after Exchange for %s\n", os.Args[1], os.Args[1])
    }
    // Stuff must be in the answer section
    for _, a := range r.Answer {
            fmt.Printf("%v\n", a)
    }
}

