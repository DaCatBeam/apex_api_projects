package validators

import (
    "fmt"

    "errors"

    "github.com/spf13/cobra"
)

// Allowed webserver-template flag choices
var (
    webserver_template_values = []string{"webhooks"}
)

const (
    webserver_template_flag = "webserver-template"
    validator_default_error_template = "The webserver flag value of %+v, for flag %+v, is not valid. Allowed values are: %v."
)

// Validator for the webserver sub command flags and values.
func WebserversCmdValidator(cmd *cobra.Command, args []string) error {
    webserver_template, err := cmd.Flags().GetString(webserver_template_flag);

    if err != nil {
        return err
    }

    switch {
        case isInvalidStringFlagValue(webserver_template, webserver_template_values):
            return errors.New(fmt.Sprintf(validator_default_error_template, webserver_template_flag, webserver_template, webserver_template_values))
        default:
            return nil
    }
}

// When the given string value is invalid, this function returns true.
func isInvalidStringFlagValue(flag_value string, allowed_values []string) bool {
    for _, v := range allowed_values {
        if flag_value == v {
            return false
        }
    }

    return true
}
