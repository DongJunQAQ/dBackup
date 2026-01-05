package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "创建备份存储库",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	vaultCmd.AddCommand(createCmd)
	var vaultName string
	var vaultSize int32
	//添加本地flag
	createCmd.Flags().StringVarP(&vaultName, "name", "n", "", "存储库名称")
	createCmd.Flags().Int32VarP(&vaultSize, "size", "s", 0, "存储库大小")
	//设置必选flag
	err := createCmd.MarkFlagRequired("name")
	err = createCmd.MarkFlagRequired("size")
	if err != nil {
		return
	}
}
