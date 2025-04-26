package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const usage = `tinydocker is a simple container runtime implementation.
The purpose of this project is to learn how docker works and how to write a docker by ourselves.
Enjoy it, just for fun.`

var rootCmd = &cobra.Command{
	Use:   "tinydocker",
	Short: "A simple container runtime",
	Long:  usage,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
	},
}

func Execute() {
	rootCmd.AddCommand(runCommand)
	rootCmd.AddCommand(initCommand)
	_ = rootCmd.Execute()
}