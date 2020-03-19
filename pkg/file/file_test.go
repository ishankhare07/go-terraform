package file

import (
	"testing"

	"github.com/ishankhare07/go-terraform/pkg/blocks"
)

func TestEmptyResource(t *testing.T) {
	r := blocks.NewResource("google_container_cluster", "primary")

	f := NewFile()

	f.AddBlock(r)

	expected := `resource "google_container_cluster" "primary" {
}
`
	if expected != string(f.Output()) {
		t.Errorf("could not get expected output.\nwant = %s\nhave = %s", expected, f.Output())
	}
}
