package gogovtrack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DefaultBaseURL is where govtrack expects its api calls
const DefaultBaseURL = "https://www.govtrack.us/api/v2"

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	OriginalBody string
	Code         int
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Invalid status code: %d", e.Code)
}

// ErrNotExpectedJSON is returned by API calls when the response isn't expected JSON
type ErrNotExpectedJSON struct {
	OriginalBody string
	Err          error
}

func (e *ErrNotExpectedJSON) Error() string {
	return fmt.Sprintf("Unexpected JSON: %s from %s", e.Err.Error(), e.OriginalBody)
}

// Client is the base for interacting with the go govtrack API.
type Client struct {
	Client  *http.Client
	BaseURL string
	filters map[string]string
}

func (c *Client) request(u string, jsonResp interface{}) error {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return err
	}
	debug := b.String()

	if resp.StatusCode != http.StatusOK {
		return &ErrBadStatusCode{
			OriginalBody: debug,
			Code:         resp.StatusCode,
		}
	}

	if err := json.NewDecoder(&b).Decode(jsonResp); err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: debug,
			Err:          err,
		}
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

// GetBill is
func (c *Client) GetBill(id int) (Bill, error) {
	var b Bill
	if err := c.request(c.urlBase("bill"), &b); err != nil {
		return b, err
	}
	return b, nil
}

// GetBills is
func (c *Client) GetBills(id int) ([]Bill, error) {
	var b []Bill
	if err := c.request(c.urlBase("bill"), &b); err != nil {
		return b, err
	}
	return b, nil
}

// GetCommittee is
func (c *Client) GetCommittee(id int) (Committee, error) {
	var cm Committee
	if err := c.request(c.urlBase("committe"), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetCommittees is
func (c *Client) GetCommittees() ([]Committee, error) {
	var cm []Committee
	if err := c.request(c.urlBase("committe"), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetRole is
func (c *Client) GetRole(id int) (Role, error) {
	var r Role
	if err := c.request(c.urlBase("role"), &r); err != nil {
		return r, err
	}
	return r, nil
}

// GetRoles is
func (c *Client) GetRoles() ([]Role, error) {
	var r []Role
	if err := c.request(c.urlBase("role"), &r); err != nil {
		return r, err
	}
	return r, nil
}

// GetPerson is
func (c *Client) GetPerson(id int) (Person, error) {
	var p Person
	if err := c.request(c.urlBase("person"), &p); err != nil {
		return p, err
	}
	return p, nil
}

// GetPersons is
func (c *Client) GetPersons() ([]Person, error) {
	var p []Person
	if err := c.request(c.urlBase("person"), &p); err != nil {
		return p, err
	}
	return p, nil
}

// GetCommitteeMember is
func (c *Client) GetCommitteeMember(id int) (CommitteeMember, error) {
	var cm CommitteeMember
	if err := c.request(c.urlBase("committe_member"), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetCommitteeMembers is
func (c *Client) GetCommitteeMembers() ([]CommitteeMember, error) {
	var p []CommitteeMember
	if err := c.request(c.urlBase("committe_member"), &p); err != nil {
		return p, err
	}
	return p, nil
}

// Filter is
func (c *Client) Filter(query map[string]string) *Client {
	c.filters = query
	return c
}

func (c *Client) urlBase(resource string) string {
	base := c.BaseURL
	if c.BaseURL == "" {
		base = DefaultBaseURL
	}
	return fmt.Sprintf("%s/%s", base, resource)
}
