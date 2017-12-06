package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var db *leveldb.DB

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/**/*.html")),
	}

	db, _ = leveldb.OpenFile("db", nil)

	defer db.Close()

	e := echo.New()
	e.Renderer = t
	e.GET("/", index)
	e.GET("/doener", doener)
	e.GET("/doenerbox", doenerbox)
	e.GET("/tuerkischepizza", tuerkischepizza)
	e.GET("/order", orders)
	e.Static("/images/*", "images")
	e.Static("/css/*", "css")
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func orders(c echo.Context) error {

	var doener []Doener

	iter := db.NewIterator(util.BytesPrefix([]byte(time.Now().Format("2006-01-02"))), nil)
	for iter.Next() {
		var d Doener
		json.Unmarshal(iter.Value(), &d)

		doener = append(doener, d)

	}
	iter.Release()

	return c.Render(http.StatusOK, "orders.html", doener)
}
