package diag_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/exp/golden"
	"github.com/nuvrel/diag"
)

func TestPrinter(t *testing.T) {
	cases := []struct {
		name   string
		config diag.Config
		diags  []diag.Diagnostic
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
			name: "detail",
			diags: []diag.Diagnostic{
				diag.NewError("you can't do that").
					Detail(
						"This is a detailed explanation of the error.",
						"This is a second paragraph with more context.",
					),
			},
		},
		{
			name: "ordering",
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").
					Help("try this instead").
					Snippet(diag.NewSnippet([]byte("var y = x + z")).
						File("main.go").
						From(1, 13).
						To(1, 14).
						Message("z is not defined"),
					).
					Note("see the docs"),
			},
		},
		{
			name: "multiple",
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").Code("E001"),
				diag.NewWarning("something looks off"),
				diag.NewError("another error"),
			},
		},
		{
			name: "margin",
			config: diag.Config{
				Margin: 4,
			},
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").
					Detail("This is a detailed explanation.").
					Help("try this instead").
					Note("see the docs"),
			},
		},
		{
			name: "skip header",
			config: diag.Config{
				SkipHeader: true,
			},
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").
					Detail("This is a detailed explanation.").
					Help("try this instead"),
			},
		},
		{
			name: "skip detail",
			config: diag.Config{
				SkipDetail: true,
			},
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").
					Detail("This is a detailed explanation.").
					Help("try this instead"),
			},
		},
		{
			name: "header func",
			config: diag.Config{
				HeaderFunc: func(sev diag.Severity, label, code, message string) string {
					header := strings.ToUpper(label)

					if code != "" {
						header += " " + code
					}

					return header + " >> " + message
				},
			},
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").Code("E001"),
			},
		},
		{
			name: "detail func",
			config: diag.Config{
				DetailFunc: func(detail []string) string {
					return ">>> " + strings.Join(detail, " | ")
				},
			},
			diags: []diag.Diagnostic{
				diag.NewError("something went wrong").
					Detail("First paragraph.", "Second paragraph."),
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
				Margin:         c.config.Margin,
				HeaderFunc:     c.config.HeaderFunc,
				DetailFunc:     c.config.DetailFunc,
				SkipHeader:     c.config.SkipHeader,
				SkipDetail:     c.config.SkipDetail,
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
