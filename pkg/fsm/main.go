package fsm

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type states int

const (
	initial states = iota
	listAll
	viewOne
	quitNow
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

func Run() {
	setup()

	for {
		if state == quitNow {
			return
		}
		fmt.Println(helpMessages[state])
		input := bufio.NewReader(os.Stdin)
		command, _ := input.ReadString('\n')
		if err := execute(command); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func execute(command string) error {

	// extract first word
	cmd := regexp.MustCompile(`[a-z]+`).FindString(command)
	if _, valid := commands[state][cmd]; !valid {
		fmt.Printf("Error: command unsupported: \"%s\"", cmd)
		time.Sleep(1 * time.Second)
		return nil
	}

	switch state {
	case initial:
		switch cmd {
		case "list":
			if err := tt.list(); err != nil {
				return err
			}
			state = listAll
		case "quit":
			tt.quit()
			state = quitNow
		}
	case listAll:
		switch cmd {
		case "prev":
			if err := tt.prev(); err != nil {
				return err
			}
		case "next":
			if err := tt.next(); err != nil {
				return err
			}
		case "selc":
			str := regexp.MustCompile(`[1-9]\d*`).FindString(command)
			num, _ := strconv.Atoi(str)
			if valid := tt.selc(num); valid {
				state = viewOne
			} else {
				time.Sleep(2 * time.Second)
			}
		case "quit":
			tt.quit()
			state = quitNow
		}
	case viewOne:
		switch cmd {
		case "back":
			tt.back()
			state = listAll
		case "quit":
			tt.quit()
			state = quitNow
		}
	}

	return nil
}
