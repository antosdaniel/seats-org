package organize

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	x = false
	o = true
)

func TestImportLayout(t *testing.T) {
	testCases := []struct {
		name string

		matrix [][]IsSeat

		want    *Layout
		wantErr error
	}{
		{
			name: "number of declared rows does not match",

			wantErr: EmptyLayoutErr,
		},
		{
			name: "number of declared cols does not match",

			matrix: [][]IsSeat{
				{o, o, o},
				{o, o},
				{o, o, o},
			},

			wantErr: errors.New("provided (2) different amount of cols than promised (3) at row 1"),
		},
		{
			name: "sets seats correctly",

			matrix: [][]IsSeat{
				{o, o, o, o},
				{o, x, o, o},
				{x, x, o, o},
				{o, x, o, o},
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
			got, err := ImportLayout(tc.matrix)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
