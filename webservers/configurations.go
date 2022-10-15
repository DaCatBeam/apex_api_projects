package webservers

import (
    "os"
    "log"

    common_utils "github.com/DaCatBeam/apex_api_projects/common"
)

type configStrategy interface {
    configure(string) error
    getStrategy() string
}

struct EnvFileConfigStrategy struct {
    isLoaded bool
}

func (strat *EnvFileConfigStrategy) getStrategy() string {
    return "EnvFileConfigStrategy"
}

func (strat *EnvFileConfigStrategy) configure(env string) error {
    env_file_path := os.Getenv(env)

    if env_file_path == "" {
        envFilePath = ".env"
    }

    env_file := common_utils.NewEnvironmentFile(env_file_path)

    if err := env_file.Load(); err != nil {
        strat.isLoaded = false
        return err
    }

    strat.isLoaded = true

    return nil
}

