package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"theo.dev/hello-universe/speedtester/model"
	"theo.dev/hello-universe/speedtester/model/PerformanceResult"
)

const databaseUrl = ""

func Insert(testResult model.TestResult) (id int, err error) {

	connection, err := pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		return
	}

	sqlQuery := "insert into results (performance) values ($1) returning id"

	var returnedId int
	row := connection.QueryRow(context.Background(), sqlQuery, testResult.Performance)

	err = row.Scan(&returnedId)

	if err != nil {
		return 0, err
	}

	return returnedId, nil

}

func Select(id int) (result PerformanceResult.PerformanceResult, err error) {

	connection, err := pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		return
	}

	sqlQuery := "SELECT performance FROM results where id=$1"

	err = connection.QueryRow(context.Background(), sqlQuery, id).Scan(&result)

	if err != nil {
		return "", err
	} else {
		return result, nil
	}

}
