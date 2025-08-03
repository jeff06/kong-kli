package actions

import (
	"fmt"
	"io"
	"kong-kli/util"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var CmdPut = &cobra.Command{
	Use:   "put",
	Short: "Create or Update Consumer using ID or username.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		controlPlaneId, _ := cmd.Flags().GetString("controlPlaneId")
		consumerId, _ := cmd.Flags().GetString("consumerId")

		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumers/%s", server, controlPlaneId, consumerId)
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

		resp := util.Execute(cmd, baseUrl, "PUT", nil, body)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)
		bodyResp, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(bodyResp))
	},
}

func init() {
	CmdPut.Flags().StringP("consumerId", "m", "", "ID of the Consumer to lookup")
	CmdPut.Flags().StringP("custom_id", "o", "", "Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or username with the request.")
	CmdPut.Flags().StringP("id", "i", "", "Valid UUID. If not provided, one will be generated.")
	CmdPut.Flags().StringP("tags", "g", "", "An optional set of strings associated with the Consumer for grouping and filtering.")
	CmdPut.Flags().StringP("username", "u", "", "The unique username of the Consumer. You must send either this field or custom_id with the request.")
	CmdPut.MarkFlagsOneRequired("username", "custom_id")
}
