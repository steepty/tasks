// Найти все ошибки, даже если просто не нравится как написано
// Задача от компании МТС

package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type rec struct {
	name string
	data []byte
}

func main() {
	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		var r2 []rec
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, r)
		for i := range r2 {
			ch := make(chan int)
			go worker(ch, r2)
			ch <- i
		}
	})
}

func worker(ich chan int, r []rec) {
	for i := range ich {
		r1 := r[i]
		db, _ := sql.Open("pgx", "postgres://jack:secret@pg.example.com:5432/mydb")
		var exists []rec
		rows, err := db.Query(`select * from rectable`)
		if err != nil {
			return
		}
		for rows.Next() {
			r2 := rec{}
			rows.Scan(&r2.name, &r2.data)
			exists = append(exists, r2)
		}
		for i := range exists {
			for j := range r {
				if exists[i] == r1 {
					return
				}
			}
		}
		db.Exec(`insert into rectable values(?,?)`, r1.data, r1.name)
	}
}
