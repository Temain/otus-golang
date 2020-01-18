package copier

import (
	"io"
	"os"
)

func Copy(from string, to string, limit int, offset int) error {
	sourceFile, _ := os.Open(from)
	defer sourceFile.Close()
	if _, err := sourceFile.Seek(int64(offset), 0); err != nil {
		return err
	}

	targetFile, _ := os.Create(to)
	defer targetFile.Close()
	if _, err := io.CopyN(targetFile, sourceFile, int64(limit)); err != nil {
		return err
	}

	return nil
}
