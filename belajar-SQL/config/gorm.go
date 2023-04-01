package config

import (
	"belajar-SQL/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Gorm struct {
	// db configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string

	// db connection
	DB *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	GORM *GormDb
)

func InitGorm() error {
	GORM = new(GormDb)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	// connect to database
	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

// func connect
func (p *Gorm) OpenConnection() error {
	// init dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)
	// passing parameter di postgres.Open
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "db_go_sql.",
			// Untuk membuat nama menjadi singular
			SingularTable: false,
		},
	})
	if err != nil {
		return err
	}
	//assign connecion yang terbuka
	p.DB = dbConnection

	// test connection
	// automigrate dengan argumen model yang ingin dimigrasi (dapat lebih dari satu)
	err = p.DB.Debug().AutoMigrate(model.Book{})
	if err != nil {
		return err
	}

	fmt.Println("Succesfully connect to database")

	return nil
}
