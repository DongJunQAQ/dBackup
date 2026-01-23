package cmd

import (
	"github.com/spf13/cobra"
)

var vaultCmd = &cobra.Command{
	Use:     "vault",
	Aliases: []string{"v"},
	Short:   "备份存储库相关操作",
	Long:    `备份存储库相关操作，如创建、列出、添加资源等`,
}

func init() {
	rootCmd.AddCommand(vaultCmd)
}
