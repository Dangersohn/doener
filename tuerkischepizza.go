package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type Tuerkischepizza struct {
	Kuerzel   string
	Gericht   string
	Preis     string
	Sosse1    string
	Sosse2    string
	Sosse3    string
	Salat1    string
	Salat2    string
	Salat3    string
	Salat4    string
	Anmerkung string
}

func tuerkischepizza(c echo.Context) error {
	tuerkischepizza := Tuerkischepizza{
		Kuerzel:   strings.ToUpper(c.QueryParam("kuerzel")),
		Gericht:   c.QueryParam("gericht"),
		Preis:     c.QueryParam("preis"),
		Sosse1:    c.QueryParam("sosse1"),
		Sosse2:    c.QueryParam("sosse2"),
		Sosse3:    c.QueryParam("sosse3"),
		Salat1:    c.QueryParam("salat1"),
		Salat2:    c.QueryParam("salat2"),
		Salat3:    c.QueryParam("salat3"),
		Salat4:    c.QueryParam("salat4"),
		Anmerkung: c.QueryParam("anmerkung"),
	}
	j, _ := json.Marshal(doener)

	t := time.Now().Format(time.RFC3339Nano)

	err := db.Put([]byte(t), j, nil)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render(http.StatusOK, "ordertuerkischepizza.html", tuerkischepizza)
}
