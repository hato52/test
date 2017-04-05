package main

import (
	"io"
	"database/sql"
	"net/http"
	"html/template"
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetList() map[int]string {
	todo := make(map[int]string)

	db, err := sql.Open("mysql", "root:natori11@/todo")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM todo")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var memo string
		if err := rows.Scan(&id, &memo); err != nil {
			panic(err.Error())
		}
		todo[id] = memo
	}

	return todo
}

func Entry(c echo.Context) error {
	memo := c.FormValue("TodoContents")

	db, err := sql.Open("mysql", "root:natori11@/todo")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	query := "INSERT INTO todo values(null, ?)"
	if _, err := db.Exec(query, memo); err != nil {
		panic(err.Error())
	}

	return c.Render(http.StatusOK, "index", GetList())
}

func main() {

	t := &Template {
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", GetList())
	})
	e.POST("/", Entry)

	e.Logger.Fatal(e.Start(":1323"))
}
