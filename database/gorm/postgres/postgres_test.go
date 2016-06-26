package postgres

// TODO: Figure out a way to mock gorm's postgres
//func TestNew(t *testing.T) {
//	var buf bytes.Buffer
//	logger := log.NewLogfmtLogger(&buf)
//
//	db, err := New(logger, gorm.Config{
//		Username: "postgres",
//		Password: "postgres",
//		Host:     "127.0.0.1",
//		Port:     5432,
//		Database: "postgres",
//	})
//
//	assert.Nil(t, err)
//	assert.NotNil(t, db)
//}
