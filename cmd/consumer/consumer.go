package consumer

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ocistok-it/notification/bootstrap"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	serveCmd = &cobra.Command{
		Use:              "consumer",
		Short:            "run event consumer server",
		PersistentPreRun: rootPreRun,
		RunE:             run,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	env := cmd.PersistentFlags().Lookup("env").Value.String()

	if env == "local" {
		initLog()
	}

	filename := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(filename)

	if err != nil {
		log.Fatal().Err(err).Str("filename", filename).Msg("error load env")
	}
}

func initLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func run(cmd *cobra.Command, args []string) error {
	bootstrap.NewConsumer().
		RegisterHandler().
		Run()

	return nil
}

func ServeCommand() *cobra.Command {
	return serveCmd
}
