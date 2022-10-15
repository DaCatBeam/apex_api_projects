package cmd

import (
    "fmt"

    "github.com/DaCatBeam/apex_api_projects/validators"

    "github.com/spf13/cobra"
)

var webserversCmd = &cobra.Command{
    Use:   "webservers",
    Short: "Act on a webserver",
    Args: validators.WebserversCmdValidator,
    Long:  "Act on a webserver",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("SUCCESSFULLY COMPLETED webservers")
    },
}

func init() {
    deployCmd.AddCommand(webserversCmd)
    webserversCmd.Flags().StringP("env-file", "e", "", "Path to the environment file.")
    webserversCmd.Flags().StringP("webserver-template", "w", "", "Webserver template location")
}
