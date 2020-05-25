package twilio

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
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

func (c *Client) newRequest(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, errors.WithMessage(err, "newRequest()")
	}
	req = req.WithContext(ctx)
	req.SetBasicAuth(c.accountSid, c.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

// DisconnectCall will disconnect the Call associated with callSids
func (c *Client) DisconnectCall(ctx context.Context, callSid string) error {
	ctx, span := trace.StartSpan(ctx, "twilio.Client.DisconnectCall()")
	defer span.End()

	params := make(url.Values)
	params.Add("Status", "completed")
	body := strings.NewReader(params.Encode())

	url := fmt.Sprintf("%s/Accounts/%s/Calls/%s.json", baseURL, c.accountSid, callSid)

	req, err := c.newRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.DisconnectCall()")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.DisconnectCall()")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.WithMessage(decodeError(res.Body), "twilio.Client.DisconnectCall()")
	}

	return nil
}

// SetMute will set the mute state for the Call associated with callSids
func (c *Client) SetMute(ctx context.Context, conferenceSid, callSid string, muted bool) error {
	ctx, span := trace.StartSpan(ctx, "twilio.Client.SetMute()")
	defer span.End()

	params := make(url.Values)
	params.Add("Muted", fmt.Sprintf("%t", muted))
	body := strings.NewReader(params.Encode())

	url := fmt.Sprintf("%s/Accounts/%s/Conferences/%s/Participants/%s.json", baseURL, c.accountSid, conferenceSid, callSid)

	req, err := c.newRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.SetMute()")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithMessage(err, "twilio.Client.SetMute()")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.WithMessage(decodeError(res.Body), "twilio.Client.SetMute()")
	}

	return nil
}

// CallResource receives call resource details
func (c *Client) CallResource(ctx context.Context, callSid string) (*CallResource, error) {
	ctx, span := trace.StartSpan(ctx, "twilio.Client.CallResource()")
	defer span.End()

	url := fmt.Sprintf("%s/Accounts/%s/Calls/%s.json", baseURL, c.accountSid, callSid)

	req, err := c.newRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.CallResource()")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.CallResource(): http.Do(")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.WithMessage(decodeError(res.Body), "twilio.Client.CallResource()")
	}

	callResource := &CallResource{}

	if err := json.NewDecoder(res.Body).Decode(callResource); err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.CallResource(): json.Decoder.Decode()")
	}

	return callResource, nil
}

// Call creates an outbound call returning the resulting CallResource
func (c *Client) Call(ctx context.Context, call *Call) (*CallResource, error) {
	ctx, span := trace.StartSpan(ctx, "twilio.Client.Call()")
	defer span.End()

	params, err := query.Values(call)
	if err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.Call(): query.Values()")
	}

	url := fmt.Sprintf("%s/Accounts/%s/Calls.json", baseURL, c.accountSid)
	body := strings.NewReader(params.Encode())

	req, err := c.newRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.Call()")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.Call(): http.Do(")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, errors.WithMessage(decodeError(res.Body), "twilio.Client.Call()")
	}

	callResource := &CallResource{}

	if err := json.NewDecoder(res.Body).Decode(callResource); err != nil {
		return nil, errors.WithMessage(err, "twilio.Client.Call(): json.Decoder.Decode()")
	}

	return callResource, nil
}
