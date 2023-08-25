package conexion

import (
	"conexionMysql/api"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*func main() {

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
		err = añadirMusica(ctx, db, 17, "Kno", "GATEWAY", "ASTRO", "K-Pop", 2020, "https://upload.wikimedia.org/wikipedia/en/5/50/ASTRO_Gateway_EP_Cover.jpg")
		if err != nil {
			panic(err)
		}
	//Quitar un libro
		err = quitarMusica(ctx, db, 7)
		if err != nil {
			panic(err)
		}
	db.Close()
} */

func CrearConexion() (*sql.DB, error) {
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

func QueryMusic(ctx context.Context, db *sql.DB, limite int64) error {
	qry := `select  * from canciones limit ?;`

	rows, err := db.QueryContext(ctx, qry, limite)
	if err != nil {
		return err
	}

	cancion := []api.Cancion{}

	for rows.Next() {
		b := api.Cancion{}
		err = rows.Scan(&b.ID, &b.Name, &b.Album, &b.Artist, &b.Genre, &b.Year, &b.Url_image)
		if err != nil {
			return err
		}
		cancion = append(cancion, b)
	}

	fmt.Println(cancion)
	return nil
}

func addMusica(ctx context.Context, db *sql.DB, id int64, name string, album string, artist string, genre string, year int64, url_image string) error {
	qryañadir := ` INSERT INTO canciones (id, name, album, artist, genre, year, url_image) VALUES (?,?,?, ?, ?, ?, ?)`

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

func deleteMusica(ctx context.Context, db *sql.DB, id int64) error {
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
