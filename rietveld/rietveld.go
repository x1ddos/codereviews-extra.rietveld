package rietveld

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Rietveld API time format
type Timestamp time.Time

const timeFmt = "2006-01-02 15:04:05"

func (t Timestamp) MarshalJSON() ([]byte, error) {
	marshaled := time.Time(t).Format(timeFmt)
	return []byte(strconv.Quote(marshaled)), nil
}

func (t *Timestamp) UnmarshalJSON(s []byte) (err error) {
	uq, err := strconv.Unquote(string(s))
	if err != nil { return }

	v, err := time.Parse(timeFmt, uq)
	if err != nil { return }

	*t = Timestamp(v)
	return nil
}

type Issue struct {
	Id          uint `json:"issue"`
	Owner       string
	OwnerEmail  string `json:"owner_email"`
	Reviewers   []string
	Cc          []string
	Subject     string
	Description string
	BaseUrl     string `json:"base_url"`
	PatchsetIds []uint `json:"patchsets"`
	Private     bool
	Closed      bool
	Created     Timestamp
	Modified    Timestamp
}

func (i Issue) String() string {
	return fmt.Sprintf(
		"[%d] %s\nBase URL: %s\nOwner: %s (%s)\nReviewers: %s\n"+
			"Private: %t\nClosed: %t\nUpdated: %s\n",
		i.Id, i.Subject, i.BaseUrl, i.Owner, i.OwnerEmail, i.Reviewers,
		i.Private, i.Closed, time.Time(i.Modified).Format(time.UnixDate))
}

type IssuesList struct {
	Cursor string
	Issues []Issue `json:"results"`
}

func (r IssuesList) String() string {
	return fmt.Sprintf("Issues count: %d, Cursor: %s", len(r.Issues), r.Cursor)
}

func Search(client *http.Client) (r *IssuesList, e error) {
	if client == nil {
		client = http.DefaultClient
	}
	url := "https://codereview.appspot.com/search?format=json&limit=10"
	resp, err := client.Get(url)
	if err != nil {
		e = err
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e = errors.New(resp.Status)
		return
	}

	r = new(IssuesList)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		e = err
		return
	}

	return
}
