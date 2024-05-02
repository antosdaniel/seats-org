package organize_test

import (
	"slices"
	"testing"

	"github.com/antosdaniel/seats-org/pkg/organize"

	"github.com/stretchr/testify/assert"
)

const (
	x = false
	o = true
)

var exampleLayout = organize.MustImportLayout([][]organize.IsSeat{
	{o, o, x, o, o},
	{o, o, x, o, o},
	{o, o, x, x, x},
	{o, o, x, o, o},
	{o, o, o, o, o},
})

type seatedPassenger struct {
	id  organize.PassengerId
	row int
	col int
}

func TestOrganize(t *testing.T) {
	testCases := []struct {
		name string

		layout     *organize.Layout
		passengers organize.Passengers

		want    []seatedPassenger
		wantErr error
	}{
		{
			name:   "passengers prefer front seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", 1, organize.FrontSeatPreference),
				organize.NewPassenger("02", 2, organize.FrontSeatPreference),
				organize.NewPassenger("03", 3, organize.FrontSeatPreference),
				organize.NewPassenger("04", 4, organize.FrontSeatPreference),
				organize.NewPassenger("05", 5, organize.FrontSeatPreference),
			},
			want: []seatedPassenger{
				{id: "01", row: 0, col: 0},
				{id: "02", row: 0, col: 1},
				{id: "03", row: 0, col: 3},
				{id: "04", row: 0, col: 4},
				{id: "05", row: 1, col: 0},
			},
		},
		{
			name:   "passengers prefer rear seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", 1, organize.RearSeatPreference),
				organize.NewPassenger("02", 2, organize.RearSeatPreference),
				organize.NewPassenger("03", 3, organize.RearSeatPreference),
				organize.NewPassenger("04", 4, organize.RearSeatPreference),
				organize.NewPassenger("05", 5, organize.RearSeatPreference),
				organize.NewPassenger("06", 6, organize.RearSeatPreference),
			},
			want: []seatedPassenger{
				{id: "06", row: 3, col: 0},
				{id: "01", row: 4, col: 0},
				{id: "02", row: 4, col: 1},
				{id: "03", row: 4, col: 2},
				{id: "04", row: 4, col: 3},
				{id: "05", row: 4, col: 4},
			},
		},
		{
			name:   "passengers prefer window seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", 1, organize.WindowSeatPreference),
				organize.NewPassenger("02", 2, organize.WindowSeatPreference),
				organize.NewPassenger("03", 3, organize.WindowSeatPreference),
			},

			want: []seatedPassenger{
				{id: "01", row: 0, col: 0},
				{id: "02", row: 0, col: 4},
				{id: "03", row: 1, col: 0},
			},
		},
		{
			name:   "passengers prefer aisle seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", 1, organize.AisleSeatPreference),
				organize.NewPassenger("02", 2, organize.AisleSeatPreference),
				organize.NewPassenger("03", 3, organize.AisleSeatPreference),
			},

			want: []seatedPassenger{
				{id: "01", row: 0, col: 1},
				{id: "02", row: 0, col: 3},
				{id: "03", row: 1, col: 1},
			},
		},
		{
			name:   "passengers prefer front window seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", 1, organize.FrontSeatPreference, organize.WindowSeatPreference),
				organize.NewPassenger("02", 2, organize.FrontSeatPreference),
				organize.NewPassenger("03", 3, organize.FrontSeatPreference),
				organize.NewPassenger("04", 4, organize.FrontSeatPreference),
				organize.NewPassenger("05", 5, organize.FrontSeatPreference, organize.WindowSeatPreference),
			},

			want: []seatedPassenger{
				{id: "01", row: 0, col: 0}, // front + window
				{id: "02", row: 0, col: 1},
				{id: "03", row: 0, col: 3},
				{id: "05", row: 0, col: 4}, // front + window
				{id: "04", row: 1, col: 0},
			},
		},
		{
			name:   "passengers prefer rear aisle seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", 1, organize.RearSeatPreference, organize.AisleSeatPreference),
				organize.NewPassenger("02", 2, organize.RearSeatPreference),
				organize.NewPassenger("03", 3, organize.RearSeatPreference, organize.AisleSeatPreference),
				organize.NewPassenger("04", 4, organize.RearSeatPreference),
				organize.NewPassenger("05", 5, organize.RearSeatPreference, organize.AisleSeatPreference),
			},

			want: []seatedPassenger{
				{id: "05", row: 2, col: 1}, // rear + aisle
				{id: "01", row: 3, col: 1}, // rear + aisle, signed up first, so takes most rear
				{id: "03", row: 3, col: 3}, // rear + aisle
				{id: "02", row: 4, col: 0},
				{id: "04", row: 4, col: 1},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := organize.Organize(tc.layout, tc.passengers)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.NoError(t, err)
			parsed := []seatedPassenger{}
			for _, i := range got.SeatedPassengers.All() {
				parsed = append(parsed, seatedPassenger{
					id:  i.Passenger().Id(),
					row: i.Row(),
					col: i.Col(),
				})
			}
			slices.SortFunc(parsed, func(a, b seatedPassenger) int {
				return organize.OrderOf(tc.layout, a.row, a.col) - organize.OrderOf(tc.layout, b.row, b.col)
			})
			assert.Equal(t, tc.want, parsed)
		})
	}
}
