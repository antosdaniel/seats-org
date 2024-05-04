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
	"44+2 (A17)": organize.MustImportLayout([][]organize.IsSeat{
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
	}),
	"49+2 (A21, A23, A27, A38, A47)": organize.MustImportLayout([][]organize.IsSeat{
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
	"49+2 (A45, A46, A53)": organize.MustImportLayout([][]organize.IsSeat{
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, x, x},
		{o, o, x, x, x},
		{o, o, x, x, x},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, o, o, o},
	}),
	"55+2 (A71)": organize.MustImportLayout([][]organize.IsSeat{
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, x, x},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, o, o, o},
	}),
	"57+2 (A15, A34)": organize.MustImportLayout([][]organize.IsSeat{
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
