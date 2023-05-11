package main

/*import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
*/
/*Tener en cuenta el driver para conectarse a la BD Postgress
 */

//getConnection obtiene una conex a la BD
/*func getConnection() *sql.DB {
	dsn := "postgres://postgres:Post456789*@dbpostgres.c5if27yrrstg.us-east-1.rds.amazonaws.com:5432/gocrud?sslmode=disable" // para conectar de modo seguros añadir:?sslmode=disable
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}*/

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnection() *sql.DB {
	dsn := "postgres://postgres:Post456789*@localhost:5432/gocrud?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
