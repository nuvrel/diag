package diag

type Diagnostic struct {
	severity Severity
	message  string
	code     string
	detail   []string
	blocks   []block
}

func NewError(msg string) Diagnostic {
	return Diagnostic{severity: SeverityError, message: msg}
}

func NewWarning(msg string) Diagnostic {
	return Diagnostic{severity: SeverityWarning, message: msg}
}

func (d Diagnostic) Code(code string) Diagnostic {
	d.code = code

	return d
}

func (d Diagnostic) Detail(paragraphs ...string) Diagnostic {
	d.detail = paragraphs

	return d
}

func (d Diagnostic) Snippet(s Snippet) Diagnostic {
	d.blocks = append(d.blocks, &s)

	return d
}

func (d Diagnostic) Help(content string) Diagnostic {
	d.blocks = append(d.blocks, &help{content: content})

	return d
}

func (d Diagnostic) Note(content string) Diagnostic {
	d.blocks = append(d.blocks, &note{content: content})

	return d
}
