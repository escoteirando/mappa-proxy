package repositories

import (
	"os"
	"testing"
	"time"

	"github.com/guionardo/mappa_proxy/mappa/domain"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	os.Remove("testing.sqlite")
}

func TestCreateSQLiteRepository(t *testing.T) {
	t.Run("Create repo", func(t *testing.T) {
		repo, err := CreateSQLiteRepository("testing.sqlite")
		if err != nil || repo == nil {
			t.Errorf("CreateSQLiteRepository() error = %v", err)
		}
		loginResponse := domain.LoginResponse{
			ID:      "a1",
			TTL:     10,
			Created: time.Now(),
			Userid:  1,
		}
		repo.SetLogin("guionardo", "1234", loginResponse, time.Date(2021, 10, 27, 0, 0, 0, 0, time.Local))
		repo.SetLogin("teste2", "abcd", loginResponse, time.Date(2021, 10, 27, 1, 0, 0, 0, time.Local))
		user, _ := repo.GetLastLogin()
		if user != "teste2" {
			t.Errorf("Expected last login to user 'teste2'")
		}
		err = repo.SaveData()
		if err != nil {
			t.Errorf("Unespected error %v", err)
		}

		repo2, err := CreateSQLiteRepository("testing.sqlite")
		if err != nil || repo2 == nil {
			t.Errorf("CreateSQLiteRepository() error = %v", err)
		}
		user, _ = repo2.GetLastLogin()
		if user != "teste2" {
			t.Errorf("Expected last login to user 'teste2'")
		}
	})
}
