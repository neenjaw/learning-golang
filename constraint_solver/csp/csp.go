package csp

import (
	"fmt"
	"maps"
	"slices"
)

type Constraint[V comparable, D any] interface {
	Variables() []V
	Satisfied(map[V]D) bool
}

type CSPSolver[V comparable, D any] struct {
	Variables   []V
	Domains     map[V][]D
	Constraints map[V][]Constraint[V, D]
}

func NewCSPSolver[V comparable, D any](
	variables []V,
	domains map[V][]D,
) (*CSPSolver[V, D], error) {
	csp := CSPSolver[V, D]{
		Variables:   slices.Clone(variables),
		Domains:     maps.Clone(domains),
		Constraints: make(map[V][]Constraint[V, D], len(variables)),
	}

	for _, variable := range variables {
		if _, ok := domains[variable]; !ok {
			return nil, fmt.Errorf("no domain for variable %v", variable)
		}
		csp.Constraints[variable] = []Constraint[V, D]{}
	}

	return &csp, nil
}

func (csp *CSPSolver[V, D]) AddConstraint(
	constraint Constraint[V, D],
) error {
	for _, variable := range constraint.Variables() {
		if !slices.Contains(csp.Variables, variable) {
			return fmt.Errorf("unknown variable in constraint: %v", variable)
		}
		csp.Constraints[variable] = append(csp.Constraints[variable], constraint)
	}
	return nil
}

func (csp *CSPSolver[V, D]) Consistent(
	variable V,
	assignment map[V]D,
) bool {
	for _, constraint := range csp.Constraints[variable] {
		if !constraint.Satisfied(assignment) {
			return false
		}
	}
	return true
}

func (csp *CSPSolver[V, D]) BacktrackingSearch(
	assignment map[V]D,
) map[V]D {
	if len(assignment) == len(csp.Variables) {
		return assignment
	}

	localVariables := slices.Clone(csp.Variables)
	unassigned := slices.DeleteFunc(localVariables, func(v V) bool {
		_, ok := assignment[v]
		return ok
	})

	if len(unassigned) == 0 {
		return nil
	}

	first := unassigned[0]
	for _, value := range csp.Domains[first] {
		localAssignment := maps.Clone(assignment)
		localAssignment[first] = value
		if csp.Consistent(first, localAssignment) {
			result := csp.BacktrackingSearch(localAssignment)
			if result != nil {
				return result
			}
		}
	}

	return nil
}
