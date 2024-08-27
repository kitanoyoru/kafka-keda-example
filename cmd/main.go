package main

import (
	"log"

	"github.com/kitanoyoru/kafka-keda-example/cmd/consumer"
	"github.com/kitanoyoru/kafka-keda-example/cmd/producer"

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
