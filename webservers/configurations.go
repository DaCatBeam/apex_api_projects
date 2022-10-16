package webservers

import (
    "fmt"
    "errors"

    "github.com/spf13/viper"
)

type ConfigObjectEnum int

func (c ConfigObjectEnum) String() string {
    return [...]string{"Database", "Webserver"}[c]
}

const (
    Database ConfigObjectEnum = iota
    Webserver
)

type IConfigObjectFactory interface {
    NewConfigObject(string) (IConfigObject, error)
}

type IConfigObject interface {
    Enum() ConfigObjectEnum
}

type DbConfigObject struct {
    Host                  string `mapstructure:"AAP_DB_HOST" required:"true"` 
    Port                  string `mapstructure:"AAP_DB_PORT" required:"true"`
    Username              string `mapstructure:"AAP_DB_USERNAME" required:"true"`
    Password              string `mapstructure:"AAP_DB_PASSWORD" required:"true"`
    Protocol              string `mapstructure:"AAP_DB_PROTOCOL" required:"true"`
    Name                  string `mapstructure:"AAP_DB_NAME" required:"true"`
    Vendor                string `mapstructure:"AAP_DB_VENDOR" default:""`
    AuthMechanism         string `mapstructure:"AAP_DB_AUTH_MECHANISM" default:""`
    TlsCaFile             string `mapstructure:"AAP_DB_TLS_CA_FILE" default:""`
    TlsCertificateKeyFile string `mapstructure:"AAP_DB_TLS_CERT_KEY_FILE" default:""`
    ConnOptions           string `mapstructure:"AAP_DB_CONN_OPTS" default:""`
}

func (c *DbConfigObject) Enum() ConfigObjectEnum {
    return Database
}

type ConfigObjectFactory struct {}

func (c *ConfigObjectFactory) NewConfigObject(conf_object_type string) (IConfigObject, error) {
    LoadAppConfig("$AAP_APP_DIRECTORY/config/envs/", conf_object_type, "env")

    switch conf_object_type {
    case "mongodb":
        var conf_object DbConfigObject

        viper.Unmarshal(&conf_object)

        return &conf_object, nil
    default:
        return nil, errors.New(fmt.Sprintf("Unknown config object type %s", conf_object_type))
    }
}

func LoadAppConfig(path, name, conf_type string) {
    if path != "" {
        viper.AddConfigPath(path)
    }
    if name != "" {
        viper.SetConfigName(name)
    }
    if conf_type != "" {
        viper.SetConfigType(conf_type)
    }

    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
    } else {
        fmt.Println("App Init Config File: ", viper.ConfigFileUsed(), ", was not found.")
    }

    viper.AutomaticEnv()
}