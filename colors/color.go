package color

import "fmt"

type (
	Black     string
	Red       string
	Green     string
	Yellow    string
	Blue      string
	Magenta   string
	Cyan      string
	White     string
	Default   string
	BlackBg   string
	RedBg     string
	GreenBg   string
	YellowBg  string
	BlueBg    string
	MagentaBg string
	CyanBg    string
	WhiteBg   string
	DefaultBg string
)

func (s Black) Format(f fmt.State, c rune)   { f.Write(bytes("30", string(s))) }
func (s Red) Format(f fmt.State, c rune)     { f.Write(bytes("31", string(s))) }
func (s Green) Format(f fmt.State, c rune)   { f.Write(bytes("32", string(s))) }
func (s Yellow) Format(f fmt.State, c rune)  { f.Write(bytes("33", string(s))) }
func (s Blue) Format(f fmt.State, c rune)    { f.Write(bytes("34", string(s))) }
func (s Magenta) Format(f fmt.State, c rune) { f.Write(bytes("35", string(s))) }
func (s Cyan) Format(f fmt.State, c rune)    { f.Write(bytes("36", string(s))) }
func (s White) Format(f fmt.State, c rune)   { f.Write(bytes("37", string(s))) }
func (s Default) Format(f fmt.State, c rune) { f.Write(bytes("39", string(s))) }

func (s BlackBg) Format(f fmt.State, c rune)   { f.Write(bytes("40", string(s))) }
func (s RedBg) Format(f fmt.State, c rune)     { f.Write(bytes("41", string(s))) }
func (s GreenBg) Format(f fmt.State, c rune)   { f.Write(bytes("42", string(s))) }
func (s YellowBg) Format(f fmt.State, c rune)  { f.Write(bytes("43", string(s))) }
func (s BlueBg) Format(f fmt.State, c rune)    { f.Write(bytes("44", string(s))) }
func (s MagentaBg) Format(f fmt.State, c rune) { f.Write(bytes("45", string(s))) }
func (s CyanBg) Format(f fmt.State, c rune)    { f.Write(bytes("46", string(s))) }
func (s WhiteBg) Format(f fmt.State, c rune)   { f.Write(bytes("47", string(s))) }
func (s DefaultBg) Format(f fmt.State, c rune) { f.Write(bytes("49", string(s))) }

func bytes(code, s string) []byte {
	return []byte("\033[" + code + "m" + s + "\033[0m")
}
