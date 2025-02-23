package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var Db *gorm.DB

func DbConnection(){
 godotenv.Load(".env")

	DNS := os.Getenv("DSN")

	var err error

	Db,err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("DB Connected")
	}

}
