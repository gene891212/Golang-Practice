package lib

import (
	"database/sql"
	"fmt"
	"gin-api-server/stru"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	env, _         = godotenv.Read(".env")
	setting string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		env["user"],
		env["password"],
		env["ip"],
		env["database"],
	)
)

func DataFromDB() []stru.User {
	db, err := sql.Open("mysql", setting)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	var (
		id           int
		userId, name string
		allUser      []stru.User
	)

	for rows.Next() {
		err = rows.Scan(&id, &userId, &name)
		if err != nil {
			log.Fatal(err)
		}
		allUser = append(
			allUser,
			stru.User{
				ID:     id,
				UserId: userId,
				Name:   name,
			},
		)
	}
	return allUser
}
