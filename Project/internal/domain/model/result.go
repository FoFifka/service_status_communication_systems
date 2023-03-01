package model

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       []SMSData       `json:"sms"`
	MMS       []MMSData       `json:"mms"`
	VoiceCall []VoiceCallData `json:"voice_call"`
	Email     []EmailData     `json:"email"`
	Billing   BillingData     `json:"billing"`
	Support   []SupportData   `json:"support"`
	Incidents []IncidentData  `json:"incidents"`
}
