package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"				
)	


var rootCmd = &cobra.Command{
	Use:   "tinydocker",	
	Short: "a simple docker CLI",		
	Long:  "tinydocker is a simple docker CLI that is easy to use and understand",
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)				
	}	
}