package consumer_groups_sub

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"kong-kli/util"
	"net/url"
)

var CmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Consumer Group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		controlPlaneId, _ := cmd.Flags().GetString("controlPlaneId")
		consumerGroupId, _ := cmd.Flags().GetString("consumerGroupId")
		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumer_groups/%s", server, controlPlaneId, consumerGroupId)
		baseUrl, _ := url.Parse(prebuildUrl)

		resp := util.Execute(cmd, baseUrl, "DELETE", nil, nil)

		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(body))
	},
}

func init() {
	CmdDelete.Flags().StringP("consumerGroupId", "u", "", "ID of the Consumer Group to lookup")
}
