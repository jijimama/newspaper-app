package database

import (
    "context"
    "errors"
    "fmt"

    "github.com/jijimama/newspaper-app/ent"
    _ "github.com/lib/pq"
)

const (
    InstancePostgres int = iota
)

var (
    errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

func NewDatabaseSQLFactory(instance int) (*ent.Client, error) {
    switch instance {
    case InstancePostgres:
        config := NewConfigPostgres()
        dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
            config.Host, config.Port, config.User, config.Name, config.Password)
        client, err := ent.Open("postgres", dsn)
        if err != nil {
            return nil, err
        }
        if err := client.Schema.Create(context.Background()); err != nil {
            return nil, err
        }
        return client, nil
    default:
        return nil, errInvalidSQLDatabaseInstance
    }
}
