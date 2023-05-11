package main

import (
	"errors"
	"time"
)

//Estuaiante estructura del estudiante
type Estudiante struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Crear registro un estudiante en la BD
func Crear(e Estudiante) error {
	q := `INSERT INTO 
			estudiantes (name, age, active)
			VALUES ($1, $2, $3)` //en sql es con "?" postgres $1
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q) //statement, sentecia preparada permite que no envien injection sql
	if err != nil {
		return err //Retorno error
	}
	defer stmt.Close() // No hubo error cierro conex

	r, err := stmt.Exec(e.Name, e.Age, e.Active) //Ejecuto sentencia eniando las variables a afectar

	if err != nil {
		return err
	}

	i, _ := r.RowsAffected() //Devuelve la cantidad de filas afectadas por la sentencia
	if i != 1 {
		return errors.New("Se esperaba 1 fila afectada")
	}
	return nil
}
