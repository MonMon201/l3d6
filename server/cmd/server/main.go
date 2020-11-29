package main

import (
	"fmt"
	"log"
	"net/http"

	forums "github.com/MonMon201/l3d6/server/forums"
)

func main() {
	http.HandleFunc("/forum", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(forums.NewForum(1, "", "", []int{1, 2, 3}))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
