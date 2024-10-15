package initializes

func InitMysql() {
	// 	// Initialize the mysql
	// 	m := global.Config.Database
	// 	path := fmt.Sprintf("%s/%s", m.Path, m.DbName)
	// 	fmt.Println("Đường dẫn đến cơ sở dữ liệu: ", path)
	// 	// Mở kết nối đến cơ sở dữ liệu (hoặc tạo mới nếu chưa tồn tại)
	// 	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
	// 		SkipDefaultTransaction: false,
	// 	})
	// 	checkErrorPanic(err, "InitMysql initialization failed")
	// 	global.Logger.Info("Mysql initialization succeeded")
	// 	global.Mdb = db
	// 	migrateTables()
	// }

	// func migrateTables() {
	// 	// Migrate the tables
	// 	err := global.Mdb.AutoMigrate(
	// 		&model.User{},
	// 		&model.Role{},
	// 	)
	// 	checkErrorPanic(err, "Migrate tables failed")
	// 	global.Logger.Info("Migrate tables succeeded")
}
