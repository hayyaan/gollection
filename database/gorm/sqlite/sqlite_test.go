package sqlite

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/gollection/gollection/database/gorm"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	c := gorm.Config{Database: "database.sqlite"}
	db, err := New(logger, c)

	assert.Nil(t, err)
	assert.NotNil(t, db)
}
