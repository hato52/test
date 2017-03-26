package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()

	tmp, err1 := template.ParseFiles("templates/index.html")
	if err1 != nil {
		panic(err1)
	}

	err2 := tmp.Execute(w, struct {
		Title string
		id int
		List []string
	}{
		Title: "Golang Template and Database Sample",
		List: []string{"hoge","fuga", "foo"},
	})
	if err2 != nil {
		panic(err2)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	/*e.GET("/users/:id", func (c echo.Context) error {
		jsonMap := map[string] string {
			"name": "hoge",
			"id": "10",
		}
		return c.JSON(http.StatusOK, jsonMap)
	})*/

	db, err := sql.Open("mysql", "root:natori11@/todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close();

	rows, err := db.Query("SELECT * FROM todo")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
