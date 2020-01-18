package copier

import (
	"github.com/cheggaaa/pb"
	"io"
	"os"
)

func Copy(from string, to string, limit int, offset int) error {
	sourceFile, _ := os.Open(from)
	defer sourceFile.Close()
	if _, err := sourceFile.Seek(int64(offset), 0); err != nil {
		return err
	}

	bar := pb.Full.Start64(int64(limit))
	sourceReader := bar.NewProxyReader(sourceFile)

	targetFile, _ := os.Create(to)
	defer targetFile.Close()
	_, err := io.CopyN(targetFile, sourceReader, int64(limit))
	bar.Finish()

	return err
}
