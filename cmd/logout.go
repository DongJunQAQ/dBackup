package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出华为云",
	Long:  `登出华为云删除dBackup配置文件并移除相关的验证信息`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		filePath := filepath.Join(homeDir, ".dbackup_config.json")
		err := os.Remove(filePath)
		if err != nil {
			pterm.Error.Printf("删除配置文件失败: %v", err)
			os.Exit(1)
		}
		pterm.Success.Println("已成功登出")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
