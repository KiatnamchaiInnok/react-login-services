package databases

import (
	"database/sql"
	"log"

	"backend/configs"
	"backend/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDBConnection(cfg *configs.Configs) (*gorm.DB, error) {
	mysqlUrl, err := utils.ConnectionUrlBuilder("mysql", cfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		log.Printf("error, can't open database connection, %s", err.Error())
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		// Logger: &SqlLogger{},
	})
	if err != nil {
		defer sqlDB.Close()
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("MySQL database has been connected 🐬")
	return db, nil
}
