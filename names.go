package names

import (
	"path/filepath"
	"regexp"
	"strings"
)

var AliasesToMatchers = map[string]string{
	"rye": `(?i)(Rye.*|Kitty.*|jl)`,
	"ia":  `(?i)(.*Papa.*|P2|Isaac.*|.*moto.*|iha|ubp52)`,
	"jr":  `(?i)(Big.*Ma.*)`,
	"kd":  `(?i)Kayleigh.*`,
	"kk":  `(?i)(KK.*|kek)`,
	"rj":  `(?i)Bob.*`,
	"pr":  `(?i)(Pam.*|Rathbone.*)`,
	"ric": `(?i)(Ric|.*A3_Pixel_XL.*|.*marlin-Pixel-222d.*)`,
	"mat": `(?i)Twenty7.*`,
	"jlc": `(?i)(.*Carlomag.*|JLC|jlc)`,
}

var MatchersToAliases = func() map[string]string {
	m := make(map[string]string, len(AliasesToMatchers))
	for a, r := range AliasesToMatchers {
		m[r] = a
	}
	return m
}()

// MatchingAliases is a map of regular expressions to cat aliases, can be used to index cats by name; all aliases are unique.
var MatchingAliases = func() map[*regexp.Regexp]string {
	m := make(map[*regexp.Regexp]string, len(AliasesToMatchers))
	for r, a := range MatchersToAliases {
		m[regexp.MustCompile(r)] = a
	}
	return m
}()

// AliasOrName returns the alias for a name, or the name if no alias is found.
func AliasOrName(name string) string {
	for r, a := range MatchingAliases {
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

// AliasOrSanitizedName returns the alias for a name, or the name if no alias is found, sanitized.
func AliasOrSanitizedName(name string) string {
	return SanitizeName(AliasOrName(name))
}
