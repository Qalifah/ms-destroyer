package destroyer

import "time"

type Target struct {
	Id        string     `json:"id" fake:"{uuid}"`
	Message   string     `json:"message" fake:"{sentence:7}"`
	CreatedOn time.Time  `json:"created_on"`
	UpdatedOn *time.Time `json:"updated_on"`
}

type Event struct {
	Id        string    `json:"id" fake:"{uuid}"`
	Name      string    `json:"name" fake:"skip"`
	Targets   []Target  `json:"targets" fakesize:"5"`
	CreatedOn time.Time `json:"created_on"`
}
