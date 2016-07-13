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

// Q is
type Q map[string]string

// Meta is
type Meta struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

// Client is the base for interacting with the go govtrack API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	filters    map[string]string
}

func (c *Client) request(u string, jsonResp interface{}) error {
	fullURL := fmt.Sprintf("%s%s", u, c.buildFilterQuery())
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
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
			OriginalBody: "",
			Err:          err,
		}
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

// GetBill is
func (c *Client) GetBill(id int) (BillResponse, error) {
	var b BillResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("bill"), id), &b); err != nil {
		return b, err
	}
	return b, nil
}

// GetBills is
func (c *Client) GetBills() (BillsResponse, error) {
	var b BillsResponse
	if err := c.request(c.urlBase("bill"), &b); err != nil {
		return b, err
	}
	return b, nil
}

// GetCommittee is
func (c *Client) GetCommittee(id int) (CommitteeResponse, error) {
	var cm CommitteeResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("committee"), id), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetCommittees is
func (c *Client) GetCommittees() (CommitteesResponse, error) {
	var cm CommitteesResponse
	if err := c.request(c.urlBase("committe"), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetRole is
func (c *Client) GetRole(id int) (RoleResponse, error) {
	var r RoleResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("role"), id), &r); err != nil {
		return r, err
	}
	return r, nil
}

// GetRoles is
func (c *Client) GetRoles() (RolesResponse, error) {
	var r RolesResponse
	if err := c.request(c.urlBase("role"), &r); err != nil {
		return r, err
	}
	return r, nil
}

// GetPerson is
func (c *Client) GetPerson(id int) (PersonResponse, error) {
	var p PersonResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("person"), id), &p); err != nil {
		return p, err
	}
	return p, nil
}

// GetPersons is
func (c *Client) GetPersons() (PersonsResponse, error) {
	var p PersonsResponse
	if err := c.request(c.urlBase("person"), &p); err != nil {
		return p, err
	}
	return p, nil
}

// GetCommitteeMember is
func (c *Client) GetCommitteeMember(id int) (CommitteeMemberResponse, error) {
	var cm CommitteeMemberResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("committee_member"), id), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetCommitteeMembers is
func (c *Client) GetCommitteeMembers() (CommitteeMembersResponse, error) {
	var cm CommitteeMembersResponse
	if err := c.request(c.urlBase("committee_member"), &cm); err != nil {
		return cm, err
	}
	return cm, nil
}

// GetCosponsorship is
func (c *Client) GetCosponsorship(id int) (CosponsorshipResponse, error) {
	var cp CosponsorshipResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("cosponsorship"), id), &cp); err != nil {
		return cp, err
	}
	return cp, nil
}

// GetCosponsorships is
func (c *Client) GetCosponsorships() (CosponsorshipsResponse, error) {
	var cp CosponsorshipsResponse
	if err := c.request(c.urlBase("cosponsorship"), &cp); err != nil {
		return cp, err
	}
	return cp, nil
}

// GetVoteVoter is
func (c *Client) GetVoteVoter(id int) (VoteVoterResponse, error) {
	var vv VoteVoterResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("vote_voter"), id), &vv); err != nil {
		return vv, err
	}
	return vv, nil
}

// GetVoteVoters is
func (c *Client) GetVoteVoters() (VoteVotersResponse, error) {
	var vv VoteVotersResponse
	if err := c.request(c.urlBase("vote_voter"), &vv); err != nil {
		return vv, err
	}
	return vv, nil
}

// GetVote is
func (c *Client) GetVote(id int) (VoteResponse, error) {
	var v VoteResponse
	if err := c.request(fmt.Sprintf("%s/%d", c.urlBase("vote"), id), &v); err != nil {
		return v, err
	}
	return v, nil
}

// GetVotes is
func (c *Client) GetVotes() (VotesResponse, error) {
	var v VotesResponse
	if err := c.request(c.urlBase("vote"), &v); err != nil {
		return v, err
	}
	return v, nil
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

func (c *Client) buildFilterQuery() string {
	// Check to see if we even have filters to build
	if len(c.filters) < 1 {
		return ""
	}

	isFirst := true
	query := "?"

	for k, v := range c.filters {
		if !isFirst {
			query += "&" + k + "=" + v
			continue
		}

		isFirst = false
		query += k + "=" + v
	}

	return query
}
