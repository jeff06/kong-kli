package consumers_sub

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var CmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a consumer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {

}
