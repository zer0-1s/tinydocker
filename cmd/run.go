package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"io"
	"strings"
	"github.com/creack/pty"
	"github.com/spf13/cobra"
)

var tty bool
var interactive bool

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run a container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Missing container command")
			return
		}
	
		parent, writePipe := NewParentProcess(tty)
	
		if tty {
			ptmx, err := pty.Start(parent)
			if err != nil {
				fmt.Println("start pty failed:", err)
				return
			}
			defer func() { _ = ptmx.Close() }()
			sendInitCommand(args, writePipe)

			go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
			_, _ = io.Copy(os.Stdout, ptmx)
		} else {
			parent.Stdin = os.Stdin
			parent.Stdout = os.Stdout
			parent.Stderr = os.Stderr
	
			if err := parent.Start(); err != nil {
				fmt.Println("Run failed:", err)
				return
			}	
			sendInitCommand(args, writePipe)
			parent.Wait()
		}
	},	
}

func init() {
	runCommand.Flags().BoolVarP(&tty, "tty", "t", false, "Enable TTY")
	runCommand.Flags().BoolVarP(&interactive, "interactive", "i", false, "Enable interactive mode")
}

func NewParentProcess(tty bool) (*exec.Cmd,*os.File) {
	// 父进程中则通过writePipe将参数写入管道，代表用户传进来的参数
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		fmt.Println("Create pipe error:", err)	
	}
	cmd := exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: 
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS 	| // 创建新的挂载namespace
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC,
		Unshareflags: syscall.CLONE_NEWNS,  // 确保执行时 unshare 成新的 mount namespace
		Setsid:  true,
		Setctty: tty,
	}
	// readPipe 是子进程的额外文件用于读取参数
	// cmd 执行时就会外带着这个文件句柄去创建子进程。
	cmd.ExtraFiles = []*os.File{readPipe} // 将管道的写入端传递给子进程
	return cmd,writePipe
}


// sendInitCommand 通过writePipe将指令发送给子进程
func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	_, _ = writePipe.WriteString(command)
	_ = writePipe.Close()
}