// ip4lookup returns an ipv4 address for a given hostname
package main

import (
	"fmt"
	"net"
	"os"
)

var usage = `

ipv6lookup - returns an ipv4 address for a given hostname.

usage: ipv6lookup golang.org

`

func main() {

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	listip, err := net.LookupIP(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, i := range listip {

		ipv4 := i.To4()
		if ipv4 == nil {
			ipv6 := i.To16()
			if ipv6 != nil {
				fmt.Println(ipv6)
			}
		}
	}
}
