package main

import (
	"flag"
	"fmt"
	"strconv"
)


func main() {
	var (
		port *int = flag.Int("p", 8000, "Select a port to run the service")
		run *bool = flag.Bool("s", false, "Run the service (Mandatory)")
	)
	flag.Parse()

	if *run {
		fmt.Println("[*] Service is running in the port " + strconv.Itoa(*port))
	}
}
