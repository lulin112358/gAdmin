package model

// 数据迁移
func migration()  {
	DB.AutoMigrate(&User{})
}