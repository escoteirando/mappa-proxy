package cache

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/escoteirando/mappa-proxy/backend/repositories"
)

func TestCreateMappaCache(t *testing.T) {
	tmp := path.Join(t.TempDir(), "db_mappa_test.db")
	conn := fmt.Sprintf("sqlite://%s", tmp)
	t.Run("Cache", func(t *testing.T) {
		repo := repositories.RepositoryFactory.GetRepository(conn)
		if repo == nil {
			t.Error("Repository not created")
		}
		cache, err := CreateMappaCache(repo)
		if err != nil {
			t.Errorf("Failed to create cache: %v", err)
		}
		cache.SetLastEventTime("test_event", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		lastEventTime := cache.GetLastEventTime("test_event")
		if lastEventTime.IsZero() {
			t.Errorf("Failed to get last event time")
		}
	})

}
