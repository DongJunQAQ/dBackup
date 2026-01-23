package cmd

import (
	"dBackup/resource"
	"github.com/spf13/cobra"
)

var resourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r"},
	Short:   "查看相关资源",
	Long:    `查看华为云上的各类资源如ECS、云硬盘等`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, _ := cmd.Flags().GetBool("ecs")
		if ecs {
			resource.ListEcs()
		}
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)
	var ecs bool
	resourceCmd.Flags().BoolVarP(&ecs, "ecs", "e", true, "列出ECS服务器")
}
