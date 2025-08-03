package actions

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var CmdPut = &cobra.Command{
	Use:   "put",
	Short: "Update a consumer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {

}
