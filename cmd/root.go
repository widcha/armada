package cmd

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/widcha/armada/configs"
)

var rootCmd = &cobra.Command{
	Use:   "Armada",
	Short: "Welcome to Armada",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Armada")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(workerCmd)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	configs.Load()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
