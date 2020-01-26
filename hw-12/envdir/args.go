package envdir

import (
	"errors"
	"os"
	"strings"
)

// ParseArgs получает слайс аргументов и определяет в них директорию и команду с аргументами.
func ParseArgs(args []string) (dir string, cmd []string, err error) {
	if len(args) < 2 {
		err = errors.New("wrong arguments, should be dir and cmd")
		return "", nil, err
	}

	dir = args[0]
	if len(strings.TrimSpace(dir)) == 0 {
		err = errors.New("empty dir argument")
		return "", nil, err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = errors.New("directory is not exists")
		return "", nil, err
	}

	cmd = args[1:]
	if len(strings.TrimSpace(cmd[0])) == 0 {
		err = errors.New("empty cmd argument")
		return dir, nil, err
	}

	return dir, cmd, nil
}
