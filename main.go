package main

import (
	"embed"
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"sort"
	"strings"

	"github.com/antosdaniel/seats-org/pkg/organize"
	"github.com/antosdaniel/seats-org/pkg/seat_layouts"
	"github.com/antosdaniel/seats-org/views/layout"
	"github.com/benbjohnson/hashfs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed assets
var embedFS embed.FS
var assetsFS = hashfs.NewFS(embedFS)

func main() {
	e := echo.New()
	
	e.Use(middleware.Logger())

	e.StaticFS("/", assetsFS)

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

		return layout.Base(assetsFS.HashName, layouts).
			Render(c.Request().Context(), c.Response())
	})

	e.POST("/organize", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		file, err := c.FormFile("passengers")
		if err != nil {
			return fmt.Errorf("file not provided: %w", err)
		}

		l, err := seat_layouts.Get(c.FormValue("layout"))
		if err != nil {
			return err
		}

		passengers, err := readListOfPassengers(file)
		if err != nil {
			return fmt.Errorf("could not read list of passengers: %w", err)
		}

		result, err := organize.Organize(l, passengers)
		if err != nil {
			return err
		}

		return layout.OrganizeResult(result).Render(c.Request().Context(), c.Response())
	})

	e.GET("/preview", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		name := c.QueryParam("layout")
		l, err := seat_layouts.Get(name)
		if err != nil {
			_, _ = c.Response().Write([]byte(err.Error()))
			return nil
		}

		selected := layout.NewSeatLayout(name, l)
		_, _ = c.Response().Write([]byte(selected.Visual))
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func readListOfPassengers(file *multipart.FileHeader) (organize.Passengers, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer src.Close()

	records, err := csv.NewReader(src).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV: %w", err)
	}

	if len(records) < 1 {
		return nil, fmt.Errorf("CSV has no rows")
	}

	idIndex := -1
	fullNameIndex := -1
	preferencesIndex := -1
	for col := range records[0] {
		header := strings.TrimSpace(strings.ToLower(records[0][col]))

		if header == "id" {
			idIndex = col
			continue
		}

		if strings.Contains(header, "nazwisko") {
			fullNameIndex = col
			continue
		}

		if strings.Contains(header, "preferencje") {
			preferencesIndex = col
			continue
		}
	}

	if idIndex < 0 {
		return nil, fmt.Errorf("ID column not found")
	}
	if fullNameIndex < 0 {
		return nil, fmt.Errorf("full name column not found")
	}
	if preferencesIndex < 0 {
		return nil, fmt.Errorf("preferences column not found")
	}

	result := organize.Passengers{}
	for i := 1; i < len(records); i++ {
		row := records[i]
		id := organize.PassengerId(row[idIndex])
		if id == "" {
			continue
		}

		preferences, err := parsePreferences(row[preferencesIndex])
		if err != nil {
			return nil, err
		}

		passenger := organize.NewPassenger(id, i, preferences...)
		passenger.Details.FullName = row[fullNameIndex]
		result = append(result, passenger)
	}

	return result, nil
}

func parsePreferences(in string) (organize.Preferences, error) {
	if in == "" {
		return nil, nil
	}

	result := organize.Preferences{}
	for _, i := range strings.Split(in, ",") {
		p := strings.TrimSpace(strings.ToLower(i))
		switch p {
		case "okno":
			result = append(result, organize.WindowSeatPreference)
		case "przejście":
			result = append(result, organize.AisleSeatPreference)
		case "przód":
			result = append(result, organize.FrontSeatPreference)
		case "tył":
			result = append(result, organize.RearSeatPreference)
		default:
			return nil, fmt.Errorf("unknown preference %q", p)
		}
	}

	// TODO: validate that preferences can be mixed
	return result, nil
}
