package consumers_sub

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"kong-kli/util"
	"net/url"
	"strings"
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

		body := make(map[string]interface{})

		bodyElement := []string{"custom_id", "id", "username"}

		for _, x := range bodyElement {
			element, _ := cmd.Flags().GetString(x)
			if element != "" {
				body[x] = element
			}
		}

		tagsStr, _ := cmd.Flags().GetString("tags")
		if tagsStr != "" {
			body["tags"] = strings.Split(tagsStr, ",")
		}

		resp := util.Execute(cmd, baseUrl, "POST", nil, body)

		defer resp.Body.Close()
		bodyResp, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(bodyResp))
	},
}

func init() {
	CmdPost.Flags().StringP("custom_id", "m", "", "Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or username with the request.")
	CmdPost.Flags().StringP("id", "i", "", "Id")
	CmdPost.Flags().StringP("tags", "g", "", "An optional set of strings associated with the Consumer for grouping and filtering.")
	CmdPost.Flags().StringP("username", "u", "", "The unique username of the Consumer. You must send either this field or custom_id with the request.")
}
