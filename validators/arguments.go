package validators

import (
    "fmt"

    "github.com/spf13/cobra"
)

func WebserversCmdValidator(cmd *cobra.Command, args []string) error {
    fmt.Println("Args: ==> ", args)

    return nil
}
