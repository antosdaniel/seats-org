package main

import (
	"encoding/csv"
	"fmt"
	"sort"

	"github.com/antosdaniel/seats-org/pkg/seat_layouts"
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

		keys := []string{}
		for name := range seat_layouts.Presets {
			keys = append(keys, name)
		}
		sort.Strings(keys)

		layouts := make([]layout.SeatLayout, len(seat_layouts.Presets))
		for i, name := range keys {
			layouts[i] = layout.NewSeatLayout(name, seat_layouts.Presets[name])
		}

		return layout.Base(layouts).
			Render(c.Request().Context(), c.Response())
	})

	e.POST("/organize", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		file, err := c.FormFile("passengers")
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

	e.GET("/preview", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		name := c.QueryParam("layout")
		s, isSet := seat_layouts.Presets[name]
		if !isSet {
			_, _ = c.Response().Write([]byte("seat layout not found"))
			return nil
		}

		selected := layout.NewSeatLayout(name, s)
		_, _ = c.Response().Write([]byte(selected.Visual))
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
