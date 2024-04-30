package organize

type happiness int

const fullHappiness happiness = 100

type happinessMap struct {
	rows         int
	cols         int
	matrix       [][]*happiness
	maxHappiness happiness
}

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

func (hm *happinessMap) setHappinessForEach(happinessEstimator func(row, col int) happiness) {
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
