package database

import (
	"fmt"
	"need-some-todo-overview/backend/internal/model"
	"net/url"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(databasePath string) (*gorm.DB, error) {
	query := url.Values{}
	query.Set("_foreign_keys", "on")
	query.Set("_busy_timeout", "5000")

	dsn := fmt.Sprintf("%s?%s", databasePath, query.Encode())

	db, err := gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("open SQLite database: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Project{},
		&model.Todo{},
		&model.Subtask{},
	)
	if err != nil {
		return fmt.Errorf("run database migrations: %w", err)
	}

	return nil
}
