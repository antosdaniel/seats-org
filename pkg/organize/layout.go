package organize

import (
	"fmt"
)

type IsSeat bool

type Seat struct {
	aisle  bool
	window bool

	mostFront bool
	front     bool
	mostRear  bool
	rear      bool
}

type Layout struct {
	rows, cols int

	matrix [][]*Seat
}

func (l *Layout) IsSeat(row, col int) bool {
	if row >= l.rows {
		return false
	}
	if col >= l.cols {
		return false
	}
	return l.matrix[row][col] != nil
}

func (l *Layout) FrontSeat() (row, col int) {
	return 0, 0
}

func (l *Layout) NextToRight(row, col int) (r, c int, exists bool) {
	col += 1

	if l.isOutsideOfBoundaries(row, col) {
		return 0, 0, false
	}

	if !l.IsSeat(row, col) {
		return l.NextToRight(row, col)
	}

	return row, col, true
}

func (l *Layout) ThisOrToTheRight(row, col int) (r, c int, exists bool) {
	if l.isOutsideOfBoundaries(row, col) {
		return 0, 0, false
	}

	if !l.IsSeat(row, col) {
		return l.ThisOrToTheRight(row, col+1)
	}

	return row, col, true
}

func (l *Layout) NextRow(row, col int) (r, c int, exists bool) {
	row += 1

	if row >= l.rows {
		return 0, 0, false
	}

	if !l.IsSeat(row, col) {
		return 0, 0, false
	}

	return row, col, true
}

func (l *Layout) NextToRightOrInNextRow(row, col int) (r, c int, exists bool) {
	if l.isOutsideOfBoundaries(row, col) {
		return 0, 0, false
	}

	r, c, exists = l.NextToRight(row, col)
	if exists {
		return
	}

	// We are checking if the first next row is an available seat, because recursive check will skip first column.
	row += 1
	col = 0
	if l.IsSeat(row, col) {
		return row, col, true
	}

	return l.NextToRightOrInNextRow(row, col)
}

func (l *Layout) isOutsideOfBoundaries(row, col int) bool {
	if row >= l.rows {
		// Outside of matrix
		return true
	}
	if col >= l.cols {
		// Outside of matrix
		return true
	}

	return false
}

func NewLayout(rows, cols int) *Layout {
	matrix := make([][]*Seat, rows)
	for row := range matrix {
		matrix[row] = make([]*Seat, cols)
	}

	return &Layout{
		rows:   rows,
		cols:   cols,
		matrix: matrix,
	}
}

func ImportLayout(rows, cols int, matrix [][]IsSeat) (*Layout, error) {
	if len(matrix) != rows {
		return nil, fmt.Errorf("provided (%d) different amount of rows than promised (%d)", len(matrix), rows)
	}

	layout := NewLayout(rows, cols)
	for row := range layout.matrix {
		if len(matrix[row]) != cols {
			return nil, fmt.Errorf(
				"provided (%d) different amount of cols than promised (%d) at row %d",
				len(matrix[row]), cols, row,
			)
		}
		for col := range layout.matrix[row] {
			if !matrix[row][col] {
				continue
			}

			s := &Seat{}
			if col == 0 || col == len(layout.matrix[row])-1 {
				s.window = true
			}

			if col+1 < cols && !matrix[row][col+1] {
				// The next space is corridor
				s.aisle = true
			} else if col > 0 && !matrix[row][col-1] {
				// The previous space is corridor
				s.aisle = true
			}

			s.mostFront = row == 0
			s.front = row < 4
			s.mostRear = row == rows-1
			s.rear = row > rows-4

			layout.matrix[row][col] = s
		}
	}

	return layout, nil
}

func MustImportLayout(rows, cols int, matrix [][]IsSeat) *Layout {
	layout, err := ImportLayout(rows, cols, matrix)
	if err != nil {
		panic(err)
	}
	return layout
}
