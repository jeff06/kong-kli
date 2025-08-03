package consumer

import (
	"github.com/spf13/cobra"
	"kong-kli/cmd/consumer/actions"
)

var CmdConsumer = &cobra.Command{
	Use:   "consumer",
	Short: "Consumer api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	CmdConsumer.AddCommand(actions.CmdPost)
	CmdConsumer.AddCommand(actions.CmdPut)
	CmdConsumer.AddCommand(actions.CmdDelete)
	CmdConsumer.AddCommand(actions.CmdGet)
}
