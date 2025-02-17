package lara

import (
	"crypto/rand"
	"os"
)

const (
	randomString = "abcdmcnzjsahdfieqpqowejalkz2712981AMALSKPWIQUEYRTFGCBZ,NCZPAOKSLAJ+-_"
)

// RandomString generate a random string
func (l *Lara) RandomString(size int) string {
	s, r := make([]rune, size), []rune(randomString)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}

func (l *Lara) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lara) CreateFileIfNotExists(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer func(file *os.File) { _ = file.Close() }(file)
	}

	return nil
}
