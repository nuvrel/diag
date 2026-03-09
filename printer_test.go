package diag_test

import (
	"bytes"
	"testing"

	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/exp/golden"
	"github.com/nuvrel/diag"
)

func TestPrinter(t *testing.T) {
	cases := []struct {
		name  string
		diags []diag.Diagnostic
	}{
		{
			name:  "error",
			diags: []diag.Diagnostic{diag.NewError("something went wrong")},
		},
		{
			name:  "error with code",
			diags: []diag.Diagnostic{diag.NewError("something went wrong").Code("E001")},
		},
		{
			name:  "warning",
			diags: []diag.Diagnostic{diag.NewWarning("something looks off")},
		},
		{
			name:  "help and note",
			diags: []diag.Diagnostic{diag.NewError("you can't do that").Help("try this instead").Note("see the docs")},
		},
		{
			name: "multiple",
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").Code("E001"),
				diag.NewWarning("something looks off"),
				diag.NewError("another error"),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			config := diag.Config{
				Profile:        colorprofile.NoTTY,
				Theme:          diag.DefaultTheme(),
				Characters:     diag.DefaultCharacters(),
				Prefixes:       diag.DefaultPrefixes(),
				SeverityLabels: diag.DefaultSeverityLabels(),
			}

			var buf bytes.Buffer

			p := diag.NewPrinter(&buf, config)

			if err := p.Print(c.diags...); err != nil {
				t.Fatal(err)
			}

			golden.RequireEqual(t, buf.Bytes())
		})
	}
}
