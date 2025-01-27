package lara

const version = "1.0.0"

func (l *Lara) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"views", "handlers", "migrations", "data", "public", "tmp", "logs", "middleware"},
	}

	err := l.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lara) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := l.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}
