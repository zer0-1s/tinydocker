package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)
// versionCmd represents the version command
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of tinydockedr",
    Long:  `All software has versions. This is tinydockedr's`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("tinydockedr v0.1 -- HEAD")
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}
