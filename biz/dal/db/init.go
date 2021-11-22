package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"xyz/dbettkk/btc/biz/container"
)

type DB struct {
	BtcDB *gorm.DB
}

func NewDefaultDB() *DB {
	return &DB{
		BtcDB: GetConnect(),
	}
}

func GetConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("btc.db"), &gorm.Config{})
	if err != nil {
		log.Println("db open error")
		return nil
	}
	return db
}

func Init() {
	db := container.C.DB

	// Migrate the schema
	//err = db.AutoMigrate(&TokenTable{})
	//if err != nil {
	//	log.Println(err)
	//}

	//// 插入内容
	db.Create(&TokenTable{Token: "test token"})
	db.Create(&TokenTable{Id: 2, Token: "test token2"})
	//
	//// 读取内容
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// 更新操作： 更新单个字段
	//db.Model(&product).Update("Price", 2000)
	//
	//// 更新操作： 更新多个字段
	//db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})
	//
	//// 删除操作：
	//db.Delete(&product, 1)
}
