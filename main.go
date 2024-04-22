package main

import (
	"fmt"

	layout "github.com/antosdaniel/seats-org/views"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "assets")

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		return layout.Base(false, c.QueryParam("name")).
			Render(c.Request().Context(), c.Response())
	})

	e.GET("/search", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		_, _ = c.Response().Write([]byte(fmt.Sprintf("%s", c.QueryParam("name"))))
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
