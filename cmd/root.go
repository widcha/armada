package cmd

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/widcha/armada/configs"
)

var rootCmd = &cobra.Command{
	Use:   "Project X",
	Short: "Welcome to Project X",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Project X")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

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
