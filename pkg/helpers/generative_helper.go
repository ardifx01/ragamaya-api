package helpers

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	nonAlnum    = regexp.MustCompile(`[^a-z0-9]+`)
	trimHyphens = regexp.MustCompile(`^-+|-+$`)
)

func Slugify(s string) string {
	if s == "" {
		return ""
	}

	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	normalized, _, _ := transform.String(t, s)

	normalized = strings.ToLower(strings.TrimSpace(normalized))
	normalized = strings.ReplaceAll(normalized, "'", "")
	normalized = nonAlnum.ReplaceAllString(normalized, "-")
	normalized = trimHyphens.ReplaceAllString(normalized, "")

	return normalized
}

func SlugifyUnique(s string) string {
	base := Slugify(s)
	id := uuid.New().String()[:8]

	return fmt.Sprintf("%s-%s", base, id)
}
