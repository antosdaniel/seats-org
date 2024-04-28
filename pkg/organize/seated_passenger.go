package organize

import (
	"fmt"
)

type SeatedPassengers struct {
	matrix [][]*SeatedPassenger
	list   []*SeatedPassenger
	seated map[PassengerId]bool
}

func NewSeatedPassengers(rows, cols int) *SeatedPassengers {
	matrix := make([][]*SeatedPassenger, rows)
	for row := range matrix {
		matrix[row] = make([]*SeatedPassenger, cols)
	}

	return &SeatedPassengers{
		matrix: matrix,
		seated: map[PassengerId]bool{},
	}
}

func (sps *SeatedPassengers) All() []*SeatedPassenger {
	return sps.list
}

func (sps *SeatedPassengers) OccupySeat(layout *Layout, passenger Passenger, row, col int) error {
	if sps.isOutsideOfBoundaries(row, col) {
		return fmt.Errorf("row (%d) and col (%d) is outside of boundary", row, col)
	}
	if sps.IsTaken(row, col) {
		return fmt.Errorf("row (%d) and col (%d) is already occupied", row, col)
	}
	if !layout.IsSeat(row, col) {
		return fmt.Errorf("row (%d) and col (%d) is not a seat", row, col)
	}

	p := &SeatedPassenger{
		row:       row,
		col:       col,
		passenger: passenger,
	}
	sps.list = append(sps.list, p)
	sps.matrix[row][col] = p
	sps.seated[passenger.Id()] = true
	return nil
}

func (sps *SeatedPassengers) IsSeated(id PassengerId) bool {
	return sps.seated[id]
}

func (sps *SeatedPassengers) IsTaken(row, col int) bool {
	return sps.matrix[row][col] != nil
}

func (sps *SeatedPassengers) IsFree(row, col int) bool {
	return !sps.IsTaken(row, col)
}

func (sps *SeatedPassengers) isOutsideOfBoundaries(row, col int) bool {
	if row >= len(sps.matrix) {
		// Outside of matrix
		return true
	}
	if col >= len(sps.matrix[row]) {
		// Outside of matrix
		return true
	}

	return false
}

type SeatedPassenger struct {
	row       int
	col       int
	passenger Passenger
}

func (s SeatedPassenger) Row() int {
	return s.row
}

func (s SeatedPassenger) Col() int {
	return s.col
}

func (s SeatedPassenger) Passenger() Passenger {
	return s.passenger
}
