package cmd

import (
	"dBackup/vault"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"}, //设置子命令的简短别名
	Short:   "列出所有备份存储库",
	Run: func(cmd *cobra.Command, args []string) {
		vault.ListVault()
	},
}

func init() {
	vaultCmd.AddCommand(listCmd)
}
