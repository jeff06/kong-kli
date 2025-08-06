package consumer_groups_consumers

import (
	"github.com/spf13/cobra"
	"kong-kli/cmd/consumer_groups/consumer_groups_consumers/consumer_groups_consumers_sub"
)

var CmdConsumerGroupsConsumers = &cobra.Command{
	Use:   "consumers",
	Short: "Consumer api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	CmdConsumerGroupsConsumers.AddCommand(consumer_groups_consumers_sub.CmdGet)
	CmdConsumerGroupsConsumers.AddCommand(consumer_groups_consumers_sub.CmdPost)
	CmdConsumerGroupsConsumers.AddCommand(consumer_groups_consumers_sub.CmdDelete)
	CmdConsumerGroupsConsumers.AddCommand(consumer_groups_consumers_sub.CmdPut)
}
