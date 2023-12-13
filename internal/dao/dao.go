package dao

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Dao struct
type Dao struct {
	// mysql
	db  *gorm.DB
	rdb *redis.Client
}

func InitMysqlPool() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"), viper.GetString("db.charset"))
	//db, err := orm.Open("mysql", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//initSharding(db)
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(4000)
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

// New init
func New() (dao *Dao) {
	dao = &Dao{
		db:  InitMysqlPool(),
		rdb: InitRedis(),
	}
	return
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})
	return rdb
}

// Close close the resource.
func (d *Dao) Close() {
	sqlDB, _ := d.db.DB()
	sqlDB.Close()
}
