package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type DoenerBox struct {
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
	Pommes    string
	Anmerkung string
}

func doenerbox(c echo.Context) error {
	doenerbox := DoenerBox{
		//Algemein
		Kuerzel:   strings.ToUpper(c.QueryParam("kuerzel")),
		Anmerkung: c.QueryParam("anmerkung"),
		//Info Gericht
		Gericht: c.QueryParam("gericht"),
		Preis:   c.QueryParam("preis"),
		//Sosse
		Sosse1: c.QueryParam("sosse1"),
		Sosse2: c.QueryParam("sosse2"),
		Sosse3: c.QueryParam("sosse3"),
		//Salat
		Salat1: c.QueryParam("salat1"),
		Salat2: c.QueryParam("salat2"),
		Salat3: c.QueryParam("salat3"),
		Salat4: c.QueryParam("salat4"),
		//Pommes
		Pommes: c.QueryParam("pommes"),
	}
	j, _ := json.Marshal(doener)

	t := time.Now().Format(time.RFC3339Nano)

	err := db.Put([]byte(t), j, nil)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render(http.StatusOK, "orderdoenerbox.html", doenerbox)
}
