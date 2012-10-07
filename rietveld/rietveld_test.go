package rietveld

import (
	"testing"
	"encoding/json"

	"os"
	"path"
)

func getFixture(filename string) (f *os.File, e error) {
	filepath := path.Join("..", "fixtures", filename)
	f, e = os.Open(filepath)
	return
}

func TestSearchJsonDecoding(t *testing.T) {
	file, err := getFixture("search.json")
	if err != nil {
		t.Error(err)
	}

	r := new(IssuesList)
	if err := json.NewDecoder(file).Decode(r); err != nil {
		t.Error(err)
	}

	if len(r.Cursor) == 0 {
		t.Error("Expected Cursor to be present")
	}
	if len(r.Issues) != 2 {
		t.Errorf("Expected 2 Issue items, got %d", len(r.Issues))
	}

	i := r.Issues[0]
	if i.Id != 12345 {
		t.Errorf("Expected Id to be 12345, go %d", i.Id)
	}
	if i.BaseUrl != "http://svn.some.base/url" {
		t.Errorf("Expected BaseUrl to be http://svn.some.base/url" +
			     ", got %s", i.BaseUrl)
	}
	if len(i.Cc) != 1 {
		t.Errorf("Expected 1 item in Cc field, got %d", len(i.Cc))
	}
	if i.Cc[0] != "dude@example.org" {
		t.Errorf("Expected Cc[0] to be dude@example.org, got %s", i.Cc[0])
	}
}
