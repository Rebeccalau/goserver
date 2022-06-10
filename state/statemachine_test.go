package state

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"start_server/config"
	"start_server/constants"
	"testing"
)

func TestStateMachine(t *testing.T) {
	suite.Run(t, new(TestStateMachineSuite))
}

type TestStateMachineSuite struct {
	suite.Suite
	TestMachine StateMachine
}

type MockChecker struct{}

func (m *MockChecker) CheckChange(nextState string, state State) bool {
	return true
}

func (suite *TestStateMachineSuite) SetupTest() {
	config := config.Configuration{}
	statemap := map[string]State{
		constants.First:          {Name: constants.First, availableStateChanges: []string{constants.OptionalSecond, constants.Third}},
		constants.OptionalSecond: {Name: constants.OptionalSecond, availableStateChanges: []string{constants.Third}},
		constants.Third:          {Name: constants.Third, availableStateChanges: []string{constants.First}},
	}

	suite.TestMachine = StateMachine{CurrentState: statemap[constants.First], AllStates: statemap, config: config, Checker: &MockChecker{}}
}

func (suite *TestStateMachineSuite) TestProcessStateChange() {
	result := suite.TestMachine.ProcessStateChange("optionalsecond")

	assert.Equal(suite.T(), "Success", result)
}
