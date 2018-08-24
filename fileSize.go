import "os"

func fileSize(filename string) int {
	f, err := os.Open(filename)
	checkErr(err)
	defer checkErr(f.Close())
	finfo, err := f.Stat()
	checkErr(err)
	return int(finfo.Size())
}
