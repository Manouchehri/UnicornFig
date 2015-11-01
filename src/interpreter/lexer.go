package interpreter

import (
	"errors"
	"regexp"
)

// States for the Parser FSM

type State int

const (
	ERROR   State = iota
	OPEN    State = iota
	LIST    State = iota
	STRING  State = iota
	COMMENT State = iota
	NUMBER  State = iota
	NAME    State = iota
)

type Instruction int

const (
	AddNothing         Instruction = iota
	AddToken           Instruction = iota
	AddChar            Instruction = iota
	AddTokenAndChar    Instruction = iota
	AddTokenAndEndSexp Instruction = iota
)

type RecursiveAction int

const (
	DoNothing RecursiveAction = iota
	Recurse   RecursiveAction = iota
	Return    RecursiveAction = iota
)

type FSMTransition struct {
	ReadMatch string
	WhatToDo  RecursiveAction
	NextState State
	WhatToAdd Instruction
	NewToken  Token
}

var TransitionsFromOpen = [...]FSMTransition{
	{"`", DoNothing, LIST, AddToken, START_LIST},
	{"\\(", Recurse, OPEN, AddToken, START_SEXP},
	{"(\"|')", DoNothing, STRING, AddToken, START_STRING},
	{";", DoNothing, COMMENT, AddToken, START_COMMENT},
	{"[0-9]", DoNothing, NUMBER, AddTokenAndChar, START_NUMBER},
	{"[a-zA-Z]", DoNothing, NAME, AddTokenAndChar, START_NAME},
	{"\\)", Return, OPEN, AddToken, END_SEXP},
}

var TransitionsFromList = [...]FSMTransition{
	{"\\(", Recurse, OPEN, AddNothing, NO_TOKEN},
	{".", DoNothing, ERROR, AddNothing, NO_TOKEN},
}

var TransitionsFromString = [...]FSMTransition{
	{"(\"|')", Return, OPEN, AddToken, END_STRING},
	{".", DoNothing, STRING, AddChar, NO_TOKEN},
}

var TransitionsFromComment = [...]FSMTransition{
	{"\n", Return, OPEN, AddToken, END_COMMENT},
	{".", DoNothing, COMMENT, AddNothing, NO_TOKEN},
}

var TransitionsFromNumber = [...]FSMTransition{
	{"\\s", Return, OPEN, AddToken, END_NUMBER},
	{"\\)", Return, OPEN, AddTokenAndEndSexp, END_NUMBER},
	{"([0-9]|\\.)", DoNothing, NUMBER, AddChar, NO_TOKEN},
}

var TransitionsFromName = [...]FSMTransition{
	{"\\s", Return, OPEN, AddToken, END_NAME},
	{"\\)", Return, OPEN, AddTokenAndEndSexp, END_NAME},
	{"[0-9a-zA-Z_]", DoNothing, NAME, AddChar, NO_TOKEN},
}

func Transition(state State, read string) (error, State, RecursiveAction, []Token) {
	var testTransitions []FSMTransition
	switch state {
	case OPEN:
		testTransitions = TransitionsFromOpen[:]
		break
	case LIST:
		testTransitions = TransitionsFromList[:]
		break
	case STRING:
		testTransitions = TransitionsFromString[:]
		break
	case COMMENT:
		testTransitions = TransitionsFromComment[:]
		break
	case NUMBER:
		testTransitions = TransitionsFromNumber[:]
		break
	case NAME:
		testTransitions = TransitionsFromName[:]
		break
	}
	for _, transition := range testTransitions {
		matched, err := regexp.MatchString(transition.ReadMatch, read)
		if err != nil {
			return err, state, DoNothing, nil
		} else if matched {
			nextState := transition.NextState
			action := transition.WhatToDo
			var tokens []Token
			switch transition.WhatToAdd {
			case AddNothing:
				tokens = []Token{}
			case AddToken:
				tokens = []Token{transition.NewToken}
			case AddChar:
				tokens = []Token{Token(read)}
			case AddTokenAndEndSexp:
				tokens = []Token{transition.NewToken, END_SEXP}
			case AddTokenAndChar:
				tokens = []Token{transition.NewToken, Token(read)}
			}
			return nil, nextState, action, tokens
		}
	}
	// TODO - Provide useful error descriptions
	return errors.New("FILL ME IN"), ERROR, DoNothing, []Token{}
}

func Lex(program string) ([]Token, int) {
	tokens := make([]Token, 0)
	currentState := OPEN
	for i := 0; i < len(program); i++ {
		char := string(program[i])
		err, nextState, action, newTokens := Transition(currentState, string(char))
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, newTokens...)
		if action == Recurse {
			nextTokens, newIndex := Lex(program[i+1:])
			tokens = append(tokens, nextTokens...)
			i = newIndex
		} else if action == Return {
			return tokens, i
		}
		currentState = nextState
	}
	return tokens, len(program)
}
