import "regexp"

var escapeshellstringre = regexp.MustCompile(`([\(\)\[\]\{\}\$\#&;` + "`" + `\|\*\?~<>\^'"\s])`)

func escapeShellString(str string) string {
	return escapeshellstringre.ReplaceAllString(str, "\\$1")
}
