package organize

import (
	"fmt"
	"slices"
)

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
		pref := len(b.Preferences()) - len(a.Preferences())
		if pref != 0 {
			return pref
		}

		// When passengers have the same number of preferences, order them by when they signed up
		return a.SignupOrder() - b.SignupOrder()
	})

	for _, passenger := range passengers {
		hm := newHappinessMap(layout, seated)
		hm.setHappinessForEach(func(row, col int) happiness {
			if len(passenger.preferences) == 0 {
				return fullHappiness
			}

			seat := layout.matrix[row][col]

			const maxPoints = 10
			possible := maxPoints * len(passenger.preferences)
			fulfilled := 0

			for _, preference := range passenger.preferences {
				switch preference {
				case FrontSeatPreference:
					result := maxPoints - row
					if result > 0 {
						fulfilled += result
					}
				case RearSeatPreference:
					diff := layout.LastRow() - row
					result := maxPoints - diff
					if result > 0 {
						fulfilled += result
					}
				case WindowSeatPreference:
					if seat.window {
						fulfilled += maxPoints
					}
				case AisleSeatPreference:
					if seat.aisle {
						fulfilled += maxPoints
					}
				default:
					panic(fmt.Sprintf("uknown preference %q", preference))
				}
			}
			return happiness(int(fullHappiness) * fulfilled / possible)
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
