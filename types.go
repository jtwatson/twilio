package twilio

// CallResource holds the details of a call resouce
type CallResource struct {
	Sid             string          `json:"sid,omitempty"`
	AccountSid      string          `json:"account_sid,omitempty"`
	Annotation      string          `json:"annotation,omitempty"`
	AnsweredBy      string          `json:"answered_by,omitempty"`
	APIVersion      string          `json:"api_version,omitempty"`
	CallerName      string          `json:"caller_name,omitempty"`
	DateCreated     string          `json:"date_created,omitempty"`
	DateUpdated     string          `json:"date_updated,omitempty"`
	Direction       string          `json:"direction,omitempty"`
	Duration        string          `json:"duration,omitempty"`
	EndTime         string          `json:"end_time,omitempty"`
	ForwardedFrom   string          `json:"forwarded_from,omitempty"`
	From            string          `json:"from,omitempty"`
	FromFormatted   string          `json:"from_formatted,omitempty"`
	GroupSid        string          `json:"group_sid,omitempty"`
	ParentCallSid   string          `json:"parent_call_sid,omitempty"`
	PhoneNumberSid  string          `json:"phone_number_sid,omitempty"`
	Price           string          `json:"price,omitempty"`
	PriceUnit       string          `json:"price_unit,omitempty"`
	QueueTime       string          `json:"queue_time,omitempty"`
	StartTime       string          `json:"start_time,omitempty"`
	Status          string          `json:"status,omitempty"`
	SubresourceUris SubresourceUris `json:"subresource_uris,omitempty"`
	To              string          `json:"to,omitempty"`
	ToFormatted     string          `json:"to_formatted,omitempty"`
	TrunkSid        string          `json:"trunk_sid,omitempty"`
	URI             string          `json:"uri,omitempty"`
}

// SubresourceUris holds details for subresource uri's
type SubresourceUris struct {
	Notifications     string `json:"notifications,omitempty"`
	Recordings        string `json:"recordings,omitempty"`
	Feedback          string `json:"feedback,omitempty"`
	FeedbackSummaries string `json:"feedback_summaries,omitempty"`
	Payments          string `json:"payments,omitempty"`
}

// Call describes a outgoing call settings
type Call struct {
	AccountSid                         string `url:"AccountSid,omitempty"`
	ApplicationSid                     string `url:"ApplicationSid,omitempty"`
	AsyncAmd                           string `url:"AsyncAmd,omitempty"`
	AsyncAmdStatusCallback             string `url:"AsyncAmdStatusCallback,omitempty"`
	AsyncAmdStatusCallbackMethod       string `url:"AsyncAmdStatusCallbackMethod,omitempty"`
	Byoc                               string `url:"Byoc,omitempty"`
	CallerId                           string `url:"CallerId,omitempty"`
	CallReason                         string `url:"CallReason,omitempty"`
	FallbackMethod                     string `url:"FallbackMethod,omitempty"`
	FallbackUrl                        string `url:"FallbackUrl,omitempty"`
	From                               string `url:"From,omitempty"`
	MachineDetection                   string `url:"MachineDetection,omitempty"`
	MachineDetectionSilenceTimeout     int    `url:"MachineDetectionSilenceTimeout,omitempty"`
	MachineDetectionSpeechEndThreshold int    `url:"MachineDetectionSpeechEndThreshold,omitempty"`
	MachineDetectionSpeechThreshold    int    `url:"MachineDetectionSpeechThreshold,omitempty"`
	MachineDetectionTimeout            int    `url:"MachineDetectionTimeout,omitempty"`
	Method                             string `url:"Method,omitempty"`
	Record                             bool   `url:"Record,omitempty"`
	RecordingChannels                  string `url:"RecordingChannels,omitempty"`
	RecordingStatusCallback            string `url:"RecordingStatusCallback,omitempty"`
	RecordingStatusCallbackEvent       string `url:"RecordingStatusCallbackEvent,omitempty"`
	RecordingStatusCallbackMethod      string `url:"RecordingStatusCallbackMethod,omitempty"`
	SendDigits                         string `url:"SendDigits,omitempty"`
	SipAuthPassword                    string `url:"SipAuthPassword,omitempty"`
	SipAuthUsername                    string `url:"SipAuthUsername,omitempty"`
	StatusCallback                     string `url:"StatusCallback,omitempty"`
	StatusCallbackEvent                string `url:"StatusCallbackEvent,omitempty"`
	StatusCallbackMethod               string `url:"StatusCallbackMethod,omitempty"`
	Timeout                            int    `url:"Timeout,omitempty"`
	To                                 string `url:"To,omitempty"`
	Trim                               string `url:"Trim,omitempty"`
	Twiml                              string `url:"Twiml,omitempty"`
	URL                                string `url:"Url,omitempty"`
}
