package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (serverState ServerState) String() string {
	return stateName[serverState]
}

func main() {
	initialState := transition(StateIdle)
	fmt.Println(initialState)

	finalState := transition(initialState)
	fmt.Println(finalState)
}

func transition(serverState ServerState) ServerState {
	switch serverState {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", serverState))
	}
}
