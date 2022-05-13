package state

import (
	"fmt"
	"start_server/config"
	"start_server/constants"
	"start_server/dosomething"
)

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

type StateMachine struct {
	CurrentState State
	AllStates    map[string]State
	config       config.Configuration
}

func (StateMachine *StateMachine) ProcessStateChange(requestType string) string {
	fmt.Println("Requesting statechange to ", requestType)
	allowedChange := StateMachine.CurrentState.CheckChange(requestType)
	fmt.Println("Checked if allowed ", allowedChange)

	if allowedChange == true {
		StateMachine.CurrentState = StateMachine.AllStates[requestType]
		fmt.Println("State changed to ", StateMachine.CurrentState)
		dosomething.ProcessRequest(StateMachine.CurrentState.Name, StateMachine.config)
		return "Success"
	}

	return "Failed"
}

func NewStateMachine(config config.Configuration) StateMachine {
	statemap := map[string]State{
		constants.First:          {Name: constants.First, availableStateChanges: []string{constants.OptionalSecond, constants.Third}},
		constants.OptionalSecond: {Name: constants.OptionalSecond, availableStateChanges: []string{constants.Third}},
		constants.Third:          {Name: constants.Third, availableStateChanges: []string{constants.OptionalFourth, constants.Fifth}},
		constants.OptionalFourth: {Name: constants.OptionalFourth, availableStateChanges: []string{constants.Fifth}},
		constants.Fifth:          {Name: constants.Fifth, availableStateChanges: []string{constants.First}},
	}

	return StateMachine{CurrentState: statemap[constants.First], AllStates: statemap, config: config}
}
