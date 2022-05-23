package infra

import (
	"fmt"
	"strings"
)

type ConnectionString struct {
	Schema         string
	ConnectionData string
}

func CreateConnectionString(connectionString string) (*ConnectionString, error) {
	words := strings.SplitN(connectionString, ":", 2)
	if len(words) != 2 {
		return nil, fmt.Errorf("expected connection string as 'schema:connection_data' - got '%s'", connectionString)
	}
	return &ConnectionString{
		Schema:         strings.ToLower(words[0]),
		ConnectionData: words[1],
	}, nil
}

func (cs *ConnectionString) String() string {
	return fmt.Sprintf("%s:%s", cs.Schema, cs.ConnectionData)
}
