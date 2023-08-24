package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Musica struct {
	id        int64
	name      string
	album     string
	artist    string
	genre     string
	year      int64
	url_image string
}

func main() {

	ctx := context.Background()

	db, err := crearConexion()
	if err != nil {
		panic(err)
	}

	//leer todos los libros que hay

	err = queryMusic(ctx, db, 20)
	if err != nil {
		panic(err)
	}

	//insertar un libro
	/*
		err = añadirMusica(ctx, db, 17, "Kno", "GATEWAY", "ASTRO", "K-Pop", 2020, "https://upload.wikimedia.org/wikipedia/en/5/50/ASTRO_Gateway_EP_Cover.jpg")
		if err != nil {
			panic(err)
		}
	*/
	//Quitar un libro
	/*
		err = quitarMusica(ctx, db, 7)
		if err != nil {
			panic(err)
		}
	*/
	db.Close()
}

func crearConexion() (*sql.DB, error) {
	//usuario, contraseña, puerto y nombre de la base de datos
	conexion := "root:1234@tcp(localhost:3306)/music"
	db, err := sql.Open("mysql", conexion)
	if err != nil {
		panic(err)
	}
	//Numero maximo de conexiones
	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func queryMusic(ctx context.Context, db *sql.DB, limite int64) error {
	qry := `select  *
	from canciones limit ?; `

	rows, err := db.QueryContext(ctx, qry, limite)
	if err != nil {
		return err
	}

	music := []Musica{}

	for rows.Next() {
		b := Musica{}
		err = rows.Scan(&b.id, &b.name, &b.album, &b.artist, &b.genre, &b.year, &b.url_image)
		if err != nil {
			return err
		}
		music = append(music, b)
	}

	fmt.Println(music)
	return nil
}

func añadirMusica(ctx context.Context, db *sql.DB, id int64, name string, album string, artist string, genre string, year int64, url_image string) error {
	qryañadir := ` INSERT INTO canciones (id, name, album, artist, genre, year, url_image) VALUES (?,?,?, ?, ?, ?, ?)

`

	result, err := db.ExecContext(ctx, qryañadir, id, name, album, artist, genre, year, url_image)
	if err != nil {
		return err
	}
	idd, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("INSERT ID:", idd)
	return nil

}

func quitarMusica(ctx context.Context, db *sql.DB, id int64) error {
	num := id
	qryquitar := `DELETE FROM canciones WHERE id = ?`

	result, err := db.Exec(qryquitar, id)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar el número de filas afectadas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Se elimino el libro con el id: %d Filas afectadas %d", num, rowsAffected)
	return nil
}
