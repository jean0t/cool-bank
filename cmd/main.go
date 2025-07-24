package main

import (
	"os/signal"
	"os"
	"fmt"
	"github.com/jean0t/cool-bank/cmd/cli"
)



func main() {
	go func() {
		var sigChan = make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		fmt.Println("")
		os.Exit(0)
	}()
	cli.Execute()
}
