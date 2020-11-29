package main

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/MonMon201/l3d6/server/db"
	forums "github.com/MonMon201/l3d6/server/forums"
	tools "github.com/MonMon201/l3d6/server/tools"
)

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "Forums",
		User:       "postgres",
		User:       "postgres",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	http.HandleFunc("/forums", func(w http.ResponseWriter, r *http.Request) {
		store := forums.NewForumStore(db)
		if r.Method == "GET" {
			store.ListChannels()
			if err != nil {
				log.Printf("Error making query to the db: %s", err)
				tools.WriteJsonInternalError(rw)
				return
			}
			tools.WriteJsonOk(rw, res)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
