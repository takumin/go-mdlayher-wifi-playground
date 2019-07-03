package main

import (
	"log"

	"github.com/mdlayher/wifi"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	client, err := wifi.New()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	devices, err := client.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range devices {
		stations, err := client.StationInfo(v)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range stations {
			log.Println(v)
		}
	}
}
