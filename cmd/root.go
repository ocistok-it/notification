package cmd

import (
	"github.com/ocistok-it/notification/cmd/consumer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Short: "Notification System",
	}
)

func Execute() {
	registerConsumerCommand()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute command")
	}
}

func registerConsumerCommand() {
	consumer.ServeCommand().PersistentFlags().StringP("env", "e", "", "set environment to use for service")
	err := consumer.ServeCommand().MarkPersistentFlagRequired("env")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to mark persistent flag required")
	}

	rootCmd.AddCommand(consumer.ServeCommand())
}
