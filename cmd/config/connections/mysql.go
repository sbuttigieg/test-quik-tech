package connections

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(c *wallet.Config, log *logrus.Logger) (*sql.DB, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPwd := os.Getenv("MYSQL_PASSWORD")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPwd, dbPort, dbName)

	var db *gorm.DB
	var dbOK bool
	var err error

	time.Sleep(c.StoreTimeout) // time for mysql to load

	for i := 0; i <= 3; i++ {
		if !dbOK {
			db, err = gorm.Open(mysql.Open(dsn))
			if err != nil {
				log.Info(fmt.Sprintf("DB load trial no. %v: ", i+1), err)
				time.Sleep(c.StoreTimeout) // time for retries

				continue
			}
		}

		dbOK = true
		break
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	return sqlDB, nil
}
