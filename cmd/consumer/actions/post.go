package actions

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
		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumers", server, controlPlaneId)
		baseUrl, _ := url.Parse(prebuildUrl)

		body := map[string]string{}
		body["custom_id"], _ = cmd.Flags().GetString("custom_id")
		body["id"], _ = cmd.Flags().GetString("id")
		body["tags"], _ = cmd.Flags().GetString("tags")
		body["username"], _ = cmd.Flags().GetString("username")

		resp := util.Execute(cmd, baseUrl, "POST", nil, body)

		defer resp.Body.Close()
		bodyResp, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(bodyResp))
	},
}

func init() {
	CmdPost.Flags().StringP("custom_id", "m", "", "created_at")
	CmdPost.Flags().StringP("id", "i", "", "created_at")
	CmdPost.Flags().StringP("tags", "g", "", "created_at")
	CmdPost.Flags().StringP("username", "u", "", "created_at")
}
