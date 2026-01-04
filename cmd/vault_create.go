package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "创建备份存储库",
	Long:  `创建备份存储库`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	vaultCmd.AddCommand(createCmd)

}
