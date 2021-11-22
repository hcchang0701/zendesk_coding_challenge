package fsm

import (
	"fmt"
	"time"
)

type ticketResp struct {
	Tickets []*Ticket `json:"tickets"`
	Meta    `json:"meta"`
}

type Ticket struct {
	URL        string      `json:"url"`
	ID         int         `json:"id"`
	ExternalID interface{} `json:"external_id"`
	Via        struct {
		Channel string `json:"channel"`
		Source  struct {
			From interface{} `json:"from"`
			To   interface{} `json:"to"`
			Rel  interface{} `json:"rel"`
		} `json:"source"`
	} `json:"via"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	Type                interface{}   `json:"type"`
	Subject             string        `json:"subject"`
	RawSubject          string        `json:"raw_subject"`
	Description         string        `json:"description"`
	Priority            interface{}   `json:"priority"`
	Status              string        `json:"status"`
	Recipient           interface{}   `json:"recipient"`
	RequesterID         int64         `json:"requester_id"`
	SubmitterID         int64         `json:"submitter_id"`
	AssigneeID          int64         `json:"assignee_id"`
	OrganizationID      int64         `json:"organization_id"`
	GroupID             int64         `json:"group_id"`
	CollaboratorIds     []interface{} `json:"collaborator_ids"`
	FollowerIds         []interface{} `json:"follower_ids"`
	EmailCcIds          []interface{} `json:"email_cc_ids"`
	ForumTopicID        interface{}   `json:"forum_topic_id"`
	ProblemID           interface{}   `json:"problem_id"`
	HasIncidents        bool          `json:"has_incidents"`
	IsPublic            bool          `json:"is_public"`
	DueAt               interface{}   `json:"due_at"`
	Tags                []string      `json:"tags"`
	CustomFields        []interface{} `json:"custom_fields"`
	SatisfactionRating  interface{}   `json:"satisfaction_rating"`
	SharingAgreementIds []interface{} `json:"sharing_agreement_ids"`
	FollowupIds         []interface{} `json:"followup_ids"`
	TicketFormID        int64         `json:"ticket_form_id"`
	BrandID             int64         `json:"brand_id"`
	AllowChannelback    bool          `json:"allow_channelback"`
	AllowAttachments    bool          `json:"allow_attachments"`
}

func (t *Ticket) String() string {

	str := fmt.Sprintf("#%d %s [stat: %s]", t.ID, t.Subject, t.Status)

	if t.Type != nil {
		str += fmt.Sprintf("[type: %v]", t.Type)
	}
	if t.Priority != nil {
		str += fmt.Sprintf("[pri: %v]", t.Priority)
	}

	return str + "\n"
}

type Meta struct {
	HasMore      bool   `json:"has_more"`
	AfterCursor  string `json:"after_cursor"`
	BeforeCursor string `json:"before_cursor"`
}
