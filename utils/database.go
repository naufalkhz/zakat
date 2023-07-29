package utils

import (
	"github.com/naufalkhz/zakat/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetConnDB(connString string) (*gorm.DB, error) {
	// connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	conn, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DSN:               connString,
			DefaultStringSize: 255,
		},
	}, &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	if err := migrate(conn); err != nil {
		return nil, err
	}

	return conn, nil
}

func migrate(conn *gorm.DB) error {
	return conn.AutoMigrate(&models.User{}, &models.Emas{}, &models.Bank{}, &models.ZakatPenghasilan{}, &models.ZakatTabungan{}, &models.ZakatPerdagangan{}, &models.ZakatEmas{})
}
