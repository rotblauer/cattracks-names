package names

import (
	"path/filepath"
	"regexp"
	"strings"
)

// Aliases is a map of regular expressions to cat aliases, can be used to index cats by name; all aliases are unique.
var Aliases = map[*regexp.Regexp]string{
	regexp.MustCompile(`(?i)(Rye.*|Kitty.*|jl)`):                          "rye",
	regexp.MustCompile(`(?i)(.*Papa.*|P2|Isaac.*|.*moto.*|iha|ubp52)`):    "ia",
	regexp.MustCompile(`(?i)(Big.*Ma.*)`):                                 "jr",
	regexp.MustCompile(`(?i)Kayleigh.*`):                                  "kd",
	regexp.MustCompile(`(?i)(KK.*|kek)`):                                  "kk",
	regexp.MustCompile(`(?i)Bob.*`):                                       "rj",
	regexp.MustCompile(`(?i)(Pam.*|Rathbone.*)`):                          "pr",
	regexp.MustCompile(`(?i)(Ric|.*A3_Pixel_XL.*|.*marlin-Pixel-222d.*)`): "ric",
	regexp.MustCompile(`(?i)Twenty7.*`):                                   "mat",
	regexp.MustCompile(`(?i)(.*Carlomag.*|JLC|jlc)`):                      "jlc",
}

// AliasOrName returns the alias for a name, or the name if no alias is found.
func AliasOrName(name string) string {
	if name == "" {
		return "_"
	}
	for r, a := range Aliases {
		if r.MatchString(name) {
			return a
		}
	}
	return name
}

// InvalidNameChars are characters that are not considered safe/sane/sanitized for cat names. Thanks Copilot.
var InvalidNameChars = []string{" ", string(filepath.Separator), "(", ")", "[", "]", "{", "}", "'", "\"", "`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "+", "=", "|", "\\", ";", ":", "<", ">", ",", ".", "?", "/", "\n", "\t"}

// InvalidReplacementChar is the character that will replace invalid characters in a name.
const InvalidReplacementChar = "_"

// SanitizeName returns a sanitized version of the name or alias.
func SanitizeName(name string) string {
	if name == "" {
		return InvalidReplacementChar
	}
	for _, c := range InvalidNameChars {
		name = strings.ReplaceAll(name, c, InvalidReplacementChar)
	}

	return name
}
