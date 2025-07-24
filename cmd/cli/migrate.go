package cli

import (
	"fmt"

	"github.com/jean0t/cool-bank/internal/database"

	"gorm.io/gorm"
	"github.com/spf13/cobra"
)

var (	
	db *gorm.DB
	err error
)

var migrateCmd = &cobra.Command{
	Use: "migrate [db_path]",
	Short: "Run migrations on the database",
	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var dbPath = args[0]
		
		db, err = database.ConnectDB(dbPath)
		if err != nil {
			return fmt.Errorf("Failed to open DB: %w", err)
		}

		fmt.Printf("[*] Database Connected to %s\n", dbPath)
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("[*] Running migrations")
		err = database.MigrateDB(db)
		if err != nil {
			return err
		}
		fmt.Println("[*] Migration succeeded")
		return nil
	},
}


func init() {
	rootCmd.AddCommand(migrateCmd)
}
