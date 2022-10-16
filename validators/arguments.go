package validators

import (
    "fmt"

    "errors"

    "github.com/spf13/cobra"
)

// Allowed webserver-template flag choices
var (
    webserver_template_choices = []string{"webhooks"}
    database_template_choices = []string{"mongodb"}
)

const (
    webserver_template_flag = "webserver-template"
    database_template_flag = "database-template"
    default_error_template = "The webserver flag value of %+v, for flag %+v, is not valid. Allowed values are: %+v."
)

// Validator for the webserver sub command flags and values.
func WebserversCmdValidator(cmd *cobra.Command, args []string) error {
    if err := validateTemplateFlag(cmd, webserver_template_flag, webserver_template_choices); err != nil {
        return err
    }

    if err := validateTemplateFlag(cmd, database_template_flag, database_template_choices); err != nil {
        return err
    }

    return nil
}

func validateTemplateFlag(cmd *cobra.Command, flag_name string, choices []string) error {
    flag_val, err := cmd.Flags().GetString(flag_name)

    if err != nil {
        return err
    }

    if isInvalidStringFlagValue(flag_val, choices) {
        return errors.New(fmt.Sprintf(default_error_template, flag_name, flag_val, choices))
    }

    return nil
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
