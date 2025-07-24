package cli

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use: "cool_bank",
	Short: "cool_bank API",
	Long: "cool_bank Official API",
}

func Execute() {
	var err error

	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
