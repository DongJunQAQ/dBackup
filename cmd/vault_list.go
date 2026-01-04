package cmd

import (
	"dBackup/vault"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有备份存储库",
	Long:  `列出所有备份存储库`,
	Run: func(cmd *cobra.Command, args []string) {
		vault.ListVault()
	},
}

func init() {
	vaultCmd.AddCommand(listCmd)
}
