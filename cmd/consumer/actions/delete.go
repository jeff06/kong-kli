package actions

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"kong-kli/util"
	"net/url"
)

var CmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a consumer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		controlPlaneId, _ := cmd.Flags().GetString("controlPlaneId")
		consumerId, _ := cmd.Flags().GetString("consumerId")
		prebuildUrl := fmt.Sprintf("https://%s/control-planes/%s/core-entities/consumers/%s", server, controlPlaneId, consumerId)
		baseUrl, _ := url.Parse(prebuildUrl)

		resp := util.Execute(cmd, baseUrl, "DELETE", nil, nil)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(body))
	},
}

func init() {
	CmdDelete.Flags().StringP("consumerId", "m", "", "Offset from which to return the next set of resources. Use the value of the ‘offset’ field from the response of a list operation as input here to paginate through all the resources")
}
