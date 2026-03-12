package cmd

import (
	"errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"path/filepath"
)

var manCmd = &cobra.Command{
	Use:   "man",
	Short: "生成dBackup的man手册页",
	Long:  `生成dBackup命令的man手册页文件到当前工作目录`,
	Run: func(cmd *cobra.Command, args []string) {
		outDir := "doc"
		if err := os.Mkdir(outDir, 0755); // 创建man手册的输出目录
		err != nil {
			if errors.Is(err, os.ErrExist) {
				pterm.Warning.Printf("%s目录已存在，man手册已生成\n", outDir)
			} else {
				pterm.Error.Printf("创建目录失败: %v\n", err)
			}
		}
		header := &doc.GenManHeader{
			Title:   "dBackup", // 手册标题
			Section: "1",       // man章节（1表示用户命令）
		}
		if err := doc.GenManTree(rootCmd, header, outDir); err != nil { // 生成man手册页
			pterm.Error.Printf("生成man手册失败: %v\n", err)
		}
		pwd, _ := os.Getwd()                   //获取当前工作路径
		fullPath := filepath.Join(pwd, outDir) //拼接路径获得完整路径
		pterm.Success.Printf("man手册已生成到: %s\n", fullPath)
		pterm.Success.Println("提示：将生成的.1文件复制到/usr/share/man/man1/后，即可用man dBackup命令查看手册")
	},
}

func init() {
	rootCmd.AddCommand(manCmd)
}
