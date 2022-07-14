package repositories

import (
	"fmt"
	"path"
	"testing"
	"time"
)

func TestDBRepository_KeyValues(t *testing.T) {
	tmp := path.Join(t.TempDir(), "db_repository_key_value_test.db")
	conn := fmt.Sprintf("sqlite://%s", tmp)

	t.Run("Test Key Values", func(t *testing.T) {
		repo := RepositoryFactory.GetRepository(conn)
		if repo == nil {
			t.Error("Repository not created")
		}

		if err := repo.SetKeyValue("key", "value", time.Duration(0)); err != nil {
			t.Errorf("Error setting key value: %v", err)
		}
		if err := repo.SetKeyValue("expired_key", "value", time.Microsecond); err != nil {
			t.Errorf("Error setting key value: %v", err)
		}

		if value := repo.GetKeyValue("key", "default"); value != "value" {
			t.Errorf("Expected value to be 'value', got '%s'", value)
		}

		if value := repo.GetKeyValue("expired_key", "default"); value != "default" {
			t.Errorf("Expected value to be 'default', got '%s'", value)
		}
	})

}
