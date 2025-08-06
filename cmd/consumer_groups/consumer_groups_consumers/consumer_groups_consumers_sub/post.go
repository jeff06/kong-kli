package consumer_groups_consumers_sub

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"kong-kli/util"
	"net/url"
)

var CmdPost = &cobra.Command{
	Use:   "post",
	Short: "Create a new consumer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		controlPlaneId, _ := cmd.Flags().GetString("controlPlaneId")
		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumer_groups", server, controlPlaneId)
		baseUrl, _ := url.Parse(prebuildUrl)

		body := make(map[string]interface{})

		bodyElement := []string{"name"}

		for _, x := range bodyElement {
			element, _ := cmd.Flags().GetString(x)
			if element != "" {
				body[x] = element
			}
		}

		resp := util.Execute(cmd, baseUrl, "POST", nil, body)

		defer resp.Body.Close()
		bodyResp, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(bodyResp))
	},
}

func init() {
	CmdPost.Flags().StringP("name", "n", "", "")
	CmdPost.MarkFlagRequired("name")
}
