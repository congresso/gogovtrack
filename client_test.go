package gogovtrack

import (
	"net/http"
	"os"
	"testing"
)

var (
	testClient = new(Client)
)

func TestMain(m *testing.M) {
	testClient = &Client{
		HTTPClient: http.DefaultClient,
	}

	os.Exit(m.Run())
}

func TestGetBill(t *testing.T) {
	b, err := testClient.GetBill(11586)
	if err != nil {
		t.Error(err)
	}

	if b.BillID != 11586 {
		t.Errorf("Expected bill to id to equal %d but got %d", 11586, b.BillID)
	}
}

func TestGetBills(t *testing.T) {
	bs, err := testClient.GetBills()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Bills) < 1 {
		t.Error("Expected at least one bill to be returned")
	}
}

func TestGetFilteredBills(t *testing.T) {
	bs, err := testClient.Filter(Q{"limit": "5"}).GetBills()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Bills) != 5 {
		t.Error("Expected filtered bill to return 5 bills")
	}
}

func TestGetPerson(t *testing.T) {
	b, err := testClient.GetPerson(300001)
	if err != nil {
		t.Error(err)
	}

	if b.PersonID != 300001 {
		t.Errorf("Expected person to id to equal %d but got %d", 300001, b.PersonID)
	}
}

func TestGetPersons(t *testing.T) {
	bs, err := testClient.GetPersons()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Persons) < 1 {
		t.Error("Expected at least one person to be returned")
	}
}

func TestGetCommittee(t *testing.T) {
	b, err := testClient.GetCommittee(2640)
	if err != nil {
		t.Error(err)
	}

	if b.CommitteeID != 2640 {
		t.Errorf("Expected committee to id to equal %d but got %d", 2640, b.CommitteeID)
	}
}

func TestGetCommittees(t *testing.T) {
	bs, err := testClient.GetCommittees()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Committees) < 1 {
		t.Error("Expected at least one committee to be returned")
	}
}

func TestGetRole(t *testing.T) {
	b, err := testClient.GetRole(1)
	if err != nil {
		t.Error(err)
	}

	if b.RoleID != 1 {
		t.Errorf("Expected role to id to equal %d but got %d", 1, b.RoleID)
	}
}

func TestGetRoles(t *testing.T) {
	bs, err := testClient.GetRoles()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Roles) < 1 {
		t.Error("Expected at least one role to be returned")
	}
}

func TestGetCommitteeMember(t *testing.T) {
	b, err := testClient.GetCommitteeMember(207975)
	if err != nil {
		t.Error(err)
	}

	if b.CommitteeMemberID != 207975 {
		t.Errorf("Expected committee member to id to equal %d but got %d", 207975, b.CommitteeMemberID)
	}
}

func TestGetCommitteeMembers(t *testing.T) {
	bs, err := testClient.GetCommitteeMembers()
	if err != nil {
		t.Error(err)
	}

	if len(bs.CommitteeMembers) < 1 {
		t.Error("Expected at least one committee member to be returned")
	}
}

func TestGetCosponsoship(t *testing.T) {
	b, err := testClient.GetCosponsorship(1)
	if err != nil {
		t.Error(err)
	}

	if b.CosponshorshipID != 1 {
		t.Errorf("Expected cosponsorship to id to equal %d but got %d", 1, b.CosponshorshipID)
	}
}

func TestGetCosponsoships(t *testing.T) {
	bs, err := testClient.GetCosponsorships()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Cosponsorships) < 1 {
		t.Error("Expected at least one cosponsorship to be returned")
	}
}

func TestGetVote(t *testing.T) {
	b, err := testClient.GetVote(67617)
	if err != nil {
		t.Error(err)
	}

	if b.VoteID != 67617 {
		t.Errorf("Expected vote to id to equal %d but got %d", 67617, b.VoteID)
	}
}

func TestGetVotes(t *testing.T) {
	bs, err := testClient.GetVotes()
	if err != nil {
		t.Error(err)
	}

	if len(bs.Votes) < 1 {
		t.Error("Expected at least one vote to be returned")
	}
}

func TestGetVoteVoter(t *testing.T) {
	b, err := testClient.GetVoteVoter(28927519)
	if err != nil {
		t.Error(err)
	}

	if b.VoteVoterID != 28927519 {
		t.Errorf("Expected vote voter to id to equal %d but got %d", 28927519, b.VoteVoterID)
	}
}

func TestGetVoteVoters(t *testing.T) {
	bs, err := testClient.GetVoteVoters()
	if err != nil {
		t.Error(err)
	}

	if len(bs.VoteVoters) < 1 {
		t.Error("Expected at least one vote voters to be returned")
	}
}

func TestFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	c := testClient.Filter(query)

	if c.filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestBuildFilterQuery(t *testing.T) {
	expected1 := "?id=1&limit=200"
	expected2 := "?limit=200&id=1"

	query := Q{"id": "1", "limit": "200"}

	testClient.Filter(query)

	result := testClient.buildFilterQuery()

	if result != expected1 {
		if result != expected2 {
			t.Errorf("Expected query to equal %s or %sbut got %s", expected1, expected2, result)
		}
	}
}
