package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var asciiArt = `
      _ ____             _                
     | |  _ \           | |               
   __| | |_) | __ _  ___| | ___   _ _ __  
  / _' |  _ < / _` + "`" + ` |/ __| |/ / | | | '_ \
 | (_| | |_) | (_| | (__|   <| |_| | |_) |
  \__,_|____/ \__,_|\___|_|\_\\__,_| .__/ 
                                   | |    
                                   |_|    
`

var rootCmd = &cobra.Command{
	Use:     "dBackup",
	Short:   "基于华为云的备份工具",
	Long:    `一个基于华为云的备份工具，通过命令行的方式备份虚拟机、云硬盘、对象存储、快照等`,
	Version: "v0.2.1",
}

func Execute() {
	blueAsciiArt := color.BlueString(asciiArt) //将Ascii设置为蓝色
	rootCmd.SetVersionTemplate(fmt.Sprintf("%s\nVersion: {{.Version}}\n", blueAsciiArt))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
