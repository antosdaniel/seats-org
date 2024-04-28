package organize

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportLayout(t *testing.T) {
	testCases := []struct {
		name string

		rows   int
		cols   int
		matrix [][]IsSeat

		want    *Layout
		wantErr error
	}{
		{
			name: "number of declared rows does not match",

			rows: 3,
			cols: 3,

			wantErr: errors.New("provided (0) different amount of rows than promised (3)"),
		},
		{
			name: "number of declared cols does not match",

			rows: 3,
			cols: 3,
			matrix: [][]IsSeat{
				{O, O, O},
				{O, O},
				{O, O, O},
			},

			wantErr: errors.New("provided (2) different amount of cols than promised (3) at row 1"),
		},
		{
			name: "sets seats correctly",

			rows: 4,
			cols: 4,
			matrix: [][]IsSeat{
				{O, O, O, O},
				{O, X, O, O},
				{X, X, O, O},
				{O, X, O, O},
			},

			want: &Layout{
				rows: 4,
				cols: 4,
				matrix: [][]*Seat{
					{
						{window: true, aisle: false},
						{window: false, aisle: false},
						{window: false, aisle: false},
						{window: true, aisle: false},
					},
					{
						{window: true, aisle: true},
						nil,
						{window: false, aisle: true},
						{window: true, aisle: false},
					},
					{
						nil,
						nil,
						{window: false, aisle: true},
						{window: true, aisle: false},
					},
					{
						{window: true, aisle: true},
						nil,
						{window: false, aisle: true},
						{window: true, aisle: false},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ImportLayout(tc.rows, tc.cols, tc.matrix)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLayout_NextToRightOrInNextRow(t *testing.T) {
	layout := MustImportLayout(3, 3, [][]IsSeat{
		{O, X, O},
		{O, X, O},
		{X, O, X},
	})

	testCases := []struct {
		name string

		row int
		col int

		wantRow    int
		wantCol    int
		wantExists bool
	}{
		{
			name: "next seat does not exist",

			row: 2,
			col: 2,

			wantExists: false,
		},
		{
			name: "next seat is the first one to the right",

			row: 0,
			col: 1,

			wantRow:    0,
			wantCol:    2,
			wantExists: true,
		},
		{
			name: "next seat is the next row",

			row: 0,
			col: 2,

			wantRow:    1,
			wantCol:    0,
			wantExists: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			row, col, exists := layout.NextToRightOrInNextRow(testCase.row, testCase.col)

			assert.Equalf(t, testCase.wantRow, row, "row")
			assert.Equalf(t, testCase.wantCol, col, "col")
			assert.Equalf(t, testCase.wantExists, exists, "exists")
		})
	}
}
