package cli

import (
	"fmt"
	"os"
	"os/signal"
	"errors"
	"net/http"
	"context"
	"time"

	"github.com/jean0t/cool-bank/internal/route"
	"github.com/jean0t/cool-bank/internal/database"

	"github.com/spf13/cobra"
)

var (
	PublicKey, PrivateKey, port, databaseName string
	ctx context.Context
	cancel context.CancelFunc
	router *http.ServeMux
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Aliases: []string{"s"},
	Short: "Start the API",
	PreRunE: func (cmd *cobra.Command, args []string) error {

		// public_key.pem and private_key.pem must be provided
		if publicKey == "" {
			return errors.New("public_key.pem is mandatory. Use --public-key or -k to provide it.")
		}
		if privateKey == "" {
			return errors.New("private_key.pem is mandatory. Use --private-key or -i to provide it.")
		}

		router = route.CreateRouter(PublicKey, PrivateKey)
		db, err = database.ConnectDB(databaseName)
		return nil

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var server *http.Server = &http.Server{
			Addr: fmt.Sprintf(":%s", port), 
			Handler: route.LoggingMiddleware(router), 
			}

		go func() {
			fmt.Println("[*] Server Started")
			fmt.Printf("[*] Access the server with the url: http://localhost:%s/\n", port)
			if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("[!] HTTP server error: %v\n", err)
			}
		}()

		
		var signalInt chan os.Signal = make(chan os.Signal, 1)
		signal.Notify(signalInt)

		<-signalInt
		fmt.Println("[*] Shutting down the server")
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err = server.Shutdown(ctx); err != nil {
			return err
		}

		fmt.Println("[*] Server Closed Cleanly")
		return nil
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "9696", "Port that the server will use to listen to requests")
	serveCmd.Flags().StringVarP(&PublicKey, "public-key", "k", "", "Path to the public_key.pem file (mandatory)")
	serveCmd.Flags().StringVarP(&PrivateKey, "private-key", "i", "", "Path to the private_key.pem file (mandatory)")
	serveCmd.Flags().StringVarP(&databaseName, "database-name", "d", "bank.db", "Name of the database file")

	serveCmd.MarkFlagRequired("public-key")
	serveCmd.MarkFlagRequired("private-key")

	rootCmd.AddCommand(serveCmd)
}
