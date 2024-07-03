package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadDatabases() *gorm.DB {
	return connectGormDB(connectDB())
}

func connectGormDB(db *sql.DB) *gorm.DB {
	config := mysql.Config{
		Conn: db,
	}
	dialector := mysql.New(config)
	gormDB, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
		return nil
	}

	return gormDB
}

func connectDB() *sql.DB {
	_driver := viper.GetString("database.driver")
	_user := viper.GetString("database.user")
	_pass := viper.GetString("database.pass")
	_host := viper.GetString("database.host")
	_port := viper.GetString("database.port")
	_scheme := viper.GetString("database.scheme")

	db, err := sql.Open(_driver, _user+":"+_pass+"@tcp("+_host+":"+_port+")/"+_scheme)

	if err != nil {
		log.Fatal("Error ao conectar ao banco de dados: ", err)
	}

	return db
}
