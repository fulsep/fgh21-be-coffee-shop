package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	user:= os.Getenv("DB_USER")
	pass:= os.Getenv("DB_PASS")
	host:= os.Getenv("DB_HOST")
	port:= os.Getenv("DB_PORT")
	db:= os.Getenv("DB_NAME")

	cstr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",user,pass,host,port,db)
	conn, err := pgx.Connect(context.Background(), cstr)

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
