package webservers

// STRATEGY PATTERN INTERFACE
type databaseStrategy interface {
    configureDatabase(string) error
    getStrategy() string
}

type MongoDbStrategy struct {
    ConnectionString string
}

func (db *MongoDbStrategy) configureDatabase(db_env_file string) error {
    return nil
}