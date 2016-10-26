package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	//fmt.Println("hilo,hhaa")
	e := echo.New()
	e.Get("/test", Test)
	e.Run(standard.New(":1323"))
}

//test
func Test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
