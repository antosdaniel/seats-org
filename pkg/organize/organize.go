package organize

import "fmt"

const (
	X = false
	O = true
)

var ExampleLayout = MustImportLayout(10, 5, [][]IsSeat{
	{O, O, O, O, O},
	{O, O, X, O, O},
	{O, O, X, O, O},
	{O, O, X, O, O},
	{X, X, X, O, O},
	{O, O, X, O, O},
	{O, O, X, O, O},
	{O, O, X, O, O},
	{O, O, X, O, O},
	{O, O, X, O, O},
})

type Organized struct {
	SeatedPassengers          *SeatedPassengers
	EmptySeatsLeft            int
	PossiblyUnhappyPassengers []any
}

func Organize(layout *Layout, passengers Passengers) Organized {
	seated := NewSeatedPassengers(layout.rows, layout.cols)
	// At first, find seats for premium passengers

	// Then, find seats

	err := findForFrontSeatPreferences(seated, layout, passengers)
	if err != nil {
		panic(err)
	}

	return Organized{
		SeatedPassengers: seated,
	}
}

func findForFrontSeatPreferences(seated *SeatedPassengers, layout *Layout, passengers Passengers) error {
	for _, passenger := range passengers.WithPreference(FrontSeatPreference) {
		if seated.IsSeated(passenger.Id()) {
			continue
		}

		row, col := layout.FrontSeat()
		for {
			if seated.IsFree(row, col) {
				err := seated.OccupySeat(layout, passenger, row, col)
				if err != nil {
					return err
				}
				break
			}

			var exists bool
			row, col, exists = layout.NextToRightOrInNextRow(row, col)
			if !exists {
				return fmt.Errorf("impossible to find next seat (row=%d, col=%d)", row, col)
			}
		}
	}

	return nil
}
