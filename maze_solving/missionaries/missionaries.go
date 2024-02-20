package missionaries

import (
	"fmt"
	"slices"
)

const MAX_NUM int = 3

type MCState struct {
	WM, WC, EM, EC int
	Boat           bool
}

func NewMCState(wm, wc int, boat bool) MCState {
	return MCState{wm, wc, MAX_NUM - wm, MAX_NUM - wc, boat}
}

func (mcs *MCState) String() string {
	boat_side := "west"
	if !mcs.Boat {
		boat_side = "east"
	}

	return fmt.Sprintf(
		"On the west bank there are %d missionaries and %d cannibals.\n"+
			"On the east bank there are %d missionaries and %d cannibals.\n"+
			"The boat is on the %s bank.\n",
		mcs.WM, mcs.WC, mcs.EM, mcs.EC, boat_side,
	)
}

func (mcs MCState) isLegal() bool {
	if mcs.WM > 0 && mcs.WM < mcs.WC {
		return false
	}
	if mcs.EM > 0 && mcs.EM < mcs.EC {
		return false
	}
	return true
}

func GoalTest(mcs MCState) bool {
	return mcs.isLegal() && mcs.EM == MAX_NUM && mcs.EC == MAX_NUM
}

func Successors(mcs MCState) []MCState {
	successors := make([]MCState, 0)

	if mcs.Boat {
		if mcs.WM > 1 {
			successors = append(successors, NewMCState(mcs.WM-2, mcs.WC, !mcs.Boat))
		}
		if mcs.WM > 0 {
			successors = append(successors, NewMCState(mcs.WM-1, mcs.WC, !mcs.Boat))
		}
		if mcs.WC > 1 {
			successors = append(successors, NewMCState(mcs.WM, mcs.WC-2, !mcs.Boat))
		}
		if mcs.WC > 0 {
			successors = append(successors, NewMCState(mcs.WM, mcs.WC-1, !mcs.Boat))
		}
		if mcs.WM > 0 && mcs.WC > 0 {
			successors = append(successors, NewMCState(mcs.WM-1, mcs.WC-1, !mcs.Boat))
		}
	} else {
		if mcs.EM > 1 {
			successors = append(successors, NewMCState(mcs.WM+2, mcs.WC, !mcs.Boat))
		}
		if mcs.EM > 0 {
			successors = append(successors, NewMCState(mcs.WM+1, mcs.WC, !mcs.Boat))
		}
		if mcs.EC > 1 {
			successors = append(successors, NewMCState(mcs.WM, mcs.WC+2, !mcs.Boat))
		}
		if mcs.EC > 0 {
			successors = append(successors, NewMCState(mcs.WM, mcs.WC+1, !mcs.Boat))
		}
		if mcs.EM > 0 && mcs.EC > 0 {
			successors = append(successors, NewMCState(mcs.WM+1, mcs.WC+1, !mcs.Boat))
		}
	}
	successors = slices.DeleteFunc(successors, func(mcs MCState) bool {
		return !mcs.isLegal()
	})
	return successors
}

func DisplaySolution(path []MCState) {
	if len(path) == 0 {
		return
	}

	oldState := path[0]
	println(oldState.String())

	for _, currentState := range path[1:] {
		if currentState.Boat {
			fmt.Printf(
				"Move %d missionaries and %d cannibals from the east bank to the west bank.\n",
				oldState.EM-currentState.EM, oldState.EC-currentState.EC,
			)
		} else {
			fmt.Printf(
				"Move %d missionaries and %d cannibals from the west bank to the east bank.\n",
				oldState.WM-currentState.WM, oldState.WC-currentState.WC,
			)
		}
		println(currentState.String())
		oldState = currentState
	}
}
