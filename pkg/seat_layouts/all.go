package seat_layouts

import "github.com/antosdaniel/seats-org/pkg/organize"

const (
	x = false
	o = true
)

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
}
