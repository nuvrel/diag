package diag

import (
	"os"

	"github.com/charmbracelet/colorprofile"
)

type Characters struct {
	Top      string
	Mid      string
	Bot      string
	Dash     string
	Dot      string
	HintHelp string
	HintNote string
}

func DefaultCharacters() Characters {
	return Characters{
		Top:      "╭",
		Mid:      "│",
		Bot:      "╰",
		Dash:     "─",
		Dot:      "·",
		HintHelp: "›",
		HintNote: "≋",
	}
}

type Prefixes struct {
	Help string
	Note string
}

func DefaultPrefixes() Prefixes {
	return Prefixes{
		Help: "help",
		Note: "note",
	}
}

type Config struct {
	Profile        colorprofile.Profile
	Theme          Theme
	Characters     Characters
	Prefixes       Prefixes
	SeverityLabels SeverityLabels
	DetailPad      int
}

func DefaultConfig() Config {
	return Config{
		Profile:        colorprofile.Detect(os.Stdout, os.Environ()),
		Theme:          DefaultTheme(),
		Characters:     DefaultCharacters(),
		Prefixes:       DefaultPrefixes(),
		SeverityLabels: DefaultSeverityLabels(),
		DetailPad:      2,
	}
}

func (c Characters) hint(ch string) string {
	if ch == "" {
		return ""
	}

	return ch + " "
}

func (c Config) effectiveDetailPad() int {
	if c.DetailPad > 0 {
		return c.DetailPad
	}

	return 2
}
