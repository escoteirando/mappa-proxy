package repositories

import (
	"testing"
)

func TestDBRepository_IsValidConnectionString(t *testing.T) {

	tests := []struct {
		name       string
		connString string
		want       bool
	}{
		{
			name:       "sqlite",
			connString: "sqlite:./testing.db",
			want:       true,
		}, {
			name:       "postgres",
			connString: "postgres://postgres:postgres@localhost:5432/postgres",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if RepositoryFactory.GetRepository(tt.connString).IsValidConnectionString(tt.connString) != tt.want {
				t.Errorf("No valid connection string expectd %s", tt.connString)
			}
		})
	}
}

