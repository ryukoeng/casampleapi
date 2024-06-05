package migrations

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func AlbumsMigrate(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'albums'").Scan(&count)
	if err != nil {
		panic(err)
	}

	// users テーブルが存在する場合は処理終了
	if count >= 1 {
		return
	}

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE albums (
		album_id serial PRIMARY KEY,
		 album_title VARCHAR(50),
		 album_artist VARCHAR(50),
		 album_price INTEGER
		 );`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		INSERT INTO
			albums (album_title, album_artist, album_price)
		VALUES
			('Blue Train','John Coltrane',55),
			('Jeru','Gerry Mulligan',22),
			('Sarah Vaughan and Clifford Brown','Sarah Vaughan',10)
	`)

	if err != nil {
		panic(err)
	}
}
