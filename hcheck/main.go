package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	tcp "github.com/tevino/tcp-shaker"
)

func main() {
	s := tcp.Shaker{}
	if err := s.Init(); err != nil {
		log.Fatal("Shaker init failed:", err)
	}
	var err error
	hostname := "localhost"
	port := "80"
	timeout := time.Second * 1
	delay := time.Microsecond * 1000

	if len(os.Args) == 2 {
		hostname = os.Args[1]

	}

	if len(os.Args) == 3 {
		hostname = os.Args[1]
		delay, err = time.ParseDuration(os.Args[2])
	}

	if len(os.Args) == 4 {
		hostname = os.Args[1]
		delay, err = time.ParseDuration(os.Args[2])
		timeout, err = time.ParseDuration(os.Args[3])
	}

	if err != nil {

		panic(err)
	}
	hostnameport := "localhost.com:80"

	for portn := 20; portn < 10000; portn++ {

		//fmt.Println(portn)
		port = strconv.Itoa(portn)

		var buffer bytes.Buffer // small!
		buffer.WriteString(hostname)
		buffer.WriteString(":")
		buffer.WriteString(port)

		hostnameport = buffer.String()

		err := s.Test(hostnameport, timeout)
		switch err {
		case tcp.ErrTimeout:
			fmt.Printf("Connect to %s timeout\n", hostnameport)
		case nil:
			fmt.Printf("Connect to %s succeeded\n", hostnameport)
		default:
			if e, ok := err.(*tcp.ErrConnect); ok {
				fmt.Fprintf(os.Stderr, "Connect to %s failed: %v\n", hostnameport, e)
			} else {
				fmt.Fprintf(os.Stderr, "Error occurred while connecting: %s\n", err)
			}
		}
		time.Sleep(delay)
	}

}
