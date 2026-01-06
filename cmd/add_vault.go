package cmd

import (
	"dBackup/vault"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "添加资源至存储库",
	RunE: func(cmd *cobra.Command, args []string) error {
		vaultId, _ := cmd.Flags().GetString("vault")
		resourcesID, _ := cmd.Flags().GetString("resources")
		resourcesType, _ := cmd.Flags().GetString("type")
		return vault.AddResources(vaultId, resourcesID, resourcesType)
	},
}

func init() {
	vaultCmd.AddCommand(addCmd)
	var vaultId string
	var resourcesID string
	var resourcesType string
	//添加本地flag
	addCmd.Flags().StringVarP(&vaultId, "vault", "v", "", "存储库ID (如ad6910b3-6c77-45bb-8651-276adce80fb8)")
	addCmd.Flags().StringVarP(&resourcesID, "resources", "r", "", "备份资源ID (如cb77b86f-6e06-4f29-96e7-ecbc012d9816)")
	addCmd.Flags().StringVarP(&resourcesType, "type", "t", "", "备份资源类型 (服务器:server, 磁盘:volume)")
	//设置必选flag
	_ = addCmd.MarkFlagRequired("vault")
	_ = addCmd.MarkFlagRequired("resources")
	_ = addCmd.MarkFlagRequired("type")
}
