package main

import (
	"encoding/json"
	"fmt"
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

type Orders struct {
	Doener          []Doener
	Doenerbox       []Doenerbox
	Tuerkischepizza []Tuerkischepizza
}

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
	e.GET("/orders", orders)
	e.Static("/images/*", "images")
	e.Static("/css/*", "css")
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func orders(c echo.Context) error {

	//var doener []Doener

	//var orders Orders

	//Holt alle Doenerboxen aus der Datenbank
	var dbox []Doenerbox
	var d []Doener
	var t []Tuerkischepizza

	iter := db.NewIterator(util.BytesPrefix([]byte("Doener Box"+time.Now().Format("2006-01-02"))), nil)
	for iter.Next() {
		var doenerbox Doenerbox
		json.Unmarshal(iter.Value(), &doenerbox)

		dbox = append(dbox, doenerbox)

	}
	iter.Release()

	//Holt alle Doener aus der Datenbank
	iter = db.NewIterator(util.BytesPrefix([]byte("Doener"+time.Now().Format("2006-01-02"))), nil)
	for iter.Next() {
		var doener Doener
		json.Unmarshal(iter.Value(), &doener)

		d = append(d, doener)

	}
	iter.Release()

	iter = db.NewIterator(util.BytesPrefix([]byte("Tuerkische Pizza"+time.Now().Format("2006-01-02"))), nil)
	for iter.Next() {
		var tuerkische Tuerkischepizza
		json.Unmarshal(iter.Value(), &tuerkische)

		t = append(t, tuerkische)

	}
	iter.Release()

	Orders := Orders{
		Doener:          d,
		Doenerbox:       dbox,
		Tuerkischepizza: t,
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Println(Orders.Doenerbox[0].Kuerzel)

	return c.Render(http.StatusOK, "orders.html", Orders)
}
