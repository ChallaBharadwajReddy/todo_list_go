package database

import (
	"database/sql"
	"fmt"

	// "os"
	// "strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
)

var Db *sql.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {

	err := godotenv.Load() //by default, it is .env so we don't have to write
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	//we read our .env file
	// host := os.Getenv("HOST")
	// port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	// user := os.Getenv("USER")
	// dbname := os.Getenv("DB_NAME")
	// pass := os.Getenv("PASSWORD")

	// set up postgres sql to open it.
	psqlSetup := "host=localhost port=5432 user=challabharadwajreddy dbname=todo_list sslmode=disable"
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		Db = db
		if err = db.Ping(); err != nil {
			fmt.Println("Error pinging DB:", err) // You should see SSL error here if it's still misconfigured
			return
		}
		// fmt.Println("Postgres connection string:", psqlSetup)
		fmt.Println("Successfully connected to database!")
	}
}
