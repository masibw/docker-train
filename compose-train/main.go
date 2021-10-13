package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func getDB() (*sqlx.DB, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT")
	mysqlConfig.User = os.Getenv("MYSQL_USER")
	mysqlConfig.Passwd = os.Getenv("MYSQL_PASSWORD")
	mysqlConfig.DBName = os.Getenv("MYSQL_DATABASE")
	mysqlConfig.Params = map[string]string{
		"time_zone": "'+00:00'",
	}
	mysqlConfig.ParseTime = true

	return sqlx.Open("mysql", mysqlConfig.FormatDSN())
}

var db *sqlx.DB

func main() {
	fmt.Println("now starting")
	var err error
	db, err = getDB()
	log.Print(err)
	defer db.Close()

	// DBが起動するまで待つ
	for {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Print(err)
		fmt.Println("waiting")
		time.Sleep(time.Second * 1)
	}

	http.HandleFunc("/", handler)

	fmt.Println("start")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
	}
	res := "users\n"
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
		}
		res += fmt.Sprintf("id: %d, name: %s\n", id, name)
	}
	fmt.Fprint(w, res)
}
