package cmd

import (
	"dBackup/auth"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录华为云",
	Long:  `用户通过AK和SK登录至华为云并将保存身份认证信息`,
	Run: func(cmd *cobra.Command, args []string) {
		ak, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("请输入AK").WithMask(" ").Show() //获取用户标准输入中的AK
		sk, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("请输入SK").WithMask(" ").Show() //获取用户标准输入中的SK
		auth.SaveAkSk(ak, sk)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
