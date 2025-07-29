package cli

import (
	"fmt"
	"errors"
	"net/http"

	"github.com/jean0t/cool-bank/internal/route"
	"github.com/jean0t/cool-bank/internal/authentication"

	"github.com/spf13/cobra"
)

var (
	port string
	router *http.ServeMux
	publicKey, privateKey string
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Aliases: []string{"s"},
	Short: "Start the API",
	PreRun: func (cmd *cobra.Command, args []string) {
		// public_key.pem and private_key.pem must be provided
		if publicKey == "" {
			return errors.New("public_key.pem is mandatory. Use --public-key or -k to provide it.")
		}
		if privateKey == "" {
			return errors.New("private_key.pem is mandatory. Use --private-key or -i to provide it.")
		}

		router = route.CreateRouter(publicKey, privateKey)
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
	serveCmd.Flags().StringVarP(&publicKey, "public-key", "k", "", "Path to the public_key.pem file (mandatory)")
	serveCmd.Flags().StringVarP(&privateKey, "private-key", "i", "", "Path to the private_key.pem file (mandatory)")

	serveCmd.MarkFlagRequired("public-key")
	serveCmd.MarkFlagRequired("private-key")

	rootCmd.AddCommand(serveCmd)
}
