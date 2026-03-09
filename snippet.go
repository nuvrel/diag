package diag

const (
	defaultPad      = 2
	defaultTabWidth = 4
)

var (
	_ block = (*Snippet)(nil)
)

type Snippet struct {
	loc      location
	content  []byte
	message  string
	pad      int
	tabWidth int
}

func (*Snippet) block() {}

func NewSnippet(content []byte) Snippet {
	return Snippet{content: content}
}

func (s Snippet) File(name string) Snippet {
	s.loc.source = name

	return s
}

func (s Snippet) From(line, col int) Snippet {
	s.loc.start = pos{line: line, col: col}

	return s
}

func (s Snippet) To(line, col int) Snippet {
	s.loc.end = pos{line: line, col: col}

	return s
}

func (s Snippet) Message(msg string) Snippet {
	s.message = msg

	return s
}

func (s Snippet) Pad(lines int) Snippet {
	s.pad = lines

	return s
}

func (s Snippet) TabWidth(width int) Snippet {
	s.tabWidth = width

	return s
}

func (s Snippet) effectivePad() int {
	if s.pad > 0 {
		return s.pad
	}

	return defaultPad
}

func (s Snippet) effectiveTabWidth() int {
	if s.tabWidth > 0 {
		return s.tabWidth
	}

	return defaultTabWidth
}

func (s Snippet) isSingleLine() bool {
	return s.loc.start.line == s.loc.end.line
}

func (s Snippet) containsLine(num int) bool {
	return num >= s.loc.start.line && num <= s.loc.end.line
}

func (s Snippet) hasCarets() bool {
	return s.loc.start.col > 0
}

func (s Snippet) hasContent() bool {
	return s.content != nil
}

func (s Snippet) hasMessage() bool {
	return s.message != ""
}
