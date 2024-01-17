// package services

// import (
// 	"auth/models"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
// 	"github.com/joho/godotenv"
// )

// //ConnectDB function: Make database connection
// func ConnectDB() *gorm.DB {

// 	//Load environmenatal variables
// 	err := godotenv.Load()

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	databaseHost := os.Getenv("DB_HOST")
// 	databaseName := os.Getenv("DB_NAME")	
// 	username := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")

// 	//Define DB connection string
// 	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

// 	//connect to db URI
// 	db, err := gorm.Open("postgres", dbURI)

// 	if err != nil {
// 		fmt.Println("error", err)
// 		panic(err)
// 	}
// 	// close db when not in use
// 	defer db.Close()

// 	// Migrate the schema
// 	db.AutoMigrate(
// 		&models.User{})

// 	fmt.Println("Successfully connected!", db)
// 	return db
// }