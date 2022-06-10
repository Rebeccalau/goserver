package state

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestState(t *testing.T) {
	suite.Run(t, new(TestStateSuite))
}

type TestStateSuite struct {
	suite.Suite
	state State
}

func (suite *TestStateSuite) SetupTest() {
	suite.state = State{Name: "1", availableStateChanges: []string{"2", "3"}}
}

func (suite *TestStateSuite) TestCheckChangedState() {
	var testCases = []struct {
		input    string
		expected bool
	}{
		{input: "1", expected: false},
		{input: "2", expected: true},
		{input: "3", expected: true},
	}

	for _, testCase := range testCases {
		result := suite.state.CheckChange(testCase.input)

		assert.Equal(suite.T(), testCase.expected, result)
	}
}
