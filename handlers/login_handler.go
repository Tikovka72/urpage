package handlers

import (
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	"go-site/jwt"
	"go-site/session"
	"go-site/structs"
	"html/template"
	"log"
	"net/http"
)

func CreateLoginHandler(_ *pgx.Conn, rdb *redis.Client) {

	loginHandler := func(writer http.ResponseWriter, request *http.Request) {
		var (
			CSRFToken string
			err       error
		)

		{ // check csrf
			_, CSRFToken, err = session.CheckSessionId(writer, request, rdb)
			if err != nil {
				http.Error(writer, "что-то пошло не так...", http.StatusInternalServerError)
				return
			}
		}

		{ // check user authed
			_, err = jwt.CheckIfUserAuth(writer, request, rdb)

			if err == nil {
				http.Redirect(writer, request, "/", http.StatusSeeOther)
				return
			}
		}

		{ // generate login page
			t, err := template.ParseFiles("templates/login.html")

			if err != nil {
				log.Println(err)
			}

			err = t.Execute(writer, structs.TemplateData{"CSRF": CSRFToken})

			if err != nil {
				log.Println(err)
			}
		}
	}

	http.HandleFunc("/login", loginHandler)
}
