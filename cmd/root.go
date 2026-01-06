package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "dBackup",
	Short:   "基于华为云的备份工具",
	Long:    `一个基于华为云的备份工具，通过命令行的方式备份虚拟机、云硬盘、对象存储、快照等`,
	Version: "v0.1.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
