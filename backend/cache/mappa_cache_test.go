package cache

import (
	"reflect"
	"testing"

	"github.com/escoteirando/mappa-proxy/backend/repositories"
)

func TestCreateMappaCache(t *testing.T) {
	
	type args struct {
		repository repositories.IRepository
	}
	tests := []struct {
		name      string
		args      args
		wantCache *MappaCache
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCache, err := CreateMappaCache(tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMappaCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCache, tt.wantCache) {
				t.Errorf("CreateMappaCache() = %v, want %v", gotCache, tt.wantCache)
			}
		})
	}
}
