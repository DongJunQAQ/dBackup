package cmd

import (
	"dBackup/vault"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "创建备份存储库",
	RunE: func(cmd *cobra.Command, args []string) error { //规范化的错误处理
		//根据flag名称获取命令行中flag对应的值
		vaultName, _ := cmd.Flags().GetString("name")
		vaultSize, _ := cmd.Flags().GetInt32("size")
		return vault.CreateVault(vaultName, vaultSize)
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
	_ = createCmd.MarkFlagRequired("name")
	_ = createCmd.MarkFlagRequired("size")
}
