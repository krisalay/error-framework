package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/krisalay/error-framework/errorframework/framework"
)

func main() {

	conn, err := pgx.Connect(context.Background(),
		"postgres://user:pass@localhost:5432/db")

	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(),
		"INSERT INTO users(id) VALUES('1')")

	if err != nil {

		appErr := framework.DB(err)

		fmt.Println(appErr.Code)
	}
}
