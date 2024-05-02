package organize

import (
	"errors"
	"fmt"
	"slices"
)

type IsSeat bool

type Seat struct {
	aisle  bool
	window bool
}

type Layout struct {
	rows, cols int

	matrix [][]*Seat
}

func (l *Layout) Rows() int {
	return l.rows
}

func (l *Layout) Cols() int {
	return l.cols
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

func (l *Layout) FirstRow() int {
	return 0
}

func (l *Layout) LastRow() int {
	return l.rows - 1
}

func (l *Layout) SeatNumber(row, col int) int {
	if !l.IsSeat(row, col) {
		return -1
	}

	number := 0
	for r := range l.Rows() {
		for c := range l.Cols() {
			if !l.IsSeat(r, c) {
				continue
			}

			number++

			if r == row && c == col {
				return number
			}
		}
	}
	return -1
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

func (l *Layout) SeatsMatrix() [][]IsSeat {
	result := make([][]IsSeat, l.rows)
	for row := range l.rows {
		result[row] = make([]IsSeat, l.cols)
		for col := range l.cols {
			result[row][col] = IsSeat(l.IsSeat(row, col))
		}
	}

	return result
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

var EmptyLayoutErr = errors.New("layout needs rows and cols")

func ImportLayout(matrix [][]IsSeat) (*Layout, error) {
	rows := len(matrix)
	if rows == 0 {
		return nil, EmptyLayoutErr
	}
	cols := len(matrix[0])
	if cols == 0 {
		return nil, EmptyLayoutErr
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

			layout.matrix[row][col] = s
		}
	}

	return layout, nil
}

func MustImportLayout(matrix [][]IsSeat) *Layout {
	layout, err := ImportLayout(matrix)
	if err != nil {
		panic(err)
	}
	return layout
}

func OrderOf(layout *Layout, row, col int) int {
	return (row+1)*layout.cols + col
}

type Cell interface {
	Row() int
	Col() int
}

func SortByOrderAsc[T Cell](items []T, layout *Layout) {
	slices.SortFunc(items, func(a, b T) int {
		return OrderOf(layout, a.Row(), a.Col()) - OrderOf(layout, b.Row(), b.Col())
	})
}
