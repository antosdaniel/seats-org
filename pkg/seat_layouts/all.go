package seat_layouts

import (
	"errors"

	"github.com/antosdaniel/seats-org/pkg/organize"
)

const (
	x = false
	o = true
)

func Get(name string) (*organize.Layout, error) {
	layout, isSet := Presets[name]
	if !isSet {
		return nil, errors.New("seat layout not found")
	}

	return layout, nil
}

var Presets = map[string]*organize.Layout{
	"Long": organize.MustImportLayout(10, 5, [][]organize.IsSeat{
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, x, x},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, o, o, o},
	}),
	"Short": organize.MustImportLayout(4, 5, [][]organize.IsSeat{
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, o, o, o},
	}),
	"VDL 49 +2": organize.MustImportLayout(13, 5, [][]organize.IsSeat{
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, x, x},
		{o, o, x, x, x},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, o, o, o},
	}),
	"A15 57+2": organize.MustImportLayout(15, 5, [][]organize.IsSeat{
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, x, x},
		{o, o, x, x, x},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, o, o, o},
	}),
}
