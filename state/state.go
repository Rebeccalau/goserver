package state

type State struct {
	Name                  string
	availableStateChanges []string
}

type ICheck interface {
	CheckChange(nextState string, state State) bool
}

type Checker struct {
}

func (c *Checker) CheckChange(nextState string, state State) bool {
	for _, x := range state.availableStateChanges {
		if nextState == x {
			return true
		}
	}
	return false
}
