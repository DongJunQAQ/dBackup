package cmd

import (
	"dBackup/resource"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var resourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r"},
	Short:   "查看相关资源",
	Long:    `查看华为云上的各类资源如ECS、云硬盘等`,
	Run: func(cmd *cobra.Command, args []string) {
		resType, _ := cmd.Flags().GetString("type")
		switch resType {
		case "ecs":
			resource.ListEcs()
		case "evs":
			resource.ListEvs()
		default:
			pterm.Error.Println("无效命令")
			_ = cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)
	var resType string
	resourceCmd.Flags().StringVarP(&resType, "type", "t", "ecs", "要查询的资源类型")
}
