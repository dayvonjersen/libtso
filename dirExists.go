import "os"

func dirExists(path string) bool {
	finfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	checkErr(err)
	return finfo.IsDir()
}
