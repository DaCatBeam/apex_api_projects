package cmd

import (
    "fmt"
    "strings"
    "os"

    flag "github.com/spf13/pflag"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
    Use: "apex_api_projects",
    Short: "CLI Tool for Apex API Golang projects",
    Long: "CLI Tool for Apex API Golang projects",
    Run: func(cmd *cobra.Command, args []string) {
        // Set command values based on preference
        cmd.Flags().VisitAll(func(f *flag.Flag) {
            env_var_suffix := ""
            if strings.Contains(f.Name, "-") {
                env_var_suffix = strings.ReplaceAll(f.Name, "-", "_")
                viper.BindEnv(f.Name, env_var_suffix)
            }

            // Apply the viper config value to the flag
            // when the flag is not set and viper has a value
            if !f.Changed && viper.GetString(env_var_suffix) != "" {
                val := viper.GetString(env_var_suffix)
                cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
            }

        })
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

// Checks for initial app.env configuration file. If found it loads it into viper.
func InitConfig() {
    viper.AutomaticEnv()
    loadAppConfig(".", "app", "env")
    // loadAppConfig("./config/", "app", "env")
}

func loadAppConfig(path, name, conf_type string) {
    if path != "" {
        viper.AddConfigPath(path)
    }
    if name != "" {
        viper.SetConfigName(name)
    }
    if conf_type != "" {
        viper.SetConfigType(conf_type)
    }

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
    } else {
        fmt.Println("App Init Config File: ", viper.ConfigFileUsed(), ", was not found.")
    }
}

// Initialize Global Flag Values and bind to environment variables with viper.
func init() {
    // Add flags in root command if required
    rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Set the app log level")

    // Initialize config variables with viper
    viper.BindPFlag("log_level", rootCmd.Flags().Lookup("log-level"))
    cobra.OnInitialize(InitConfig)
}
