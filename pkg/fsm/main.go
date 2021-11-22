package fsm

import (
	"fmt"
	"regexp"
	"time"
)

type states int

const (
	initial states = iota
	listAll
	viewOne
)

var commands = map[states]map[string]struct{}{
	// states -> valid commands
	initial: {"list": struct{}{}, "quit": struct{}{}},
	listAll: {"next": struct{}{}, "prev": struct{}{}, "selc": struct{}{}, "quit": struct{}{}},
	viewOne: {"back": struct{}{}, "quit": struct{}{}},
}

var helpMessages = map[states]string{
	initial: "\nWelcome to TicketViewer!!!\n" +
		"To start with, type 'list' to list all of your tickets or enter 'quit' to leave:\n",
	listAll: "\nPlease use one of the following commands to move on:\n" +
		"\tprev        Go to the previous page\n" +
		"\tnext        Go to the next page\n" +
		"\tselc num    View more details about Ticket #{num}\n" +
		"\tquit        Leave this program and say goodbye\n",
	viewOne: "\nType 'back' to go back to the list view, otherwise enter 'quit' to leave:\n",
}

var state states

func init() {
	state = initial
}

func Run() {

	for {
		// print some words
		fmt.Println(helpMessages[state])

		var command string
		//command = "list"
		fmt.Scan(&command)

		if err := execute(command); err != nil {
			return
		}
	}
}

func execute(command string) error {

	// normalize
	re := regexp.MustCompile(`[a-z]+`)
	cmd := re.FindString(command)

	if _, valid := commands[state][cmd]; !valid {
		fmt.Printf("Error: command unsupported: \"%s\"", cmd)
		time.Sleep(2 * time.Second)
		return nil
	}

	switch state {
	case initial:
		switch cmd {
		case "list":
			list()
			state = listAll
		case "quit":
			quit()
		}
	case listAll:
		switch cmd {
		case "prev":
			prev()
		case "next":
			next()
		case "selc":
			state = viewOne
		case "quit":
			quit()
		}
	case viewOne:
		switch cmd {
		case "back":
			list()
			state = listAll
		case "quit":
			quit()
		}
	}

	return nil
}
