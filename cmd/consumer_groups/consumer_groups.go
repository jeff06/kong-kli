package consumer_groups

import (
	"github.com/spf13/cobra"
	"kong-kli/cmd/consumer_groups/consumer_groups_sub"
)

var CmdConsumerGroups = &cobra.Command{
	Use:   "consumerGroups",
	Short: "Consumer api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	CmdConsumerGroups.AddCommand(consumer_groups_sub.CmdGet)
	CmdConsumerGroups.AddCommand(consumer_groups_sub.CmdPost)
	CmdConsumerGroups.AddCommand(consumer_groups_sub.CmdDelete)
	CmdConsumerGroups.AddCommand(consumer_groups_sub.CmdPut)
}
