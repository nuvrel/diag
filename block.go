package diag

import "strconv"

type block interface {
	block()
}

type pos struct {
	line int
	col  int
}

type location struct {
	source string
	start  pos
	end    pos
}

func (l location) String() string {
	hasLine := l.start.line != 0

	if !hasLine && l.source == "" {
		return ""
	}

	if !hasLine {
		return l.source
	}

	return l.source + ":" + strconv.Itoa(l.start.line) + ":" + strconv.Itoa(l.start.col)
}
