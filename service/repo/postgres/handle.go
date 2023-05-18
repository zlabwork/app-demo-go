package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var conn *gorm.DB
var err error

// Connect
// @docs https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func connect(dsn string) (*gorm.DB, error) {

	opts := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: os.Getenv("TABLE_PREFIX"),
			// SingularTable: true,
		},
	}

	// https://github.com/go-gorm/postgres
	conn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), opts)

	db, err := conn.DB()

	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)

	// 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)

	return conn, err
}

func getHandle() (*gorm.DB, error) {

	if conn != nil {
		return conn, nil
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	name := os.Getenv("POSTGRES_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, pass, name, port)
	return connect(dsn)
}
