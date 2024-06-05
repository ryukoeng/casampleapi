package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price`
}

func Albums(db *sql.DB) {
	http.HandleFunc("/albums/get", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM albums")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		albums := []Album{}
		for rows.Next() {
			album := Album{}
			err := rows.Scan(
				&album.ID,
				&album.Title,
				&album.Artist,
				&album.Price)
			if err != nil {
				log.Fatal(err)
			}
			albums = append(albums, album)
		}

		json.NewEncoder(w).Encode(albums)
	})

	http.HandleFunc("/albums/create", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		album := Album{}
		err = json.Unmarshal(body, &album)
		if err != nil {
			log.Fatal(err)
		}

		stmt, err := db.Prepare("INSERT INTO albums (album_id,album_title,album_artist,album_price) VALUES ($1,$2,$3,$4) RETURNING album_id,album_title,album_artist,album_price;")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		err = stmt.QueryRow(album.ID).Scan(
			&album.ID,
			&album.Title,
			&album.Artist,
			&album.Price)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(album)
	})
}
