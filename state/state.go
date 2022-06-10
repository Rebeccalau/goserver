package state

type State struct {
	Name                  string
	availableStateChanges []string
}

func (state *State) CheckChange(nextState string) bool {
	for _, x := range state.availableStateChanges {
		if nextState == x {
			return true
		}
	}
	return false
}
