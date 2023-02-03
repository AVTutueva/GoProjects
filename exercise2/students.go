package exercise2

import "fmt"

func Students() {
	results := map[string][]int{
		"Aleksandra": {5, 6, 7, 8, 10},
		"Ana":        {6, 7, 8, 9, 8},
		"Alina":      {6, 7, 7, 6, 7},
		"Simon":      {8, 6, 4, 5, 5},
		"Lukas":      {5, 6, 4, 5, 5},
	}

	for student, marks := range results {
		max_mark := 0
		for _, mark := range marks {
			if mark > max_mark {
				max_mark = mark
			}
		}
		fmt.Printf("%s: %d\n", student, max_mark)
	}

}
