package actions

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdPost = &cobra.Command{
	Use:   "post",
	Short: "Create or update a consumer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		fmt.Println("Token:", token)
	},
}

func init() {
	CmdPost.Flags().StringP("name", "n", "", "Consumer name")
}
