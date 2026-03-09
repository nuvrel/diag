package diag

import (
	"bytes"
	"strconv"
	"strings"
	"unicode/utf8"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

var (
	_ printer = (*snippeter)(nil)
)

type snippeter struct {
	printer  *Printer
	snippet  *Snippet
	severity Severity
}

func (s snippeter) print() {
	start := max(1, s.snippet.loc.start.line)
	end := max(start, s.snippet.loc.end.line)
	pad := s.snippet.effectivePad()
	width := len(strconv.Itoa(end + pad))

	s.printHeader(width)
	s.printSpacer(width)

	if s.snippet.hasContent() {
		s.printLines(width, start, end, pad)
	}

	s.printSpacer(width)
	s.printBottom(width)
}

func (s snippeter) printHeader(width int) {
	header := s.printer.config.Characters.Top + s.printer.config.Characters.Dash

	if loc := s.snippet.loc.String(); loc != "" {
		header += "[" + loc + "]"
	}

	s.writeBorder(width, header)
}

func (s snippeter) printBottom(width int) {
	bottom := s.printer.config.Characters.Bot + s.printer.config.Characters.Dash

	s.writeBorder(width, bottom)
}

func (s snippeter) printSpacer(width int) {
	s.writeGutter(width, s.printer.config.Characters.Mid)
	s.printer.writeln()
}

func (s snippeter) writeGutter(width int, spacer string) {
	s.writeIndent(width)
	s.printer.writeStyled(s.printer.config.Theme.Muted, spacer+" ")
}

func (s snippeter) writeBorder(width int, content string) {
	s.writeIndent(width)
	s.printer.writeStyled(s.printer.config.Theme.Muted, content)
	s.printer.writeln()
}

func (s snippeter) writeIndent(width int) {
	s.printer.write(strings.Repeat(" ", width+3))
}

func (s snippeter) printLines(width, start, end, pad int) {
	lines := bytes.Split(s.snippet.content, []byte("\n"))
	first, last := max(1, start-pad), min(len(lines), end+pad)

	for idx, raw := range lines[first-1 : last] {
		num := first + idx
		content := string(raw)

		style := s.printer.config.Theme.Muted

		if s.snippet.containsLine(num) {
			style = s.printer.styleFor(s.severity)
		}

		s.writeLineNumber(num, width, style)
		s.printer.write(s.expandTabs(content))
		s.printer.writeln()

		isCaretLine := num == start

		if s.snippet.isSingleLine() && isCaretLine && s.snippet.hasCarets() {
			s.writeGutter(width, s.printer.config.Characters.Dot)
			s.writeCarets(content)
		}
	}

	if !s.snippet.isSingleLine() && s.snippet.hasMessage() {
		s.printSpacer(width)
		s.writeGutter(width, s.printer.config.Characters.Dot)
		s.printer.write(s.snippet.message)
		s.printer.writeln()
	}
}

func (s snippeter) writeCarets(line string) {
	start := s.visualColumn(line, s.snippet.loc.start.col-1)
	end := s.visualColumn(line, s.snippet.loc.end.col-1)
	width := ansi.StringWidth(s.expandTabs(line))
	start, end = max(0, min(start, width)), max(start+1, min(end, width))

	s.printer.write(strings.Repeat(" ", start))
	s.printer.writeStyled(s.printer.styleFor(s.severity), strings.Repeat("^", end-start))

	if s.snippet.hasMessage() {
		s.printer.write(strings.Repeat(" ", 2) + s.snippet.message)
	}

	s.printer.writeln()
}

func (s snippeter) writeLineNumber(num, width int, style lipgloss.Style) {
	str := strconv.Itoa(num)

	s.printer.write(strings.Repeat(" ", width-len(str)+2))
	s.printer.writeStyled(style, str)
	s.printer.writeStyled(s.printer.config.Theme.Muted, " "+s.printer.config.Characters.Mid+" ")
}

func (s snippeter) expandTabs(line string) string {
	if !strings.ContainsRune(line, '\t') {
		return line
	}

	tab := s.snippet.effectiveTabWidth()
	col := 0

	var b strings.Builder

	for _, ch := range line {
		if ch != '\t' {
			b.WriteRune(ch)
			col++

			continue
		}

		pad := tab - (col % tab)
		b.WriteString(strings.Repeat(" ", pad))
		col += pad
	}

	return b.String()
}

func (s snippeter) visualColumn(line string, offset int) int {
	if offset <= 0 {
		return 0
	}

	tab := s.snippet.effectiveTabWidth()
	col, pos := 0, 0

	for _, ch := range line {
		if pos >= offset {
			break
		}

		if ch != '\t' {
			col++
			pos += utf8.RuneLen(ch)

			continue
		}

		col += tab - (col % tab)
		pos++
	}

	return col
}
