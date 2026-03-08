package generator

type CommandConfig struct {
	Rows   int
	Files  int
	OutDir string
}

func (c CommandConfig) Run() error {
	return nil
}
