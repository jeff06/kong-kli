package consumer_groups_consumers_sub

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"kong-kli/util"
	"net/url"
)

var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "List all consumers in a consumer group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		controlPlaneId, _ := cmd.Flags().GetString("controlPlaneId")
		consumerGroupId, _ := cmd.Flags().GetString("consumerGroupId")
		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumer_groups/%s/consumers", server, controlPlaneId, consumerGroupId)
		baseUrl, _ := url.Parse(prebuildUrl)

		queryParam := make(map[string]interface{})
		queryParam["size"], _ = cmd.Flags().GetInt("size")
		queryParam["offset"], _ = cmd.Flags().GetString("offset")
		queryParam["tags"], _ = cmd.Flags().GetString("tags")

		resp := util.Execute(cmd, baseUrl, "GET", queryParam, nil)

		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(body))
	},
}

func init() {
	CmdGet.Flags().StringP("consumerGroupId", "u", "", "ID of the Consumer Group to lookup")
	CmdGet.Flags().IntP("size", "z", 100, "Number of resources to be returned.")
	CmdGet.Flags().StringP("offset", "o", "", "Offset from which to return the next set of resources. Use the value of the ‘offset’ field from the response of a list operation as input here to paginate through all the resources")
	CmdGet.Flags().StringP("tags", "g", "", "A list of tags to filter the list of resources on. Multiple tags can be concatenated using ‘,’ to mean AND or using ‘/’ to mean OR.")
	CmdGet.MarkFlagsMutuallyExclusive("consumerGroupId", "size")
	CmdGet.MarkFlagsMutuallyExclusive("consumerGroupId", "offset")
	CmdGet.MarkFlagsMutuallyExclusive("consumerGroupId", "tags")
}
