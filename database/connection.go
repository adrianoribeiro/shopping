package database

import (
	"bufio"
	"errors"
	"github.com/shopping/repo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect (){
	//To run mysql
	//docker run --name some-mysql1 -p3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:latest
	dsn := "root:root@tcp(127.0.0.1:3308)/marketPlaceDB?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.Product{})
	//loadDummyData()
}

func loadDummyData() {
	file, _ := OpenFile("./database/data.sql")
	defer file.Close()

	//clean table
	DB.Exec("delete from products")

	//Fill table
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		DB.Exec(scanner.Text())  // line by line
	}
}

func OpenFile(jsonFile string) (*os.File, error) {
	jsonData, err := os.Open(jsonFile)
	if err != nil {
		return nil, errors.New("Fail to open data.sql")
	}
	return jsonData, nil
}
