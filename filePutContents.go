import (
	"io"
	"os"
)

func filePutContents(filename, contents string) {
	f, err := os.Create(filename)
	checkErr(err)
	_, err = io.WriteString(f, contents)
	checkErr(err)
	checkErr(f.Close())
}
