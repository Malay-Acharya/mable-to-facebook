package models

type FacebookData struct {
	Data          []FacebookEvent `json:"data"`
	TestEventCode string          `json:"test_event_code"`
}

type FacebookEvent struct {
	EventName    string `json:"event_name"`
	EventTime    int64  `json:"event_time"`
	ActionSource string `json:"action_source"`
	UserData     struct {
		Em []string `json:"em"`
		Ph []string `json:"ph"`
	} `json:"user_data"`
	CustomData struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	} `json:"custom_data"`
}
