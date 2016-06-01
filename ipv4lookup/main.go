// ip4lookup returns an ipv4 address for a given hostname
package main

import (
	"fmt"
	"net"
	"os"
)

var usage = `

ip4lookup - returns an ipv4 address for a given hostname.

usage: ipv4lookup golang.org

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
		if ipv4 != nil {
			fmt.Println(ipv4)
		}
	}
}
