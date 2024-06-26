package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart [PROCESS]",
	Short: "Restart a process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := getClient().RestartProcess(name)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to restart process %s", name)
		}
		fmt.Printf("Process %s restarted\n", name)
	},
}

func init() {
	processCmd.AddCommand(restartCmd)
}
