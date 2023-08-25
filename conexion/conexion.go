package conexion

import (
	"conexionMysql/modelo"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CrearConexion() *sql.DB {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	//usuario, contraseña, puerto y nombre de la base de datos
	db, _ := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+
		"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	//Numero maximo de conexiones
	db.SetMaxOpenConns(5)
	return db
}

func QueryMusic(ctx context.Context, db *sql.DB, limite int) []modelo.Cancion {
	qry := `select  * from canciones limit ?;`

	rows, err := db.QueryContext(ctx, qry, limite)
	if err != nil {
		panic(err)
	}

	cancion := []modelo.Cancion{}

	for rows.Next() {
		b := modelo.Cancion{}
		err = rows.Scan(&b.ID, &b.Name, &b.Album, &b.Artist, &b.Genre, &b.Year, &b.Url_image)
		if err != nil {
			panic(err)
		}
		cancion = append(cancion, b)
	}

	return cancion
}

func AddMusica(ctx context.Context, db *sql.DB, id int64, name string, album string, artist string, genre string, year int64, url_image string) error {
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

func UpdateMusica(ctx context.Context, db *sql.DB, id int64, name string, album string, artist string, genre string, year int64, url_image string) error {
	query := `UPDATE canciones SET name = ?, album = ?, artist = ?, genre = ?, year = ?, url_image = ? where id = ?;`
	_, err := db.ExecContext(ctx, query, name, album, artist, genre, year, url_image, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMusica(ctx context.Context, db *sql.DB, id int64) error {
	qryquitar := `DELETE FROM canciones WHERE id = ?`

	_, err := db.Exec(qryquitar, id)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar el número de filas afectadas
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Printf("Se elimino el libro con el id: %d Filas afectadas %d", num, rowsAffected)
	return nil
}
