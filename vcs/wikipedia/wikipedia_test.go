package wikipedia

import "testing"

// A random wikipedia article and user
const wikiPage = "Jack Parsons (rocket engineer)"
const wikiUsername = "John"
const minimumExpectedEdits = 30

func TestGetUserEdits(t *testing.T) {
	t.Logf("Querying " + wikiPage + " for user '" + wikiUsername + "'")
	count, err := GetUserEdits(wikiPage, wikiUsername)
	if err != nil {
		t.Error("Error: ", err)
		t.FailNow()
	}
	if count < minimumExpectedEdits {
		t.Errorf("GetUserEdits returned: %d, expected at least: %d", count, minimumExpectedEdits)
	}
}
