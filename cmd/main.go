package main

import (
	"flag"
	"fmt"
	"strconv"
	"net/http"
	
	"github.com/jean0t/cool-bank/internal/route"
)


func main() {
	var (
		port *int = flag.Int("p", 8000, "Select a port to run the service")
		run *bool = flag.Bool("s", false, "Run the service (Mandatory)")
		migrate *bool = flag.Bool("M", false, "Make the migration")
	)
	flag.Parse()

	if *migrate {
		//database.MigrateDB()
	}

	if *run {
		var router *http.ServeMux = route.CreateRouter()
		fmt.Println("[*] Service is running in the port " + strconv.Itoa(*port))
		http.ListenAndServe(":9999", router)
	}
}
