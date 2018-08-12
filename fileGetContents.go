import (
	"bytes"
	"io"
	"os"
)

func fileGetContents(filename string) string {
	contents := new(bytes.Buffer)
	f, err := os.Open(filename)
	checkErr(err)
	_, err = io.Copy(contents, f)
	f.Close()
	if err != io.EOF {
		checkErr(err)
	}
	return contents.String()
}
