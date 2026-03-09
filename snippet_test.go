package diag_test

import (
	"bytes"
	"testing"

	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/exp/golden"
	"github.com/nuvrel/diag"
)

func TestSnippet(t *testing.T) {
	src := []byte(`package main

import (
	"errors"
	"strconv"
)

func fetch(id int) (string, error) {
	if id < 0 {
		return "", errors.New("negative id")
	}

	return "item-" + strconv.Itoa(id), nil
}`)

	cases := []struct {
		name string
		diag diag.Diagnostic
	}{
		{
			name: "with carets",
			diag: diag.NewError("you can't do that").
				Snippet(diag.NewSnippet(src).
					File("fetch.go").
					From(10, 14).
					To(10, 24).
					Message("dynamic error created here"),
				),
		},
		{
			name: "without location",
			diag: diag.NewError("you can't do that").Snippet(diag.NewSnippet(src)),
		},
		{
			name: "multiline",
			diag: diag.NewError("you can't do that").
				Snippet(diag.NewSnippet(src).
					File("fetch.go").
					From(9, 0).
					To(11, 0).
					Message("this whole block is wrong"),
				),
		},
		{
			name: "custom pad",
			diag: diag.NewError("you can't do that").
				Snippet(diag.NewSnippet(src).
					File("fetch.go").
					From(10, 14).
					To(10, 24).
					Pad(1),
				),
		},
		{
			name: "file only",
			diag: diag.NewError("you can't do that").Snippet(diag.NewSnippet(src).File("fetch.go")),
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

			if err := p.Print(c.diag); err != nil {
				t.Fatal(err)
			}

			golden.RequireEqual(t, buf.Bytes())
		})
	}
}
