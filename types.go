package twilio

import (
	"fmt"
	"strings"
	"time"
)

// CallResource holds the details of a call resouce
type CallResource struct {
	Sid             string          `json:"sid,omitempty"`
	AccountSid      string          `json:"account_sid,omitempty"`
	Annotation      string          `json:"annotation,omitempty"`
	AnsweredBy      string          `json:"answered_by,omitempty"`
	APIVersion      string          `json:"api_version,omitempty"`
	CallerName      string          `json:"caller_name,omitempty"`
	DateCreated     TwilioTime      `json:"date_created,omitempty"`
	DateUpdated     TwilioTime      `json:"date_updated,omitempty"`
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

// Call describes outgoing call settings
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

// ConferenceResource holds the details of a conference
type ConferenceResource struct {
	AccountSid              string            `json:"account_sid,omitempty"`
	DateCreated             TwilioTime        `json:"date_created,omitempty"`
	DateUpdated             TwilioTime        `json:"date_updated,omitempty"`
	ApiVersion              string            `json:"api_version,omitempty"`
	FriendlyName            string            `json:"friendly_name,omitempty"`
	Region                  string            `json:"region,omitempty"`
	Sid                     string            `json:"sid,omitempty"`
	Status                  string            `json:"status,omitempty"` // init, in-progress, or completed.
	Uri                     string            `json:"uri,omitempty"`
	SubresourceUris         map[string]string `json:"subresource_uris,omitempty"`
	ReasonConferenceEnded   string            `json:"reason_conference_ended,omitempty"` // conference-ended-via-api, participant-with-end-conference-on-exit-left, participant-with-end-conference-on-exit-kicked, last-participant-kicked, or last-participant-left.
	CallSidEndingConference string            `json:"call_sid_ending_conference,omitempty"`
}

// ParticipantResource holds the details of a participant
type ParticipantResource struct {
	AccountSid             string     `json:"account_sid,omitempty"`
	CallSid                string     `json:"call_sid,omitempty"`
	CallSidToCoach         string     `json:"call_sid_to_coach,omitempty"`
	Coaching               bool       `json:"coaching,omitempty"`
	ConferenceSid          string     `json:"conference_sid,omitempty"`
	DateCreated            TwilioTime `json:"date_created,omitempty"`
	DateUpdated            TwilioTime `json:"date_updated,omitempty"`
	EndConferenceOnExit    bool       `json:"end_conference_on_exit,omitempty"`
	Muted                  bool       `json:"muted,omitempty"`
	Hold                   bool       `json:"hold,omitempty"`
	StartConferenceOnEnter bool       `json:"start_conference_on_enter,omitempty"`
	Status                 string     `json:"status,omitempty"` // queued, connecting, ringing, connected, complete, or failed
	Uri                    string     `json:"uri,omitempty"`
}

// APIError holds the details of errors returned from twilio
type APIError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

// Error returns string representation of the error
func (a *APIError) Error() string {
	return fmt.Sprintf("APIError: %s: more_info: %s", a.Message, a.MoreInfo)
}

// TwilioTime implements interfaces for json Marshalling and Unmarshalling
type TwilioTime struct {
	time.Time
}

const ttLayout = "Mon, 02 Jan 2006 15:04:05 -0700" // 2006/01/02|15:04:05

// UnmarshalJSON implements the Unmarshaler interface
func (tt *TwilioTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		tt.Time = time.Time{}
		return
	}
	tt.Time, err = time.Parse(ttLayout, s)
	return
}

// MarshslJSON implements the Marshaler interface
func (tt *TwilioTime) MarshalJSON() ([]byte, error) {
	if tt.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tt.Time.Format(ttLayout))), nil
}
