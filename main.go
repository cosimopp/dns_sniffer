package main

import (
    "net"
    "github.com/miekg/dns"
    "fmt"
)

func main() {
    laddr := &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 53}
    resolver := "8.8.8.8:53"

    conn, _ := net.ListenUDP("udp4", laddr)
    defer conn.Close()
    for {
        buf := make([]byte, 512)
        _, raddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            continue
        }

        req := new(dns.Msg)
        req.Unpack(buf)

        name := req.Question[0].Name
        fmt.Println(raddr, name) 
        
        client := new(dns.Client)
        res, _, err := client.Exchange(req, resolver)
        if err != nil {
            continue
        }
        pack, err := res.Pack()
        if err != nil {
            continue
        }
        conn.WriteToUDP(pack, raddr) 
    }
}
