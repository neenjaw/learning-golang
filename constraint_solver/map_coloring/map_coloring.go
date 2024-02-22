package map_coloring

import "constraint_solver/csp"

type MapColoringConstraint struct {
	First, Second string
}

func (m MapColoringConstraint) Variables() []string {
	return []string{m.First, m.Second}
}

func (m MapColoringConstraint) Satisfied(assignment map[string]string) bool {
	if _, ok := assignment[m.First]; !ok {
		return true
	}
	if _, ok := assignment[m.Second]; !ok {
		return true
	}
	return assignment[m.First] != assignment[m.Second]
}

func MapVariables() []string {
	return []string{
		"Western Australia",
		"Northern Territory",
		"South Australia",
		"Queensland",
		"New South Wales",
		"Victoria",
		"Tasmania",
	}
}

func MapColors() []string {
	return []string{"red", "green", "blue"}
}

func MapColoringCSP() *csp.CSPSolver[string, string] {
	variables := MapVariables()
	domains := make(map[string][]string)
	for _, variable := range variables {
		domains[variable] = MapColors()
	}
	csp, _ := csp.NewCSPSolver(variables, domains)

	csp.AddConstraint(MapColoringConstraint{"Western Australia", "Northern Territory"})
	csp.AddConstraint(MapColoringConstraint{"Western Australia", "South Australia"})
	csp.AddConstraint(MapColoringConstraint{"South Australia", "Northern Territory"})
	csp.AddConstraint(MapColoringConstraint{"Queensland", "Northern Territory"})
	csp.AddConstraint(MapColoringConstraint{"Queensland", "South Australia"})
	csp.AddConstraint(MapColoringConstraint{"Queensland", "New South Wales"})
	csp.AddConstraint(MapColoringConstraint{"New South Wales", "South Australia"})
	csp.AddConstraint(MapColoringConstraint{"Victoria", "South Australia"})
	csp.AddConstraint(MapColoringConstraint{"Victoria", "New South Wales"})
	csp.AddConstraint(MapColoringConstraint{"Victoria", "Tasmania"})

	return csp
}

// uniqueCombinations returns all unique combinations of two elements from the
// given slice of strings.
func uniqueCombinations(
	elements []string,
) [][2]string {
	combinations := [][2]string{}

	for i := 0; i < len(elements); i++ {
		for j := i + 1; j < len(elements); j++ {
			combinations = append(combinations, [2]string{elements[i], elements[j]})
		}
	}

	return combinations
}
