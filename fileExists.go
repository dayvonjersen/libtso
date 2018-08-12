import (
	"os"
)

func fileExists(filename string) bool {
	f, err := os.Open(filename)
	checkErr(f.Close())
	if os.IsNotExist(err) {
		return false
	}
	checkErr(err)
	return true
}
