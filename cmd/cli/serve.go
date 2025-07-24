package cli

import (
	"fmt"
	"net/http"

	"github.com/jean0t/cool-bank/internal/route"

	"github.com/spf13/cobra"
)

var (
	port string
	router *http.ServeMux
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Aliases: []string{"s"},
	Short: "Start the API",
	PreRun: func (cmd *cobra.Command, args []string) {
		router = route.CreateRouter()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("[*] Server Started")
		fmt.Printf("[*] Access the server with the url: http://localhost:%s/\n", port)
		err = http.ListenAndServe(":"+port, router)
		return err
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "9696", "Port that the server will use to listen to requests")
	rootCmd.AddCommand(serveCmd)
}
