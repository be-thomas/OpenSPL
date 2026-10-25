package spl2

import (
	"testing"

	"github.com/be-thomas/openspl/pkg/diag"
)

func TestParseIsExplicitlyUnsupported(t *testing.T) {
	_, err := Parse(`from idx | where status = 500`)
	if err == nil {
		t.Fatal("expected an error")
	}
	ds, ok := diag.FromError(err)
	if !ok || len(ds) == 0 {
		t.Fatalf("error is not a diag.Error: %v", err)
	}
	if ds[0].Code != "spl2-not-implemented" {
		t.Fatalf("code = %q", ds[0].Code)
	}
	if ds[0].Pos.Line != 1 || ds[0].Pos.Col != 1 {
		t.Fatalf("position = %#v", ds[0].Pos)
	}
}
