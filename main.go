package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/o-t-k-t/sql_boiler_exercise/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "dbname=feeder_development user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)

	// 削除
	if _, err := models.Users().DeleteAll(ctx, db); err != nil {
		panic(err)
	}

	// 追加
	user := models.User{
		Name: "foo bar",
	}
	if err := user.Insert(ctx, db, boil.Infer()); err != nil {
		panic(err)
	}

	// 取得
	users, err := models.Users().All(ctx, db)

	for _, user := range users {
		fmt.Printf("Hello %#v", user)
	}
}
