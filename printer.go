package diag

import (
	"bytes"
	"fmt"
	"io"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
)

type printer interface {
	print()
}

type Printer struct {
	writer io.Writer
	buffer bytes.Buffer
	config Config
}

func NewPrinter(w io.Writer, c Config) *Printer {
	return &Printer{writer: w, config: c}
}

func (p *Printer) Print(diags ...Diagnostic) error {
	p.buffer.Reset()

	for i, d := range diags {
		if i > 0 {
			p.writeln()
		}

		p.print(d)
	}

	return p.flush()
}

func (p *Printer) print(d Diagnostic) {
	p.printHeader(d)

	for _, b := range d.blocks {
		switch b := b.(type) {
		case *Snippet:
			(&snippeter{printer: p, snippet: b, severity: d.severity}).print()
		case *help:
			p.writeStyled(p.config.Theme.Help, p.config.Prefixes.Help+": ")
			p.write(b.content)
			p.writeln()
		case *note:
			p.writeStyled(p.config.Theme.Note, p.config.Prefixes.Note+": ")
			p.write(b.content)
			p.writeln()
		}
	}
}

func (p *Printer) flush() error {
	_, err := p.writer.Write(p.buffer.Bytes())
	if err != nil {
		return fmt.Errorf("flushing buffer: %w", err)
	}

	return nil
}

func (p *Printer) write(text string) {
	p.buffer.WriteString(text)
}

func (p *Printer) writeln() {
	p.buffer.WriteByte('\n')
}

func (p *Printer) styled(s lipgloss.Style, text string) string {
	if p.config.Profile != colorprofile.Ascii && p.config.Profile != colorprofile.NoTTY {
		return s.Render(text)
	}

	return text
}

func (p *Printer) writeStyled(s lipgloss.Style, text string) {
	p.write(p.styled(s, text))
}

func (p *Printer) styleFor(s Severity) lipgloss.Style {
	switch s {
	case SeverityError:
		return p.config.Theme.Error
	case SeverityWarning:
		return p.config.Theme.Warning
	}

	return lipgloss.NewStyle()
}

func (p *Printer) printHeader(d Diagnostic) {
	label := p.config.SeverityLabels.labelFor(d.severity)

	if d.code != "" {
		label += "[" + d.code + "]"
	}

	label += ":"

	p.writeStyled(p.styleFor(d.severity), label)
	p.write(" ")
	p.writeStyled(p.config.Theme.Message, d.message)
	p.writeln()
}
