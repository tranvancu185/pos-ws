package initializes

import (
	"database/sql"
	"path/filepath"
	"tranvancu185/vey-pos-ws/global"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

func InitMysqlC() {
	// Initialize the mysql
	m := global.Config.Database
	dbPath := m.Path
	if global.Config.Server.Mode == "production" {
		dbPath = global.Config.Database.Path
	}
	path := filepath.Join(dbPath, m.DbName)
	// Mở kết nối đến cơ sở dữ liệu (hoặc tạo mới nếu chưa tồn tại)
	db, err := sql.Open("sqlite3", path)
	checkErrorPanic(err, "InitMysql initialization failed")
	global.Logger.Info("Mysql initialization succeeded")
	global.Mdbc = db
	erro := migrateTablesC()
	checkErrorPanic(erro, "Migrate tables failed")
}

func migrateTablesC() error {
	// Migrate the tables
	goose.SetDialect("sqlite3")
	pathMigration := filepath.Join(global.Config.Database.PathMigration, "sql/schema")
	if err := goose.Up(global.Mdbc, pathMigration); err != nil {
		return err
	}
	return nil
}
