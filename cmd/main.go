package main

import (
	"log"

	"avro/cmd/consumer"
	"avro/cmd/producer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "kafka",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(producer.Command())
	rootCmd.AddCommand(consumer.Command())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
