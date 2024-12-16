package operation

import "time"

type OperationNature int8

const (
	OP_SURVEY_CLOSE OperationNature = iota + 1
	OP_SURVEY_OPEN
	OP_SURVEY_CANCEL
	OP_SURVEY_DISCARD
	OP_SURVEY_APPROVE
	OP_SURVEY_RESULT
	OP_SURVEY_LIST
	OP_ITEM_AUTHENTICATE
	OP_ITEM_WEIGH
	OP_VOTE_STARTED
	OP_VOTE_RECEIVE
	OP_VOTE_ACCEPTED
	OP_VOTE_DENIED
	OP_VOTE_DEF_MESSAGE
)

type OpError struct {
	Microservice string `json:"ms_id"`
	At           uint64 `json:"at"`
	Error        string `json:"error"`
}

type Operation struct {
	Id          string    `json:"id"`
	RequestedBy string    `json:"requested_by"`
	HandledBy   string    `json:"handled_by"`
	StartedAt   time.Time `json:"started_at"`
	LastUpdate  time.Time `json:"last_update"`
	ErrorStack  []OpError `json:"error_stack"`
	Path        []string  `json:"path"`
}
