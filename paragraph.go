package diag

var (
	_ block = (*help)(nil)
	_ block = (*note)(nil)
)

type help struct {
	content string
}

func (*help) block() {}

type note struct {
	content string
}

func (*note) block() {}
