package layout

import (
	"fmt"
	"strings"

	"github.com/antosdaniel/seats-org/pkg/organize"
)

type OrganizedPassenger struct {
	ID          string
	FullName    string
	SeatID      string
	Seat        string
	Preferences string
}

func OrganizedPassengerList(organized organize.Organized) []OrganizedPassenger {
	sorted := make([]*organize.SeatedPassenger, len(organized.SeatedPassengers.All()))
	copy(sorted, organized.SeatedPassengers.All())
	organize.SortByOrderAsc(sorted, organized.Layout)

	result := []OrganizedPassenger{}
	for _, i := range sorted {
		p := OrganizedPassenger{
			ID:          string(i.Passenger().Id()),
			FullName:    i.Passenger().Details.FullName,
			SeatID:      seatID(i.Row(), i.Col()),
			Seat:        fmt.Sprintf("%2d - %2d", i.Row(), i.Col()),
			Preferences: printPreferences(i.Passenger().Preferences()),
		}
		result = append(result, p)
	}

	return result
}

func printPreferences(preferences organize.Preferences) string {
	if len(preferences) == 0 {
		return "-"
	}

	result := []string{}
	for _, p := range preferences {
		result = append(result, printPreference(p))
	}
	return strings.Join(result, ", ")
}

func printPreference(preference organize.Preference) string {
	switch preference {
	case organize.FrontSeatPreference:
		return "przód"
	case organize.RearSeatPreference:
		return "tył"
	case organize.WindowSeatPreference:
		return "przy oknie"
	case organize.AisleSeatPreference:
		return "przy przejściu"
	default:
		panic(fmt.Errorf("unknown preference %q", preference))
	}
}

func cellAttrs(layout *organize.Layout, row, col int) string {
	result := seatID(row, col)
	if layout.IsSeat(row, col) {
		result += " seat"
	}
	return result
}

func seatID(row, col int) string {
	return fmt.Sprintf("seat-%d-%d", row, col)
}
