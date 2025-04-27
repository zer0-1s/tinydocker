package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"io"	
	"strings"
	"github.com/spf13/cobra"
)

const fdIndex = 3

func readUserCommand() []string {
	// uintptr(3 ）就是指 index 为3的文件描述符，也就是传递进来的管道的另一端，至于为什么是3，具体解释如下：
	/*	因为每个进程默认都会有3个文件描述符，分别是标准输入、标准输出、标准错误。这3个是子进程一创建的时候就会默认带着的，
		前面通过ExtraFiles方式带过来的 readPipe 理所当然地就成为了第4个。
		在进程中可以通过index方式读取对应的文件，比如
		index0：标准输入
		index1：标准输出
		index2：标准错误
		index3：带过来的第一个FD，也就是readPipe
	*/
	pipe := os.NewFile(uintptr(fdIndex), "pipe")
	msg, err := io.ReadAll(pipe)
	if err != nil {
		fmt.Println("init read pipe error")
		return nil
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize container environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init container")

		// if len(args) < 1 {
		// 	fmt.Println("Missing container command in init")
		// 	os.Exit(1)
		// }

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

		// 从 pipe 中读取命令
		cmdargs := readUserCommand()
		if len(cmdargs) == 0 {
			fmt.Println("run container get user command error, cmdArray is nil")
			os.Exit(1)
		}

		// 执行用户传入的命令，比如 /bin/sh
		commandPath, err := exec.LookPath(cmdargs[0])
		if err != nil {
			fmt.Println("Exec look path error:", err)
			os.Exit(1)
		}
		// runC实现的方式之一，底层调用 execve 系统调用，作用是用新的程序替换当前进程的镜像，数据，堆栈，环境变量和参数
		// 效果：把自己的进程的变成PID为1的进程
		if err := syscall.Exec(commandPath, cmdargs, os.Environ()); err != nil {
			fmt.Println("Exec error:", err)
			os.Exit(1)
		}
	},
}
