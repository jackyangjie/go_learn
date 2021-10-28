package trs

type Argument struct {
	Role string `json:"role"`
	Argument string `json:"argument"`
}

type Trigger struct {
	Trigger string `json:"trigger"`
	Arguments [] Argument `json:"arguments"`
}

type Recode struct {
	Text string `json:"text"`
	Event_list []Trigger `json:"event_list"`
}
