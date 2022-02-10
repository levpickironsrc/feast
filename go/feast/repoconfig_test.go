package feast

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewRepoConfig(t *testing.T) {
	dir, err := os.MkdirTemp("", "feature_repo_*")
	assert.Nil(t, err)
	defer func() {
		assert.Nil(t, os.RemoveAll(dir))
	}()
	filePath := filepath.Join(dir, "feature_store.yaml")
	data := []byte(`
project: feature_repo
registry: data/registry.db
provider: local
online_store:
    type: redis
    connection_string: "localhost:6379"
`)
	err = os.WriteFile(filePath, data, 0666)
	assert.Nil(t, err)
	config, err := NewRepoConfig(dir, "")
	assert.Nil(t, err)
	assert.Equal(t, "feature_repo", config.Project)
	assert.Equal(t, filepath.Join(dir, "data/registry.db"), config.Registry)
	assert.Equal(t, "local", config.Provider)
	assert.Equal(t, map[string]interface{}{
		"type":              "redis",
		"connection_string": "localhost:6379",
	}, config.OnlineStore)
	assert.Empty(t, config.OfflineStore)
	assert.Empty(t, config.FeatureServer)
	assert.Empty(t, config.Flags)
}
