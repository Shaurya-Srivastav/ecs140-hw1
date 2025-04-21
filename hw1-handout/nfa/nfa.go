package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// If input is empty, check if start and final are the same state
	if len(input) == 0 {
		return start == final
	}

	// Get all possible next states from the current state after reading the first symbol
	nextStates := transitions(start, input[0])

	// If there are no possible next states, return false
	if nextStates == nil {
		return false
	}

	// For each possible next state, recursively check if the final state is reachable
	// using the remaining input
	for _, nextState := range nextStates {
		if Reachable(transitions, nextState, final, input[1:]) {
			return true
		}
	}

	// If none of the paths lead to the final state, return false
	return false
}
