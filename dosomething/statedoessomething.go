package dosomething

import (
	"fmt"
	"start_server/config"
	"start_server/constants"
)

func FirstStateDoesSomething(config config.Configuration) {
	fmt.Println("First State Function Called; Also", config.First)
}

func OtherStates(config config.Configuration) {
	fmt.Println("Other States Functions; Also ", config.Other)
}

func ProcessRequest(requestType string, config config.Configuration) {
	switch requestType {
	case constants.First:
		go FirstStateDoesSomething(config)
	case constants.OptionalSecond, constants.Third, constants.OptionalFourth, constants.Fifth:
		go OtherStates(config)
	default:
		fmt.Println("Error Request")
	}
}
