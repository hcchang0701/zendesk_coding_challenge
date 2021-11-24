package fsm

type ITicketFetcher interface {
	getTicketsWithCursor(string, string) ([]*Ticket, *Meta, error)
}

type ITransitionFunctions interface {
	list() error
	next() error
	prev() error
	selc(int) bool
	back()
	quit()
}
