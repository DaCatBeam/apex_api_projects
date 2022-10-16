package webservers

import (
    "strings"
)

// STRATEGY PATTERN INTERFACE
type databaseStrategy interface {
    configureDatabase(DbConfigObject) error
    getStrategy() string
}

type MongoDbStrategy struct {
    ConnectionString string
    DbConfigObject DbConfigObject
}

func (db *MongoDbStrategy) configureDatabase(config_object DbConfigObject) error {
    db.DbConfigObject = config_object

    return nil
}

func getConnectionString(config_object DbConfigObject) string {
    var conn_uri_segments []string

    conn_uri_segments = append(conn_uri_segments, config_object.Protocol, "://")

    if config_object.AuthMechanism == "User" {
        conn_uri_segments = append(conn_uri_segments, config_object.Username, ":", config_object.Password, "@")
    }

    conn_uri_segments = append(conn_uri_segments, config_object.Host, ":", config_object.Port)

    if config_object.ConnOptions != "" {
        conn_uri_segments = append(conn_uri_segments, "/?", config_object.ConnOptions)
    }

    return strings.Join(conn_uri_segments, "")
}

func (db *MongoDbStrategy) getStrategy() string {
    return "mongodb"
}
