package main

import (
	"encoding/csv"
	"fmt"

	"github.com/antosdaniel/seats-org/views/layout"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/", "assets")

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		return layout.Base(false, c.QueryParam("name")).
			Render(c.Request().Context(), c.Response())
	})

	e.POST("/organize", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		file, err := c.FormFile("file")
		if err != nil {
			return fmt.Errorf("file not provided: %w", err)
		}

		src, err := file.Open()
		if err != nil {
			return fmt.Errorf("could not open file: %w", err)
		}
		defer src.Close()

		records, err := csv.NewReader(src).ReadAll()
		if err != nil {
			return fmt.Errorf("could not read CSV: %w", err)
		}

		_, _ = c.Response().Write([]byte(fmt.Sprintf("%v", records)))
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
