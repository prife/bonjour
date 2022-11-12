package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/oleksandr/bonjour"
)

// _apple-mobdev2._tcp. 70:ea:5a:2a:88:ad@fe80::72ea:5aff:fe2a:88ad
func main() {
	// Run registration (blocking call)
	s, err := bonjour.RegisterProxy("70:ea:5a:2a:88:ad@fe80::72ea:5aff:fe2a:88ad", "_apple-mobdev._tcp.", "", 32498, "00008027-001E30411EE9802E", "127.0.39.237", nil, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Ctrl+C handling
	handler := make(chan os.Signal, 1)
	signal.Notify(handler, os.Interrupt)
	for sig := range handler {
		if sig == os.Interrupt {
			s.Shutdown()
			time.Sleep(1e9)
			break
		}
	}
}
