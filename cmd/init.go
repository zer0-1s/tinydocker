package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize container environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init container")

		if len(args) < 1 {
			fmt.Println("Missing container command in init")
			os.Exit(1)
		}

		// 1. 先把整个挂载空间设置为私有
		if err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
			fmt.Println("Set mount propagation private failed:", err)
			os.Exit(1)
		}

		// 2. 挂载新的 proc
		if err := syscall.Mount("proc", "/proc", "proc", uintptr(syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV), ""); err != nil {
			fmt.Println("Mount proc error:", err)
			os.Exit(1)
		}

		// 执行用户传入的命令，比如 /bin/sh
		commandPath, err := exec.LookPath(args[0])
		if err != nil {
			fmt.Println("Exec look path error:", err)
			os.Exit(1)
		}

		if err := syscall.Exec(commandPath, args, os.Environ()); err != nil {
			fmt.Println("Exec error:", err)
			os.Exit(1)
		}
	},
}
