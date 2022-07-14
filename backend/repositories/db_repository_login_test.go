package repositories

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
)

func TestDBRepository_Logins(t *testing.T) {
	tmp := path.Join(t.TempDir(), "db_repository_logins_test.db")
	conn := fmt.Sprintf("sqlite://%s", tmp)
	t.Run("Test Key Values", func(t *testing.T) {
		repo := RepositoryFactory.GetRepository(conn)
		if repo == nil {
			t.Error("Repository not created")
		}
		loginResponse := responses.MappaLoginResponse{
			ID:      "ABCD",
			TTL:     3600,
			Created: time.Now(),
			Userid:  1234,
		}
		if err := repo.SetLogin("username", "password", loginResponse, time.Now()); err != nil {
			t.Errorf("Error setting login: %v", err)
		}
		if err := repo.SetLogin("user2", "password2", loginResponse, time.Now()); err != nil {
			t.Errorf("Error setting login: %v", err)
		}
		logins, err := repo.GetLogins()
		if err != nil {
			t.Errorf("Error getting logins: %v", err)
		}
		if len(logins) != 2 {
			t.Errorf("Expected 1 login, got %d", len(logins))
		}
		if err = repo.DeleteLogin("username"); err != nil {
			t.Errorf("Error deleting login: %v", err)
		}
		if logins, err = repo.GetLogins(); err != nil || len(logins) != 1 {
			t.Errorf("Expected 0 logins, got: %v %v", logins, err)
		}
	})

}
