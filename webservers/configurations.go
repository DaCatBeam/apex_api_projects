package webservers

import (
    "fmt"
    "errors"

    "github.com/kelseyhightower/envconfig"
)

// type ConfigObjectEnum int

const (
    // db_config_object_enum ConfigObjectEnum = iota
    db_config_object_enum int = 0
)

type IConfigObjectFactory interface {
    NewConfigObject(string) (IConfigObject, error)
}

type IConfigObject interface {
    ConfigObjectEnum() int
}

type DbConfigObject struct {
    Host     string `envconfig:"AAP_DB_HOST" required:"true"` 
    Port     string `envconfig:"AAP_DB_PORT" required:"true"`
    Username string `envconfig:"AAP_DB_USERNAME" required:"true"`
    Password string `envconfig:"AAP_DB_PASSWORD" required:"true"`
    Protocol string `envconfig:"AAP_DB_PROTOCOL" required:"true"`
    Name     string `envconfig:"AAP_DB_NAME" required:"true"`
    Vendor   string `envconfig:"AAP_DB_VENDOR" default:""`
}

func (c *DbConfigObject) ConfigObjectEnum() int {
    return db_config_object_enum
}

type ConfigObjectFactory struct {}

func (c *ConfigObjectFactory) NewConfigObject(conf_object_type string) (IConfigObject, error) {
    var conf_object IConfigObject

    switch conf_object_type {
    case "db_config_object":
        conf_object = &DbConfigObject{}
    default:
        return nil, errors.New(fmt.Sprintf("Unknown config object type %s", conf_object_type))
    }

    err := envconfig.Process("", conf_object)

    return conf_object, err
}
