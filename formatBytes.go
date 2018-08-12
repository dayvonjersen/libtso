import "fmt"

var byteUnits = [6]string{"", "K", "M", "G", "T", "P"}

func formatBytes(i int) string {
	k := 0
	b := float64(i)
	for b >= 1024 {
		k++
		b /= 1024
	}
	return fmt.Sprintf("%.3f %sB", b, byteUnits[k])
}
