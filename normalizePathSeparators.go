import "strings"

func normalizePathSeparators(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}
