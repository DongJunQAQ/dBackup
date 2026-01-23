package cmd

import (
	"dBackup/backup"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"b"},
	Short:   "开始备份",
	Long:    `为存储库创建检查点`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vaultId, _ := cmd.Flags().GetString("vault")
		return backup.CreateCheckpoint(vaultId)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	var vaultId string
	backupCmd.Flags().StringVarP(&vaultId, "vault", "v", "", "存储库ID (格式为xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)")
	_ = backupCmd.MarkFlagRequired("vault")
}
