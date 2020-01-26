package envdir

import (
	"io/ioutil"
	"path/filepath"
)

// ReadDir вычитывает переменные окружения из файлов в директории.
func ReadDir(dir string) (map[string]string, error) {
	env := make(map[string]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return env, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(dir, file.Name())
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return env, err
		}
		env[file.Name()] = string(bytes)
	}

	return env, nil
}
