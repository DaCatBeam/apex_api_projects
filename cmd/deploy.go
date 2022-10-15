package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
    Use:   "deploy",
    Short: "Deploy a project",
    Long:  "Deploy a project",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("SUCCESSFULLY COMPLETED deploy")
    },
}

func init() {
    rootCmd.AddCommand(deployCmd)
}
