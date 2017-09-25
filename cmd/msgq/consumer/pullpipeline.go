//pullpipeline.go

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ibmendoza/msgq"
)

//Example:
//pullpipeline -ip tcp://192.168.0.150:65000

var ip = flag.String("ip", "", "IP address")

func main() {
	flag.Parse()

	if *ip == "" {
		log.Fatal("ip address required")
	}

	pullpipeline, err := msgq.NewPullPipeline(*ip)
	if err != nil {
		log.Fatal("Error creating pullpipeline socket")
	}

	err = msgq.Listen(pullpipeline, *ip)
	if err != nil {
		log.Fatal("Error listening to pullpipeline socket")
	}

	for {
		msg, e := msgq.Receive(pullpipeline)
		if e != nil {
			log.Println("Error receiving message from pullpipeline socket")
		}

		fmt.Println(string(msg))
	}
}
