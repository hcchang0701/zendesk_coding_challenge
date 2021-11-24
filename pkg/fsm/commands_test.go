package fsm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTicketFetcher struct{}

var _ ITicketFetcher = (*mockTicketFetcher)(nil)

var get func(string, string) ([]*Ticket, *Meta, error)

func (m *mockTicketFetcher) getTicketsWithCursor(before, after string) ([]*Ticket, *Meta, error) {
	return get(before, after)
}

func TestListCommandSuccess(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	get = func(string, string) ([]*Ticket, *Meta, error) {
		return []*Ticket{
			{
				ID:      1,
				Subject: "First ticket",
			},
		}, nil, nil
	}
	err := tt.list()
	assert.NoError(t, err)
}

func TestListCommandFailed(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	get = func(string, string) ([]*Ticket, *Meta, error) {
		return nil, nil, errors.New("Something's wrong")
	}
	err := tt.list()
	assert.Error(t, err)
}

func TestPrevCommandSuccess(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	meta = &Meta{HasMore: true}
	get = func(string, string) ([]*Ticket, *Meta, error) {
		return []*Ticket{
			{
				ID:      1,
				Subject: "First ticket",
			},
		}, nil, nil
	}
	err := tt.prev()
	assert.NoError(t, err)
}

func TestPrevCommandFailed(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	meta = &Meta{HasMore: true}
	get = func(string, string) ([]*Ticket, *Meta, error) {
		return nil, nil, errors.New("Something's wrong")
	}
	err := tt.prev()
	assert.Error(t, err)
}

func TestNextCommandSuccess(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	meta = &Meta{HasMore: true}
	get = func(string, string) ([]*Ticket, *Meta, error) {
		return []*Ticket{
			{
				ID:      1,
				Subject: "First ticket",
			},
		}, nil, nil
	}
	err := tt.next()
	assert.NoError(t, err)
}

func TestNextCommandFailed(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	meta = &Meta{HasMore: true}
	get = func(string, string) ([]*Ticket, *Meta, error) {
		return nil, nil, errors.New("Something's wrong")
	}
	err := tt.next()
	assert.Error(t, err)
}

func TestSelcCommandSuccess(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	tickets = []*Ticket{
		{
			ID: 1,
		},
	}
	res := tt.selc(1)
	assert.Equal(t, true, res)
}

func TestSelcCommandFailed(t *testing.T) {
	tf = new(mockTicketFetcher)
	tt = new(transition)
	tickets = []*Ticket{
		{
			ID: 1,
		},
	}
	res := tt.selc(2)
	assert.Equal(t, false, res)
}
