package diag

type Severity int

const (
	SeverityError Severity = iota
	SeverityWarning
)

type SeverityLabels struct {
	Error   string
	Warning string
}

func DefaultSeverityLabels() SeverityLabels {
	return SeverityLabels{
		Error:   "error",
		Warning: "warning",
	}
}

func (sl SeverityLabels) labelFor(s Severity) string {
	switch s {
	case SeverityError:
		return sl.Error
	case SeverityWarning:
		return sl.Warning
	}

	return ""
}
