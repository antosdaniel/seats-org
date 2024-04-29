package organize

import (
	"fmt"
	"slices"
)

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

func Organize(layout *Layout, passengers Passengers) (Organized, error) {
	seated := NewSeatedPassengers(layout.rows, layout.cols)

	// At first, reserve seats

	// The more preferences passenger has, the harder it is to find a seat for them.
	// We should start with passengers who have the most preferences.
	slices.SortFunc(passengers, func(a, b Passenger) int {
		return len(b.Preferences()) - len(a.Preferences())
	})

	for _, passenger := range passengers {
		hm := newHappinessMap(layout, seated)
		hm.setHappinesForEach(func(row, col int) happiness {
			if len(passenger.preferences) == 0 {
				return fullHappiness
			}

			seat := layout.matrix[row][col]
			fulfilled := 0
			possible := 0
			for _, preference := range passenger.preferences {
				switch preference {
				case FrontSeatPreference:
					possible += 2
					if seat.mostFront {
						fulfilled += 2
					} else if seat.front {
						fulfilled += 1
					}
				case RearSeatPreference:
					possible += 2
					if seat.mostRear {
						fulfilled += 2
					} else if seat.rear {
						fulfilled += 1
					}
				default:
					panic(fmt.Sprintf("uknown preference %q", preference))
				}
			}
			return fullHappiness * happiness(fulfilled/possible)
		})

		row, col, exists := hm.firstHappiestSeat()
		if !exists {
			// TODO
			panic(fmt.Sprintf("did not find a seat for passanger %q", passenger.Id()))
		}

		err := seated.OccupySeat(layout, passenger, row, col)
		if err != nil {
			return Organized{}, fmt.Errorf("could not seat passenger %q at the happiest seat: %w", passenger.id, err)
		}
	}

	return Organized{
		SeatedPassengers: seated,
	}, nil
}

type happinessMap struct {
	rows         int
	cols         int
	matrix       [][]*happiness
	maxHappiness happiness
}

type happiness int

const fullHappiness happiness = 100

func newHappinessMap(layout *Layout, seated *SeatedPassengers) happinessMap {
	matrix := make([][]*happiness, layout.rows)
	for row := range matrix {
		matrix[row] = make([]*happiness, layout.cols)
		for col := range layout.cols {
			if !layout.IsSeat(row, col) {
				continue
			}
			if seated.IsTaken(row, col) {
				continue
			}

			tmp := happiness(0)
			matrix[row][col] = &tmp
		}
	}

	return happinessMap{
		rows:   layout.rows,
		cols:   layout.cols,
		matrix: matrix,
	}
}

func (hm *happinessMap) setHappinesForEach(happinessEstimator func(row, col int) happiness) {
	for row := range hm.rows {
		for col := range hm.cols {
			if hm.matrix[row][col] == nil {
				continue
			}

			h := happinessEstimator(row, col)
			hm.matrix[row][col] = &h

			if h > hm.maxHappiness {
				hm.maxHappiness = h
			}
		}
	}
}

func (hm *happinessMap) firstHappiestSeat() (row, col int, exists bool) {
	for row := range hm.rows {
		for col := range hm.cols {
			if hm.matrix[row][col] == nil {
				continue
			}

			if *hm.matrix[row][col] == hm.maxHappiness {
				return row, col, true
			}
		}
	}

	return 0, 0, false
}
