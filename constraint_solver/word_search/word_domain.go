package word_search

import "math/rand"

func (w *WordGrid) generateDomain(word string) [][]GridLocation {
	domain := [][]GridLocation{}
	length := len(word)

	for row := 0; row < w.rows; row++ {
		for column := 0; column < w.columns; column++ {
			if column+length <= w.columns {
				// left to right
				fillRight(&domain, row, column, length)
				if row+length <= w.rows {
					// left to right diagonal
					fillDiagonalRight(&domain, row, column, length)
				}
			}
			if row+length <= w.rows {
				// top to bottom
				fillDown(&domain, row, column, length)
				if column-length >= 0 {
					// top to bottom diagonal
					fillDiagonalLeft(&domain, row, column, length)
				}
			}
		}
	}

	rand.Shuffle(len(domain), func(i, j int) {
		domain[i], domain[j] = domain[j], domain[i]
	})

	return domain
}

func fillRight(domain *[][]GridLocation, row, column, length int) {
	locations := []GridLocation{}
	for i := 0; i < length; i++ {
		locations = append(locations, GridLocation{row, column + i})
	}
	*domain = append(*domain, locations)
}

func fillDiagonalRight(domain *[][]GridLocation, row, column, length int) {
	locations := []GridLocation{}
	for i := 0; i < length; i++ {
		locations = append(locations, GridLocation{row + i, column + i})
	}
	*domain = append(*domain, locations)
}

func fillDown(domain *[][]GridLocation, row, column, length int) {
	locations := []GridLocation{}
	for i := 0; i < length; i++ {
		locations = append(locations, GridLocation{row + i, column})
	}
	*domain = append(*domain, locations)
}

func fillDiagonalLeft(domain *[][]GridLocation, row, column, length int) {
	locations := []GridLocation{}
	for i := 0; i < length; i++ {
		locations = append(locations, GridLocation{row + i, column - i})
	}
	*domain = append(*domain, locations)
}
