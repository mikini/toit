// Code generated by tedi; DO NOT EDIT.

package path

import (
	"github.com/jstroem/tedi"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	t := tedi.New(m)

	// TestLabels:
	t.TestLabel("unit")
	t.TestLabel("integration")
	t.TestLabel("regression")

	os.Exit(t.Run())
}
