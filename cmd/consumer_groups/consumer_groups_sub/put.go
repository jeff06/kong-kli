package consumer_groups_sub

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"kong-kli/util"
	"net/url"
)

var CmdPut = &cobra.Command{
	Use:   "put",
	Short: "Create or Update Consumer Group using ID.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		controlPlaneId, _ := cmd.Flags().GetString("controlPlaneId")
		consumerGroupId, _ := cmd.Flags().GetString("consumerGroupId")
		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumer_groups/%s", server, controlPlaneId, consumerGroupId)
		baseUrl, _ := url.Parse(prebuildUrl)

		body := make(map[string]interface{})

		bodyElement := []string{"name"}

		for _, x := range bodyElement {
			element, _ := cmd.Flags().GetString(x)
			if element != "" {
				body[x] = element
			}
		}

		resp := util.Execute(cmd, baseUrl, "PUT", nil, body)

		defer resp.Body.Close()
		bodyResp, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(bodyResp))
	},
}

func init() {
	CmdPut.Flags().StringP("name", "n", "", "")
	CmdPut.Flags().StringP("consumerGroupId", "u", "", "ID of the Consumer Group to lookup")
	CmdPut.MarkFlagRequired("name")
}
