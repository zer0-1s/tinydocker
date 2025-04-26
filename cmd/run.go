package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"io"
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
	
		c := NewParentProcess(tty, args)
	
		if tty {
			ptmx, err := pty.Start(c)
			if err != nil {
				fmt.Println("start pty failed:", err)
				return
			}
			defer func() { _ = ptmx.Close() }()
			go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
			_, _ = io.Copy(os.Stdout, ptmx)
		} else {
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
	
			if err := c.Start(); err != nil {
				fmt.Println("Run failed:", err)
				return
			}
			c.Wait()
		}
	},	
}

func init() {
	runCommand.Flags().BoolVarP(&tty, "tty", "t", false, "Enable TTY")
	runCommand.Flags().BoolVarP(&interactive, "interactive", "i", false, "Enable interactive mode")
}

func NewParentProcess(tty bool, args []string) *exec.Cmd {
	cmdArgs := append([]string{"init"}, args...) // 构建：/proc/self/exe init /bin/bash
	cmd := exec.Command("/proc/self/exe", cmdArgs...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: 
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC,
		Setsid:  true,
		Setctty: tty,
	}
	return cmd
}
