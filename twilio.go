package twilio

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

var baseURL = "https://api.twilio.com/2010-04-01"

// Client is an http client that talks to Twilio's Rest API
type Client struct {
	httpClient *http.Client
	accountSid string
	authToken  string
}

// New returns a new Twillio Client
func New(client *http.Client, accountSid, authToken string) *Client {
	return &Client{httpClient: client, accountSid: accountSid, authToken: authToken}
}

func (c *Client) newRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, errors.WithMessage(err, "newRequest()")
	}
	req.SetBasicAuth(c.accountSid, c.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

// DisconnectCall will disconnect the Call associated with callSids
func (c *Client) DisconnectCall(callSid string) error {
	params := make(url.Values)
	params.Add("Status", "completed")
	body := strings.NewReader(params.Encode())

	url := fmt.Sprintf("%s/Accounts/%s/Calls/%s.json", baseURL, c.accountSid, callSid)

	req, err := c.newRequest(http.MethodPost, url, body)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.DisconnectCall()")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.DisconnectCall()")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("twilio.Client.DisconnectCall(): expected status code 200, got %d", res.StatusCode)
	}

	return nil
}

// SetMute will set the mute state for the Call associated with callSids
func (c *Client) SetMute(conferenceSid, callSid string, muted bool) error {
	params := make(url.Values)
	params.Add("Muted", fmt.Sprintf("%t", muted))
	body := strings.NewReader(params.Encode())

	url := fmt.Sprintf("%s/Accounts/%s/Conferences/%s/Participants/%s.json", baseURL, c.accountSid, conferenceSid, callSid)

	req, err := c.newRequest(http.MethodPost, url, body)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.SetMute()")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.SetMute()")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("twilio.Client.SetMute(): expected status code 200, got %d", res.StatusCode)
	}

	return nil
}

// CallResource recieves call resource details
func (c *Client) CallResource(callSid string) {

}

// CallResource holds the details of a call resouce
type CallResource struct {
	Sid             string          `json:"sid,omitempty"`
	DateCreated     string          `json:"date_created,omitempty"`
	DateUpdated     string          `json:"date_updated,omitempty"`
	ParentCallSid   string          `json:"parent_call_sid,omitempty"`
	AccountSid      string          `json:"account_sid,omitempty"`
	To              string          `json:"to,omitempty"`
	From            string          `json:"from,omitempty"`
	PhoneNumberSid  string          `json:"phone_number_sid,omitempty"`
	Status          string          `json:"status,omitempty"`
	StartTime       string          `json:"start_time,omitempty"`
	EndTime         string          `json:"end_time,omitempty"`
	Duration        string          `json:"duration,omitempty"`
	Price           string          `json:"price,omitempty"`
	Direction       string          `json:"direction,omitempty"`
	AnsweredBy      string          `json:"answered_by,omitempty"`
	APIVersion      string          `json:"api_version,omitempty"`
	ForwardedFrom   string          `json:"forwarded_from,omitempty"`
	CallerName      string          `json:"caller_name,omitempty"`
	URI             string          `json:"uri,omitempty"`
	SubresourceUris SubresourceUris `json:"subresource_uris,omitempty"`
}

type SubresourceUris struct {
	Notifications string `json:"notifications,omitempty"`
	Recordings    string `json:"recordings,omitempty"`
}
