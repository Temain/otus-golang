package copier

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"os"
)

// Copy копирует исходный файл частично или полностью.
func Copy(from string, to string, limit int, offset int) (err error) {
	// Открытие исходного файла для чтения.
	sourceFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Проверка размера сдвига и сдвиг в исходном файле.
	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	sourceFileSize := sourceFileInfo.Size()
	if sourceFileSize <= int64(offset) {
		return fmt.Errorf("offset is too big, max %v", sourceFileSize-1)
	}
	if offset > 0 {
		_, err = sourceFile.Seek(int64(offset), 0)
		if err != nil {
			return err
		}
	}

	// Обертка для вывода прогресса копирования.
	barSize := sourceFileSize
	if limit != 0 {
		barSize = int64(limit)
	}
	bar := pb.Full.Start64(barSize)
	sourceReader := bar.NewProxyReader(sourceFile)

	// Копирование исходного файла.
	targetFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer targetFile.Close()
	if limit != 0 {
		_, err = io.CopyN(targetFile, sourceReader, int64(limit))
	} else {
		_, err = io.Copy(targetFile, sourceReader)
	}

	bar.Finish()

	return err
}
