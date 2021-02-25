package DatabaseUtils

import (
	"database/sql"
	"fmt"
)

func CreateDatabaseConnection(serverName string, userId string, password string, databaseName string) (*sql.DB, error){
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", serverName, userId, password, databaseName)
	db,err := sql.Open("mssql", connectionString)
	return db, err
}


func PerformSelectQuery(db *sql.DB, query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return rows, nil
}



