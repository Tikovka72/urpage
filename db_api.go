package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	UserId     int
	Username   string
	Password   string
	Email      string
	CreateDate time.Time
	ImagePath  string
	Links      []string
}

func connect(username string, password string, dbname string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://"+username+":"+password+"@localhost:5432/"+dbname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func getUserViaId(userId int) *User {
	user := User{}

	var image *string
	var links *string

	err := conn.QueryRow(context.Background(), "SELECT * from user_info where user_id=$1", userId).Scan(
		&user.UserId,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreateDate,
		&image,
		&links)
	if err != nil {
		return &User{}
	}

	if image != nil {
		user.ImagePath = userImages + *image
	} else {
		user.ImagePath = userImages + "test.jpeg"
	}

	if links != nil {
		linksLst := strings.Split(*links, " ")
		user.Links = linksLst
	}
	return &user
}

func addUser(username string, password string, email string, imagePath string, links string) int {
	userId := -1
	err := conn.QueryRow(context.Background(),
		"INSERT INTO user_info (Username, Password, Email, create_date, image_path, Links)"+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id", username, password, email, time.Now(), imagePath, links).Scan(&userId)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	return userId
}
