package layout

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antosdaniel/seats-org/pkg/organize"
)

type SeatLayout struct {
	Name   string
	Matrix string
	Visual string
}

func NewSeatLayout(name string, layout *organize.Layout) SeatLayout {
	matrix := layout.SeatsMatrix()
	result, err := json.Marshal(matrix)
	if err != nil {
		panic(fmt.Errorf("could not marshal layout seats: %w", err))
	}

	rows := make([]string, layout.Rows())
	for row := range layout.Rows() {
		for col := range layout.Cols() {
			val := represent(matrix[row][col])
			if rows[row] != "" {
				rows[row] += " "
			}
			rows[row] += val
		}
	}

	return SeatLayout{
		Name:   name,
		Matrix: string(result),
		Visual: strings.Join(rows, "\n"),
	}
}

const (
	available   = "O"
	unavailable = "X"
)

func represent(in organize.IsSeat) string {
	if in {
		return available
	}

	return unavailable
}
