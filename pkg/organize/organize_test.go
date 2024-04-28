package organize_test

import (
	"testing"

	"github.com/antosdaniel/seats-org/pkg/organize"

	"github.com/stretchr/testify/assert"
)

const (
	x = false
	o = true
)

var exampleLayout = organize.MustImportLayout(5, 5, [][]organize.IsSeat{
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

		want []seatedPassenger
	}{
		{
			name:   "passengers prefer front seats",
			layout: exampleLayout,
			passengers: organize.Passengers{
				organize.NewPassenger("01", "", organize.FrontSeatPreference),
				organize.NewPassenger("02", "", organize.FrontSeatPreference),
				organize.NewPassenger("03", "", organize.FrontSeatPreference),
				organize.NewPassenger("04", "", organize.FrontSeatPreference),
				organize.NewPassenger("05", "", organize.FrontSeatPreference),
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
				organize.NewPassenger("01", "", organize.RearSeatPreference),
				organize.NewPassenger("02", "", organize.RearSeatPreference),
				organize.NewPassenger("03", "", organize.RearSeatPreference),
				organize.NewPassenger("04", "", organize.RearSeatPreference),
				organize.NewPassenger("05", "", organize.RearSeatPreference),
			},
			want: []seatedPassenger{},
		},
		{
			name:       "passengers prefer window seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "passengers prefer aisle seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "passengers prefer front window seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "passengers prefer front aisle seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "passengers prefer rear window seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "passengers prefer rear aisle seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "group sits together",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "solo travelers sit alone",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
		{
			name:       "solo travelers of the same gender sit together, if there are no available seats",
			layout:     exampleLayout,
			passengers: organize.Passengers{},
			//want:       organize.Organized{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := organize.Organize(tc.layout, tc.passengers)

			parsed := []seatedPassenger{}
			for _, i := range got.SeatedPassengers.All() {
				parsed = append(parsed, seatedPassenger{
					id:  i.Passenger().Id(),
					row: i.Row(),
					col: i.Col(),
				})
			}
			assert.Equal(t, tc.want, parsed)
		})
	}
}
