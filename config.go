package diag

import (
	"os"

	"github.com/charmbracelet/colorprofile"
)

type Characters struct {
	Top  string
	Mid  string
	Bot  string
	Dash string
	Dot  string
}

func DefaultCharacters() Characters {
	return Characters{
		Top:  "╭",
		Mid:  "│",
		Bot:  "╰",
		Dash: "─",
		Dot:  "·",
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
}

func DefaultConfig() Config {
	return Config{
		Profile:        colorprofile.Detect(os.Stdout, os.Environ()),
		Theme:          DefaultTheme(),
		Characters:     DefaultCharacters(),
		Prefixes:       DefaultPrefixes(),
		SeverityLabels: DefaultSeverityLabels(),
	}
}
