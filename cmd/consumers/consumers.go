package consumers

import (
	"github.com/spf13/cobra"
	"kong-kli/cmd/consumers/consumers_sub"
)

var CmdConsumer = &cobra.Command{
	Use:   "consumer",
	Short: "Consumer api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	CmdConsumer.AddCommand(consumers_sub.CmdPost)
	CmdConsumer.AddCommand(consumers_sub.CmdPut)
	CmdConsumer.AddCommand(consumers_sub.CmdDelete)
	CmdConsumer.AddCommand(consumers_sub.CmdGet)
}
